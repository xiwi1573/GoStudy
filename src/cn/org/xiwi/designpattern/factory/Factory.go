package factory

import "cn/org/xiwi/designpattern/strategy"

func NetStragegy(t string) (result strategy.Strategy) {
	switch t {
	case "s":
		result = strategy.Sub{}
	case "m":
		result = strategy.Multi{}
	case "a":
		result = strategy.Add{}
	case "d":
		result = strategy.Div{}
	default:
		result = strategy.Add{}
	}
	return
}
