package control

import(
	"fmt"
	"math/rand"
)

func Guess_number(){
	var n int
	n = rand.Intn(100)

	for{
		var input int
		fmt.Scanf("%d\n",&input)
		flag := false
		switch {
		case input == n:
			fmt.Println("yes you are right!")
			flag = true
		case input > n:
			fmt.Println("your number is larger")
		case input < n :
			fmt.Println("your number is smaller")
		}

		if(flag){
			break
		}
	}
}