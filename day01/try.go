package main

import(
	"fmt"
	"day01/calic"
)
func main(){
	var a int
	a = 1
	b := 2
	fmt.Printf("a = %d\nb = %d\n",a,b)

	arry := [5]int{1,2,3,4,5}
	for i,value := range arry{
		fmt.Printf("arry[%d] = %d , ",i,value)
	}
	//var 切片变量 = make([]类型， 长度, 容量)

	var sli1 = make([]int,4,8)
	var sli2 = []int{1,2,3,4,5}
	fmt.Println("\nsli1 = ",sli1)
	fmt.Println("sli2 = ",sli2)
	sli1 = append(sli1,1,2,3,4)
	fmt.Println("sli1 = ",sli1)
	
	// 仅声明

	m1 := make(map[string]string)

	m1["open"] = "close"
	fmt.Println(m1["open"])
	m2 := map[string]int{
		"math":5,
		"chinses":6,
	}
	for key,value := range m2{
		fmt.Printf("%s : %d\n",key,value)
	}
	str1 := "lin"
	var p *string = &str1
	*p = "shuai"
	fmt.Printf("str1 = %s",str1)
	fmt.Printf("add = %d",calic.Addnum(a,b))

}
