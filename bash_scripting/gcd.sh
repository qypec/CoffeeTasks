#!/bin/bash

# Курс https://stepik.org/course/73/promo

# Напишите скрипт на bash, который будет искать наибольший общий делитель 
# (НОД, greatest common divisor, GCD) двух чисел. При запуске ваш скрипт 
# ждет ввода двух натуральных чисел через пробел. После ввода чисел скрипт
# считает их НОД и выводит на экран сообщение "GCD is <посчитанное значение>".
# После этого скрипт опять входит в режим ожидания двух натуральных чисел. 
# Если в какой-то момент работы пользователь ввел вместо этого пустую строку,
# то нужно написать на экран "bye" и закончить свою работу. 

gcd() # $a $b
{
	if [[ $a -eq $b ]]; then
		result=$a; return
	elif [[ $a -gt $b ]]; then
		let "a -= b"
		gcd $a $b
	else
		let "b -= a"
		gcd $a $b
	fi
}

read_arguments () # nothing 
{
	read a b
	if [[ $a == "" || $b == "" ]]; then
		echo "bye"; exit
	fi
}

result="empty"
while true; do
	read_arguments
	gcd $a $b
	if [[ $result != "empty" ]]; then
		echo "GCD is $result"
		result="empty"
	fi
done
