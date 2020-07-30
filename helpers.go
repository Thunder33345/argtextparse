package argtextparse

import "fmt"

func safeArrayIndex(str string, index int) string {
	if index >= 0 && index < len(str) {
		return string(str[index])
	}
	return ""
}

func hasTMode(mode terminatorMode) bool {
	return mode != terminatorMode(tNone)
}

func isSpace(r string) bool {
	switch r {
	//case ' ', '\t', '\r', '\n':
	//	return true
	case " ", "\t", "\r", "\n":
		return true
	}
	return false
}

func parseDebug(input string) {
	full := Parse(input)
	str := full.Sink
	short := full.ShortArg
	long := full.LongArg

	fmt.Printf("\ninput:%s\n", input)
	fmt.Printf("str:%s\n", str)
	fmt.Printf("Short:%v\n", short)
	for i, s := range short {
		fmt.Printf("Key:%s, value:%s, flag:%t, counter:%d\n", i, s.Value, s.Flag, s.Counter)
	}
	fmt.Printf("Long:%v\n", long)
	for i, s := range long {
		fmt.Printf("Key:%s, value:%s, flag:%t, counter:%d\n", i, s.Value, s.Flag, s.Counter)
	}
}
