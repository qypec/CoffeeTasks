#!/bin/bash

# Курс https://stepik.org/course/73/promo

# Напишите калькулятор на bash. При запуске ваш скрипт должен ожидать ввода пользователем
# команды (при этом на экран выводить ничего не нужно). Команды могут быть трех типов: 
#   1) Слово "exit". В этом случае скрипт должен вывести на экран слово "bye" и завершить работу. 
#   2) Три аргумента через пробел -- первый операнд (целое число), операция 
#	(одна из "+", "-", "*", "/", "%", "**") и второй операнд (целое число). В этом случае нужно
# 	произвести указанную операцию над заданными числами и вывести результат на экран. После этого
# 	переходим в режим ожидания новой команды.
#   3) Любая другая команда из одного аргумента или из трех аргументов, но с операцией не из списка.
# 	В этом случае нужно вывести на экран слово "error" и завершить работу.

read_arguments () # nothing 
{
	read a operation b
	if [[ $a == "exit" ]]; then
		echo "bye"; exit
	fi
	if [[ $a == "" || $operation == "" || $b == "" ]]; then
		echo "error"; exit
	fi
}

while true; do
	read_arguments
	case $operation in
		"+")
			let "result = a + b"
			echo $result
			;;
		"-")
			let "result = a - b"
			echo $result
			;;
		"*")
			let "result = a * b"
			echo $result
			;;
		"**")
			let "result = a ** b"
			echo $result
			;;
		"/")
			let "result = a / b"
			echo $result
			;;
		"%")
			let "result = a % b"
			echo $result
			;;
		*)
			echo "error"; exit
	esac
done
