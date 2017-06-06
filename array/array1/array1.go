package main

import "fmt"
import "reflect"

func main() {
	var arr [5]int
	fmt.Println(reflect.TypeOf(arr)) //[5]int
	var sli []int
	fmt.Println(reflect.TypeOf(sli)) //[]int

	arr1 := [1]int{100}
	arr2 := [2]int{1, 2}
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr, "\n", arr1, "\n", arr2, "\n", arr3) //[0 0 0 0 0],[100],[1 2],[1 2 3 4 5]

	fmt.Println("二维数组")

	arr4 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr4) //[[1 2 3] [4 5 6]]
	arr5 := [2][3]string{{"jack", "lily", "lucy"}, {"张三", "李四", "王五"}}
	fmt.Println(arr5)       //[[jack lily lucy] [张三 李四 王五]]
	fmt.Println(arr5[0][2]) //lucy
	fmt.Println(arr5[1][0]) //张三

	fmt.Println("多维数组")
	//数组定义，接受不同类型
	arr6 := [2][3][4]interface{}{
		{
			{1, 2, 3, 4},
			{"21", 22.22, false, 24},
			{31, 32.0, 33, nil}},
		{
			{41, 42, "43", 44},
			{51, "jack", 53, 54},
			{61, 62, nil, 64}}}
	fmt.Println(arr6)
	fmt.Println(arr6[0:1])
	fmt.Println(arr6[0:1][1:2])

	fmt.Println("数组的比较")
	arr10 := [3]int{1, 2, 3}
	arr11 := [3]int{1, 2, 3}
	arr12 := [...]int{1, 2, 3}
	sli13 := []int{1, 2, 3} //slice
	fmt.Println(sli13)      //[1 2 3]

	fmt.Println(arr10 == arr11) //true
	fmt.Println(arr10 == arr12) //true
	// fmt.Println(arr10 == sli13) ERRIR mismatched types [3]int and []int 数组不能和slice比较
}
