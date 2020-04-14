/*
	Формат входа. Строка s[1 . . . n], состоящая из заглавных и прописных букв латинского алфавита, 
		цифр, знаков препинания и скобок из множества []{}().

	Формат выхода. Если скобки в s расставлены правильно, выведите
		строку “Success". В противном случае выведите индекс
		(используя индексацию с единицы) первой закрывающей скобки, для
		которой нет соответствующей открывающей. Если такой нет,s
		выведите индекс первой открывающей скобки, для которой нет
		соответствующей закрывающей.
*/

package main

import (
	"container/list"
	"fmt"
)

type StackValue struct {
	index   int
	bracket rune
}

func lastBracket(bracketStack *list.List) rune {
	return bracketStack.Back().Value.(*StackValue).bracket
}

func isPairOfBrackets(openingBracket rune, closingBracket rune) bool {
	if openingBracket == '(' && closingBracket == ')' {
		return true
	} else if openingBracket == '[' && closingBracket == ']' {
		return true
	} else if openingBracket == '{' && closingBracket == '}' {
		return true
	}
	return false
}

func isBracket(elem rune) bool {
	for _, item := range "()[]{}" {
		if elem == item {
			return true
		}
	}
	return false
}

func isClosingBracket(elem rune) bool {
	for _, item := range ")]}" {
		if elem == item {
			return true
		}
	}
	return false
}

func isOpeningBracket(elem rune) bool {
	for _, item := range "([{" {
		if elem == item {
			return true
		}
	}
	return false
}

func findErrorPosition(bracketStack *list.List) {
	firstOpeningBracketIndex := 0
	closingBracketIndex := 0
	for e := bracketStack.Front(); e != nil; e = e.Next() {
		if firstOpeningBracketIndex == 0 && isOpeningBracket(e.Value.(*StackValue).bracket) {
			firstOpeningBracketIndex = e.Value.(*StackValue).index
		}
		if isClosingBracket(e.Value.(*StackValue).bracket) {
			closingBracketIndex = e.Value.(*StackValue).index
			break
		}
	}
	if closingBracketIndex != 0 {
		fmt.Println(closingBracketIndex)
	} else {
		fmt.Println(firstOpeningBracketIndex)
	}
}

func main() {
	var inputCode string
	fmt.Scan(&inputCode)

	bracketStack := list.New()
	for i, codeCh := range inputCode {
		if isBracket(codeCh) {
			if bracketStack.Len() != 0 && isPairOfBrackets(lastBracket(bracketStack), codeCh) {
				bracketStack.Remove(bracketStack.Back())
			} else {
				bracketStack.PushBack(&StackValue{i + 1, codeCh})
			}
		}
	}
	if bracketStack.Len() == 0 {
		fmt.Println("Success")
	} else {
		findErrorPosition(bracketStack)
	}
}
