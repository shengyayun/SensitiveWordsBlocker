package main

import (
	"fmt"

	"./lib"
)

func main() {
	words := []string{"敏感词", "敏感语句", "屏蔽字"}
	dfa := lib.NewDFA(words)
	fmt.Println(dfa.Analyze("测试敏感词与屏蔽字"))
}
