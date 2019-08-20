package lib

import "fmt"

// 打印自动机的内容
func (dfa *DFA) Println() {
	dfa.Root.Println("$")
}

// 递归打印状态
func (state *State) Println(payload string) {
	for k, v := range *state {
		if len(*v) != 0 {
			v.Println(payload + "->" + k)
		} else {
			fmt.Println(payload + "->" + k)
		}
	}
}
