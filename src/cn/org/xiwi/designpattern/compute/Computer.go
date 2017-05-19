package compute

import (
	"fmt"

	s "cn/org/xiwi/designpattern/strategy"
)

type Computer struct {
	Num1, Num2 int
	strate     s.Strategy
}

func (p *Computer) SetStrategy(strate s.Strategy) {
	p.strate = strate
}

func (p Computer) Do() int {

	defer func() {
		if f := recover(); f != nil {
			fmt.Println(f)
		}
	}()

	if p.strate == nil {
		panic("strategy is null")
	}

	return p.strate.Computer(p.Num1, p.Num2)
}
