package buffIo


import (
	"fmt"
	"os"
	"bufio"
)

func TestBufio(){
	reader := bufio.NewReader(os.Stdin)
	str,err :=reader.ReadString('\n') //byte类型 使用单引号
	if err != nil{
		fmt.Println("read string failed...\n")
		return
	}

	fmt.Printf("read str success,ret:%s\n",str)
}




