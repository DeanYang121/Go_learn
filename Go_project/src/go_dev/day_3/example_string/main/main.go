package main

import(
	"fmt"
	"strings"
	"go_dev/day_3/example_string/strtests"
)

func urlProcess(url string)string{
	result := strings.HasPrefix(url,"http://")
	if !result {
		url = fmt.Sprintf("http://%s",url)
	}
	return url
}

func pathProcess(path string)string{
	result := strings.HasSuffix(path,"/")
	if !result {
		path = fmt.Sprintf("%s/",path)
	}
	return path
}
//strings.Replace(,,,)  "heheheworld","he","we",2
//strings.Count(str string,"he")int
//strings.Repeat()
//strings.ToLoweer
//strings.Join()
//strings.Split
//string.TrimSpace()
func main(){
// 	var (
// 		url string
// 		path string
// 	)
// 	fmt.Println("please input: url path")
// 	fmt.Scanf("%s%s",&url,&path)
// 	url = urlProcess(url)
// 	path = pathProcess(path)

// 	fmt.Println(url,path)
	strtests.Strlearn()
}