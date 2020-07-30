package argtextparse

const terminator = "-"

func main() {
	//parseDebug("Test - lol -h el lo -eeee k --hello world")
	//parseDebug("-hello wor - ld")
	//parseDebug("-vv")
	//parseDebug("--hello world - hey --hello world")
}

func Parse(cli string) ArgumentCollection {

	var sink string
	var shortArg = make(map[string]ArgumentValue)
	var longArg = make(map[string]ArgumentValue)

	var focus focusMode

	var tmpSKeys []string //maybe make them share the same storage, but explode it on saving?
	var tmpLKey string
	var tmpValues string

	endWildcard := false
	tMode := tNone
	for i, r := range cli {
		if (i == 0 || isSpace(safeArrayIndex(cli, i-1)) || hasTMode(tMode)) && string(r) == terminator && !isSpace(safeArrayIndex(cli, i+1)) {
			endWildcard = true
			focus = fKey //resets
			if tMode == tNone {
				tMode = tShort
			} else if tMode == tShort {
				tMode = tLong
			}
			continue
		} else if !endWildcard {
			if isSpace(string(r)) && safeArrayIndex(cli, i+1) == terminator && !isSpace(safeArrayIndex(cli, i+2)) {
				endWildcard = true
				continue
			}
			sink += string(r)
			continue
		}
		switch tMode {
		case tShort:
			if isSpace(string(r)) && safeArrayIndex(cli, i+1) == terminator && !isSpace(safeArrayIndex(cli, i+2)) || len(cli) == i+1 {
				if !isSpace(string(r)) {
					switch focus {
					case fKey:
						tmpSKeys = append(tmpSKeys, string(r))
					case fValue:
						tmpValues += string(r)
					}
				}
				for _, sk := range tmpSKeys {
					sArg, isSet := shortArg[sk]
					if !isSet {
						sArg = ArgumentValue{}
					}
					sArg.Value = tmpValues
					sArg.Flag = true
					sArg.Counter++
					shortArg[sk] = sArg
				}
				tmpValues = ""
				tmpSKeys = []string{}
				tMode = tNone
				continue
			}
			if focus == fKey {
				if isSpace(string(r)) {
					focus = fValue
					continue
				}
				tmpSKeys = append(tmpSKeys, string(r))
			}
			if focus == fValue {
				tmpValues += string(r)
			}
			continue
		case tLong:
			if isSpace(string(r)) && safeArrayIndex(cli, i+1) == terminator && !isSpace(safeArrayIndex(cli, i+2)) || len(cli) == i+1 {
				if !isSpace(string(r)) {
					switch focus {
					case fKey:
						tmpLKey += string(r)
					case fValue:
						tmpValues += string(r)
					}
				}

				lArg, isSet := longArg[tmpLKey]
				if !isSet {
					lArg = ArgumentValue{}
				}
				lArg.Value = tmpValues
				lArg.Flag = true
				lArg.Counter++
				longArg[tmpLKey] = lArg

				tmpValues = ""
				tmpLKey = ""
				tMode = tNone
				continue
			}
			if focus == fKey {
				if isSpace(string(r)) {
					focus = fValue
					continue
				}
				tmpLKey += string(r)
			}
			if focus == fValue {
				tmpValues += string(r)
			}
			continue
		}
	}
	return ArgumentCollection{
		Sink:     sink,
		ShortArg: shortArg,
		LongArg:  longArg,
	}
}
