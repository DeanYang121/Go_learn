package control


import(
	"fmt"
)

//go 的switch不需要break 不会一直往下走 fallTrought 会一直往下走
func SwitchA(){
	var a int = 10
	switch a{
	case 0:
		fmt.Println("a is equal 0")
		fallthrough
	case 10: 
		fmt.Println("a is equal 10")
	default:
		fmt.Println("you input a error number")
	}
}