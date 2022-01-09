package main

import "fmt"

func main() {
	// var n [10]int
	// var i, j int

	// for i = 0; i < 10; i++ {
	// 	n[i] = i + 100
	// }

	// for j = 0; j < 10; j++ {
	// 	fmt.Printf("Element[%d] = %d\n", j, n[j])
	// }

	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i := 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j := 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k := 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	// listDemo2()

	var balance4 = [5]int{1000, 2, 3, 17, 50}
	var avg float32
	avg = listDemo3(balance4, 5)
	fmt.Println(avg)
}

func listDemo2() {
	values := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	values = append(values, row1)
	values = append(values, row2)

	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	fmt.Printf("第一个元素为: %d\n", values[0][0])

	a := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11}}

	sites := [2][2]string{}

	sites[0][0] = "Google"
	sites[0][1] = "Runppb"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"

	fmt.Println(a)

	fmt.Println(sites)

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func listDemo3(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
