package sortsmethod

func Bsort(a []int){

	for i := 0 ; i < len(a);i++{
		for j := 1 ; j < len(a) - i ; j++{
			if(a[j] > a[j-1]){
				a[j] , a[j-1] = a[j-1],a[j]
			}
		}
	}

}
