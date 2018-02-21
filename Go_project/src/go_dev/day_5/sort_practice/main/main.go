package main
import(
	"fmt"
	"go_dev/day_5/sort_practice/sortsmethod"
)

func main(){
	b := [...]int{8,12,10,5,4,3,15}

	// sortsmethod.Bsort(b[:])
	sortsmethod.Isort(b[:])
	fmt.Println(b)

}

