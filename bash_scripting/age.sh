# Курс https://stepik.org/course/73/promo

# Напишите скрипт на bash, который будет определять в какую возрастную группу попадают пользователи. При
# запуске скрипт должен вывести сообщение "enter your name:" и ждать от пользователя ввода имени
# (используйте read, чтобы прочитать его). Когда имя введено, то скрипт должен написать "enter your age:"
# и ждать ввода возраста (опять нужен read). Когда возраст введен, скрипт пишет на экран "<Имя>, your
# group is <группа>", где <группа> определяется на основе возраста по следующим правилам:
# 	младше либо равно 16: "child",
# 	от 17 до 25 (включительно): "youth",
# 	старше 25: "adult".
# После этого скрипт опять выводит сообщение "enter your name:" и всё начинается по новой. Если в
# какой-то момент работы скрипта будет введено пустое имя или возраст 0, то скрипт должен написать на
# экран "bye" и закончить свою работу.

#!/bin/bash

check_arguments () # nothing 
{
	echo "enter your name:"; read name
	if [[ $name == "" ]]; then
		echo "bye"; exit
	fi

	echo "enter your age:"; read age
	if [[ $age -eq 0 ]]; then
		echo "bye"; exit
	fi
}

name=""
age=0
while true ; do
	check_arguments 
    if [[ $age -le 16 ]]; then
        group="child"
    elif [[ $age -le 25 && $age -ge 17 ]]; then
        group="youth"
    else
        group="adult"
    fi
    echo "$name, your group is $group"
done
