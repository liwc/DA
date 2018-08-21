package usage

import (
	"fmt"
	"strconv"
	"DA/3_stack/stack"
)

const cnt int = 8

// 运算符集合
var opType = [cnt]string{"+", "-", "*", "/", "%", "(", ")", "#"}

// 各运算符相遇时，优先级比较 1:大于，2:小于，3:等于，0:不可能，错误
var opPrior = [cnt][cnt]int{
	{1, 1, 2, 2, 2, 2, 1, 1},
	{1, 1, 2, 2, 2, 2, 1, 1},
	{1, 1, 1, 1, 1, 2, 1, 1},
	{1, 1, 1, 1, 1, 2, 1, 1},
	{1, 1, 1, 1, 1, 2, 1, 1},
	{2, 2, 2, 2, 2, 2, 3, 0},
	{1, 1, 1, 1, 1, 0, 1, 1},
	{2, 2, 2, 2, 2, 2, 0, 3},
}

// 是否是opType中的操作符
func isOpType(str string) bool {
	for i := 0; i < cnt; i++ {
		if str == opType[i] {
			return true
		}
	}

	return false
}

// 比较优先级
func getPrior(str1, str2 string) int {
	row, col := 0, 0
	for i := 0; i < cnt; i++ {
		if str1 == opType[i] {
			row = i
		}
		if str2 == opType[i] {
			col = i
		}
	}

	return opPrior[row][col]
}

// 计算
func compute(a int, o string, b int) int {
	res := 0

	switch o {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		fallthrough
	case "%":
		if b == 0 {
			fmt.Println("error")
			return -99
		}
		if o == "/" {
			res = a / b
		} else {
			res = a % b
		}
	}

	return res
}

// ExpTest ...
func ExpTest() {
	// 操作数栈
	numStack := stack.OsInit()
	// 操作符栈
	opStack := stack.OsInit()
	opStack.Push("#")

	exp := []string{"(", "-1", "+", "(", "2", "-", "14", ")", "*", "2", "/", "-2", ")", "%", "2", "#"}
	for i := 0; i < len(exp); {
		fmt.Print("exp[i]: ")
		fmt.Println(exp[i])

		if !isOpType(exp[i]) {
			value, err := strconv.Atoi(exp[i])
			if err == nil {
				numStack.Push(value)
			}
		} else {
			top, _ := opStack.Top().(string)
			prior := getPrior(top, exp[i])
			switch prior {
			case 1:
				o, _ := opStack.Pop().(string)
				b, _ := numStack.Pop().(int)
				a, _ := numStack.Pop().(int)
				numStack.Push(compute(a, o, b))
				continue
			case 2:
				opStack.Push(exp[i])
			case 3:
				opStack.Pop()
			case 0:
				fmt.Println("error")
				break
			}
		}

		i++
		fmt.Print("numStack: ")
		numStack.Output()
		fmt.Print("opStack: ")
		opStack.Output()
	}

	fmt.Println()
	fmt.Printf("result: %v", numStack.Pop())
}
