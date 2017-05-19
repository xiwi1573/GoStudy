package strategy

type Add struct {

}

func (p Add) Computer(num1,num2 int) int  {
	var res int = num1+num2
	return res
}
