package argtextparse

import (
	"testing"
)

func TestBasic(t *testing.T) {
	tests := []testTemplate{
		{"simple", "hello", ArgumentCollection{Sink: "hello", ShortArg: nil, LongArg: nil}},
		{"just arg", "--hello", ArgumentCollection{Sink: "", ShortArg: nil, LongArg: map[string]ArgumentValue{"hello": newArgData("", 1)}}},
		{"simple arg", "hello --foo bar", ArgumentCollection{Sink: "hello", ShortArg: nil, LongArg: map[string]ArgumentValue{
			"foo": newArgData("bar", 1),
		}}},
		{"join sink", "hello - world", ArgumentCollection{Sink: "hello - world", ShortArg: nil, LongArg: nil}},
		{"joined sinkless", "--hello world - hey --hi hello", ArgumentCollection{Sink: "", ShortArg: nil, LongArg: map[string]ArgumentValue{
			"hello": newArgData("world - hey", 1),
			"hi":    newArgData("hello", 1),
		}}},
		{"counter", "-vvvvv", ArgumentCollection{Sink: "", ShortArg: map[string]ArgumentValue{"v": newArgData("", 5)}, LongArg: nil}},
		{"sink counter", "hello -vvvvv", ArgumentCollection{Sink: "hello", ShortArg: map[string]ArgumentValue{"v": newArgData("", 5)}, LongArg: nil}},
		{"sink counter", "hello -hi o/ -vvvvv hi", ArgumentCollection{Sink: "hello", ShortArg: map[string]ArgumentValue{
			"v": newArgData("hi", 5),
			"h": newArgData("o/", 1),
			"i": newArgData("o/", 1),
		}, LongArg: nil}},
		{"sinkless", "-hello wor - ld", ArgumentCollection{Sink: "", ShortArg: map[string]ArgumentValue{
			"h": newArgData("wor - ld", 1),
			"e": newArgData("wor - ld", 1),
			"l": newArgData("wor - ld", 2),
			"o": newArgData("wor - ld", 1),
		}, LongArg: nil}},
		{"all", "hello -wo r - ld --hello world -vv", ArgumentCollection{Sink: "hello", ShortArg: map[string]ArgumentValue{
			"w": newArgData("r - ld", 1),
			"o": newArgData("r - ld", 1),
			"v": newArgData("", 2),
		}, LongArg: map[string]ArgumentValue{
			"hello": newArgData("world", 1),
		}}},
	}

	for _, test := range tests {
		t.Run(t.Name(), testHelper(test))
	}
}

func testHelper(test testTemplate) func(*testing.T) {
	return func(t *testing.T) {
		p := Parse(test.input)
		if !isEqual(p, test.output) {
			t.Errorf("Test#%s\nInput:  [%v]\nReturn: %v\nWant:   %v", test.name, test.input, p, test.output)
		}
	}
}

type testTemplate struct {
	name   string
	input  string
	output ArgumentCollection
}

func isEqual(a, b ArgumentCollection) bool {
	if a.Sink != b.Sink {
		return false
	}
	if !isArgDataEqual(a.ShortArg, b.ShortArg) {
		return false
	}
	if !isArgDataEqual(a.LongArg, b.LongArg) {
		return false
	}
	return true
}
func isArgDataEqual(a, b map[string]ArgumentValue) bool {
	if len(a) != len(b) {
		return false
	}
	for ai, av := range a {
		bv, be := b[ai]
		if !be {
			return false
		}
		if av != bv {
			return false
		}
	}
	return true
}
func newArgData(val string, ctr int) ArgumentValue {
	return ArgumentValue{
		Value:   val,
		Flag:    true,
		Counter: ctr,
	}
}

/*func isEqualStr(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
*/
