package lib

import (
	"strings"
)

// 状态
type State map[string]*State

// 有穷自动机
type DFA struct {
	Root   *State //开始状态
	Cursor *State //当前状态
	Cache  string //敏感字缓存
}

// 通过敏感关键字数组生成自动机
func NewDFA(words []string) DFA {
	state := &State{}
	for i := 0; i < len(words); i++ {
		letters := strings.Split(words[i], "")
		cursor := state
		for j := 0; j < len(letters); j++ {
			if _, ok := (*cursor)[letters[j]]; !ok {
				(*cursor)[letters[j]] = &State{}
			}
			cursor = (*cursor)[letters[j]]
		}
	}
	return DFA{Root: state, Cursor: state, Cache: ""}
}

// 该状态是否是终结状态
func (state *State) isEnd() bool {
	return len(*state) == 0
}

// 重置状态机
func (dfa *DFA) Reset() {
	dfa.Cursor = dfa.Root
	dfa.Cache = ""
}

// 自动机开始逐字读取
func (dfa *DFA) Read(word string) (hit string) {
	if state, ok := (*dfa.Cursor)[word]; ok {
		dfa.Cursor = state
		dfa.Cache += word
		if dfa.Cursor.isEnd() {
			hit = dfa.Cache
			dfa.Reset()
		}
	} else {
		dfa.Reset()
	}
	return
}

// 通过自动机分析文本
func (dfa *DFA) Analyze(content string) []string {
	words := make([]string, 0)
	list := strings.Split(content, "")
	for index := 0; index < len(list); index++ {
		if hit := dfa.Read(list[index]); hit != "" {
			words = append(words, hit)
		}
	}
	return words
}
