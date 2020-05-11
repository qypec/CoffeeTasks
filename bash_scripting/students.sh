# Напишите скрипт на bash, который принимает на вход один аргумент 
# (целое число от 0 до бесконечности), который будет обозначать число
# студентов в аудитории. В зависимости от значения числа нужно вывести разные сообщения. 
# 
# Соответствие входа и выхода должно быть таким:
# 0 -->  No students
# 1 -->  1 student
# 2 -->  2 students
# 3 -->  3 students
# 4 -->  4 students
# 5 и больше --> A lot of students

#!/bin/bash

if [[ $1 -eq 0 ]]
then
    echo "No students"
elif [[ $1 -eq 1 ]]
then
    echo "1 student"
elif [[ $1 -ge 5 ]]
then
    echo "A lot of students"
else
    echo "$1 students"
fi
