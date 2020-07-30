package argtextparse

type ArgumentCollection struct {
	Sink     string
	ShortArg map[string]ArgumentValue
	LongArg  map[string]ArgumentValue
}

func (data ArgumentCollection) GetPair(short string, long string) (arg ArgumentValue, success bool) {
	arg, success = data.ShortArg[short]
	if success {
		return arg, success
	}
	arg, success = data.LongArg[long]
	if success {
		return arg, success
	}
	return ArgumentValue{}, false
}

type ArgumentValue struct {
	Value   string
	Flag    bool
	Counter int
}

type terminatorMode int

const (
	tNone terminatorMode = iota
	tShort
	tLong
)

type focusMode int

const (
	fKey focusMode = iota
	fValue
)
