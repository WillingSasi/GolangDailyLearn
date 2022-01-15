package main

import (
	"fmt"
	"math/big"
	"time"
)

const LIM = 1000

var fibs [LIM]*big.Int

func main() {
	// big1 := new(big.Int).SetUint64(uint64(1000))
	// fmt.Println("big1 is: ", big1)

	// big2 := big1.Uint64()
	// fmt.Println("big2 is: ", big2)

	// big3, _ := new(big.Int).SetString("1000", 10)
	// fmt.Println("big3 is: ", big3)

	// big4 := big3.Uint64()
	// fmt.Println("big4 is: ", big4)

	result := big.NewInt(0)
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("数列第 %d 位： %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行完成，所耗时间为： %s\n", delta)
}

func fibonacci(n int) (res *big.Int) {
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(fibs[n-1], fibs[n-2])
	}

	fibs[n] = res
	return
}
