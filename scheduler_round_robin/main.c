#include "header.h"

/**
 * timeslice - квант времени, который нужно использовать.
 * Поток смещается с CPU, если пока он занимал CPU функция
 * timer_tick была вызвана timeslice раз.
 **/
void				scheduler_setup(int timeslice)
{
	delete_scheduler(&g_scheduler);
	if ((g_scheduler = init_scheduler(timeslice)) == NULL)
		return ;
}

/**
 * Функция вызывается, когда создается новый поток управления.
 * thread_id - идентификатор этого потока и гарантируется, что
 * никакие два потока не могут иметь одинаковый идентификатор.
 **/
void				new_thread(int thread_id)
{
	t_thread    	*new_thread;

	if ((new_thread = create_thread(thread_id)) == NULL)
		return ;
	if (g_scheduler->running_thread == NULL)
		g_scheduler->running_thread = new_thread;
	else
		ft_lstpushback(&g_scheduler->thread_queue, ft_lstnew(new_thread, sizeof(t_thread *)));
}

/**
 * Функция вызывается, когда поток, исполняющийся на CPU,
 * завершается. Завершится может только поток, который сейчас
 * исполняется, поэтому thread_id не передается. CPU должен
 * быть отдан другому потоку, если есть активный
 * (незаблокированный и незавершившийся) поток.
 **/
void				exit_thread()
{
	t_list			*tmp;

	delete_thread(&g_scheduler->running_thread);
	tmp = ft_extracthead(g_scheduler->thread_queue);
	if (tmp != NULL)
		g_scheduler->running_thread = (t_thread *)tmp->content;
	else
		g_scheduler->running_thread = NULL;
	tmp->content = NULL;
	ft_lstdelone(&tmp, del_content_thread);
	g_scheduler->running_time_thread = 0;
}

/**
 * Функция вызывается, когда поток, исполняющийся на CPU,
 * блокируется. Заблокироваться может только поток, который
 * сейчас исполняется, поэтому thread_id не передается. CPU
 * должен быть отдан другому активному потоку, если таковой
 * имеется.
 **/
void				block_thread()
{
	t_list			*tmp;

	ft_lstpushback(&g_scheduler->blocked_thread, ft_lstnew(g_scheduler->running_thread, sizeof(t_thread *)));
	tmp = ft_extracthead(g_scheduler->thread_queue);
	if (tmp != NULL)
		g_scheduler->running_thread = (t_thread *)tmp->content;
	else
		g_scheduler->running_thread = NULL;
	tmp->content = NULL;
	ft_lstdelone(&tmp, del_content_thread);
	g_scheduler->running_time_thread = 0;
}

/**
 * Функция вызывается, когда один из заблокированных потоков
 * разблокируется. Гарантируется, что thread_id - идентификатор
 * ранее заблокированного потока.
 **/
void				wake_thread(int thread_id)
{
	t_list			*list_runner;
	t_list			*tmp;

	list_runner = g_scheduler->blocked_thread;
	while (list_runner != NULL)
	{
		if (((t_thread *)list_runner->content)->id == thread_id)
			break ;
		list_runner = list_runner->next;
	}
	tmp = list_runner;
	list_runner = list_runner->next;
	tmp->next = NULL;
	ft_lstpushback(&g_scheduler->thread_queue, tmp);
}

/**
 * Ваш таймер. Вызывается каждый раз, когда проходит единица
 * времени.
 **/
void				timer_tick()
{
	t_list			*tmp;

	g_scheduler->running_time_thread++;
	if (g_scheduler->running_time_thread == g_scheduler->timeslice)
	{
		g_scheduler->running_time_thread = 0;
		ft_lstpushback(&g_scheduler->thread_queue, ft_lstnew(g_scheduler->running_thread, sizeof(t_thread *)));
		tmp = ft_extracthead(g_scheduler->thread_queue);
		if (tmp != NULL)
			g_scheduler->running_thread = (t_thread *)tmp->content;
		else
			g_scheduler->running_thread = NULL;
		tmp->content = NULL;
		ft_lstdelone(&tmp, del_content_thread);
	}
}

/**
 * Функция должна возвращать идентификатор потока, который в
 * данный момент занимает CPU, или -1 если такого потока нет.
 * Единственная ситуация, когда функция может вернуть -1, это
 * когда нет ни одного активного потока (все созданные потоки
 * либо уже завершены, либо заблокированы).
 **/
int 				current_thread()
{
	if (g_scheduler->running_thread != NULL)
		return g_scheduler->running_thread->id;
	return -1;
}

int					main(void)
{
	scheduler_setup(2);

/* Test 01 */	
	printf("Test 01\n");
	new_thread(1);
	printf("expected: 1, actual: %d\n", current_thread());
	exit_thread();
	printf("expected: -1, actual: %d\n", current_thread());
	printf("\n");

/* Test 02 */
	printf("Test 02\n");
	new_thread(1);
	printf("expected: 1, actual: %d\n", current_thread());
	timer_tick();
	timer_tick();
	printf("expected: 1, actual: %d\n", current_thread());
	printf("\n");

/* Test 03 */
	printf("Test 03\n");
	new_thread(1);
	printf("expected: 1, actual: %d\n", current_thread());
	new_thread(2);
	printf("expected: 1, actual: %d\n", current_thread());
	timer_tick();
	timer_tick();
	printf("expected: 2, actual: %d\n", current_thread());
	timer_tick();
	timer_tick();
	printf("expected: 1, actual: %d\n", current_thread());
	printf("\n");
}
