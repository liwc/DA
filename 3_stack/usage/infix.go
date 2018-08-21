package usage

import (
	"fmt"
	"strconv"
	"DA/3_stack/stack"
)

func reverse(src string) string {
	runes := []rune(src)
	for from, to := 0, len(runes) - 1; from < to; from, to = from + 1, to - 1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// infix to prefix
func in2pre(infix []string) string {
	tmpStk1 := stack.OsInit()
	tmpStk2 := stack.OsInit()

	for i := len(infix) - 1; i >= 0; {
		if !isOpType(infix[i]) {
			value, err := strconv.Atoi(infix[i])
			if err == nil {
				tmpStk2.Push(value)
			}
		} else {
			if tmpStk1.Empty() || infix[i] == ")" {
				// 如果S1为空或infix[i]为右括号“)”，则直接将此运算符入栈
				tmpStk1.Push(infix[i])
			} else {
				top, _ := tmpStk1.Top().(string)
				if infix[i] == "(" {
					if top == ")" {
						tmpStk1.Pop()
					} else {
						tmpStk2.Push(tmpStk1.Pop())
						fmt.Print("tmpStk1: ")
						tmpStk1.Output()
						fmt.Print("tmpStk2: ")
						tmpStk2.Output()
						continue
					}
				} else if top == ")" {
					// 如果顶运算符为右括号“)”，则直接将此运算符入栈
					tmpStk1.Push(infix[i])
				} else {
					prior := getPrior(infix[i], top)
					if prior == 1 || prior == 3 {
						// 若优先级比栈顶运算符的较高或相等，也将运算符压入S1
						tmpStk1.Push(infix[i])
					} else {
						// 将S1栈顶的运算符弹出并压入到S2中，再次与S1中新的栈顶运算符相比较
						tmpStk2.Push(tmpStk1.Pop())
						fmt.Print("tmpStk1: ")
						tmpStk1.Output()
						fmt.Print("tmpStk2: ")
						tmpStk2.Output()
						continue
					}
				}
			}
		}

		i-- 
		fmt.Print("tmpStk1: ")
		tmpStk1.Output()
		fmt.Print("tmpStk2: ")
		tmpStk2.Output()
	}

	for !tmpStk1.Empty() {
		tmpStk2.Push(tmpStk1.Pop())
	}

	fmt.Print("tmpStk1: ")
	tmpStk1.Output()
	fmt.Print("tmpStk2: ")
	tmpStk2.Output()

	var res string
	for !tmpStk2.Empty() {
		value := tmpStk2.Pop()
		value1, ok1 := value.(int)
		if ok1 {
			res += strconv.Itoa(value1)
		} else {
			value2, ok2 := value.(string)
			if ok2 {
				res += value2
			}
		}
	}

	return res
}

func parsePre(prefix string) int {
	tmpStk := stack.OsInit()

	for i := len(prefix) - 1; i >= 0; i-- {
		if !isOpType(string(prefix[i])) {
			value, _ := strconv.Atoi(string(prefix[i]))
			tmpStk.Push(value)
		} else {
			a, _ := tmpStk.Pop().(int)
			b, _ := tmpStk.Pop().(int)
			c := compute(a, string(prefix[i]), b)
			tmpStk.Push(c)
		}

		fmt.Print("tmpStk: ")
		tmpStk.Output()
	}

	res, _ := tmpStk.Pop().(int)
	return res
}

// infix to postfix
func in2post(infix []string) string {
	tmpStk1 := stack.OsInit()
	tmpStk2 := stack.OsInit()

	for i := 0; i < len(infix); {
		if !isOpType(infix[i]) {
			value, err := strconv.Atoi(infix[i])
			if err == nil {
				tmpStk2.Push(value)
			}
		} else {
			if tmpStk1.Empty() || infix[i] == "(" {
				tmpStk1.Push(infix[i])
			} else {
				top, _ := tmpStk1.Top().(string)
				if infix[i] == ")" {
					if top == "(" {
						tmpStk1.Pop()
					} else {
						tmpStk2.Push(tmpStk1.Pop())
						fmt.Print("tmpStk1: ")
						tmpStk1.Output()
						fmt.Print("tmpStk2: ")
						tmpStk2.Output()
						continue
					}
				} else if top == "(" {
					tmpStk1.Push(infix[i])
				} else {
					prior := getPrior(infix[i], top)
					if prior == 1 {
						// 若优先级比栈顶运算符的较高或相等，也将运算符压入S1
						tmpStk1.Push(infix[i])
					} else {
						// 将S1栈顶的运算符弹出并压入到S2中，再次与S1中新的栈顶运算符相比较
						tmpStk2.Push(tmpStk1.Pop())
						fmt.Print("tmpStk1: ")
						tmpStk1.Output()
						fmt.Print("tmpStk2: ")
						tmpStk2.Output()
						continue
					}
				}
			}
		}

		i++ 
		fmt.Print("tmpStk1: ")
		tmpStk1.Output()
		fmt.Print("tmpStk2: ")
		tmpStk2.Output()
	}

	for !tmpStk1.Empty() {
		tmpStk2.Push(tmpStk1.Pop())
	}

	fmt.Print("tmpStk1: ")
	tmpStk1.Output()
	fmt.Print("tmpStk2: ")
	tmpStk2.Output()

	var res string
	for !tmpStk2.Empty() {
		value := tmpStk2.Pop()
		value1, ok1 := value.(int)
		if ok1 {
			res += strconv.Itoa(value1)
		} else {
			value2, ok2 := value.(string)
			if ok2 {
				res += value2
			}
		}
	}

	return reverse(res)
}

func parsePost(postfix string) int {
	tmpStk := stack.OsInit()

	for i := 0; i < len(postfix); i++ {
		if !isOpType(string(postfix[i])) {
			value, _ := strconv.Atoi(string(postfix[i]))
			tmpStk.Push(value)
		} else {
			b, _ := tmpStk.Pop().(int)
			a, _ := tmpStk.Pop().(int)
			c := compute(a, string(postfix[i]), b)
			tmpStk.Push(c)
		}

		fmt.Print("tmpStk: ")
		tmpStk.Output()
	}

	res, _ := tmpStk.Pop().(int)
	return res
}

// InfixTest ...
func InfixTest() {
	infix := []string{"(", "1", "+", "2", ")", "*", "3", "-", "4"}
	prefix := in2pre(infix)
	fmt.Printf("prefix: %s\n", prefix)
	fmt.Printf("result: %d\n", parsePre(prefix))

	postfix := in2post(infix)
	fmt.Printf("prefix: %s\n", postfix)
	fmt.Printf("result: %d\n", parsePost(postfix))

	fmt.Println()
}