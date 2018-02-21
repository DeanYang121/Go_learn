package buffIo

import (
	"io"
	"fmt"
	"os"
	"bufio"
)

type CharCount struct{
	chCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func TestBufio2(){
	file,err := os.Open("l:/go-project/src/go_dev/test.log")
	if err != nil {
		fmt.Println("read file err:",err)
		return
	}

	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)
	for{
		str ,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}

		if err != nil {
			fmt.Println("read string failed,err:",err)
			break
		}
		runeArr := []rune(str)
		for _,v := range runeArr{
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.chCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++ 
			}
		}
	}

	fmt.Printf("char count:%d\n",count.chCount)
	fmt.Printf("num count:%d\n",count.NumCount)
	fmt.Printf("space count:%d\n",count.SpaceCount)
	fmt.Printf("other count:%d\n",count.OtherCount)
	// fmt.Printf("read str succ,return :%s \n",str)
}




