package strategy

type Multi struct {

}

func (p Multi) Computer(num1, num2 int) int {
	var res int = num1 * num2
	return res
}