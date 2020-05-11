#!/bin/bash

# Курс https://stepik.org/course/73/promo

# Посмотрите на функцию counter.
# Впишите в форму ниже строку, которую выведет на экран команда
# echo "counters are $c1 and $c2" если она находится в скрипте
# после десяти вызовов функции counter с параметрами сначала 1,
# затем 2, затем 3 и т.д., последний вызов с параметром 10.

counter ()  # takes one argument
{
  local let "c1+=$1"
  let "c2+=${1}*2"
}

for i in {1..10}; do
	counter i
done

echo "counters are $c1 and $c2"
