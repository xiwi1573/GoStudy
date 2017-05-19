package strategy

type Sub struct {

}

func (p Sub) Computer(num1,num2 int) int  {
	var res int = num1-num2
	return res
}
