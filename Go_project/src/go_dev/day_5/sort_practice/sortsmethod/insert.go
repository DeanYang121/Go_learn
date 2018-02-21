package sortsmethod

func Isort(a []int){
	for i := 1;i<len(a);i++{
		for j := i; j > 0;j-- {
			if a[j] > a[j-1]{
				break
			}
			a[j],a[j-1] = a[j-1],a[j]
		}
	}
}