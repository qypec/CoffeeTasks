/*
	Вам нужно реализовать работу со списком таких дел на месяц,
	а именно, реализовать набор следующих операций:
		ADD i s: Назначить дело с названием s на день i текущего месяца.
		DUMP i: Вывести все дела, запланированные на день i текущего месяца.
		NEXT:
			Перейти к списку дел на новый месяц. При выполнении данной команды вместо
			текущего (старого) списка дел на текущий месяц создаётся и становится
			активным (новый) список дел на следующий месяц: все дела со старого списка
			дел копируются в новый список. После выполнения данной команды новый список
			дел и следующий месяц становятся текущими, а работа со старым списком дел прекращается.
			При переходе к новому месяцу необходимо обратить внимание на разное количество дней в месяцах:
				если следующий месяц имеет больше дней, чем текущий, «дополнительные» дни
				необходимо оставить пустыми (не содержащими дел);
				если следующий месяц имеет меньше дней, чем текущий, дела со всех «лишних»
				дней необходимо переместить на последний день следующего месяца.
*/

#include <iostream>
#include <map>
#include <vector>

using namespace std;

void						next(vector<vector<string>>& list, int del)
{
	int						newBack = list.size() - del - 1;

	if (del > 0)
	{
		for (auto i = list.size() - 1; i > newBack; i--)
		{
			for (auto t : list[i])
				list[newBack].push_back(t);
		}
	}
	list.resize(list.size() - del);
}

int							main()
{
	int						Q, dayNumber, monthCounter;
	string					operation, task;
	vector<int>				monthDay = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
	vector<vector<string>>	list(monthDay[0]);

	cin >> Q;
	monthCounter = 0;
	for (auto i = 0; i < Q; i++)
	{
		cin >> operation;
		if (operation == "ADD")
		{
			cin >> dayNumber >> task;
			list[dayNumber - 1].push_back(task);
		}
		else if (operation == "DUMP")
		{
			cin >> dayNumber;
			cout << list[dayNumber - 1].size() << " ";
			for (auto t : list[dayNumber - 1])
				cout << t << " ";
			cout << endl;
		}
		else if (operation == "NEXT")
		{
			next(list, monthDay[monthCounter] - monthDay[(monthCounter + 1) % 12]);
			monthCounter = (monthCounter + 1) % 12;
		}
	}
}
