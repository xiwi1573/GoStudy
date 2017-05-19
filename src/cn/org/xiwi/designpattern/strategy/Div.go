package strategy

type Div struct {

}

func (p Div) Computer(num1, num2 int) int {

	if num2 == 0 {
		panic("num2 must not be 0!")
	}

	var res int = num1 / num2
	return res
}