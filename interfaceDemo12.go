package main

import "fmt"

type Alipay struct {
}

func (a *Alipay) CanUseFaceID() {

}

type Cash struct {
}

func (a *Cash) Stolen() {

}

type ContainCanUseFaceID interface {
	CanUseFaceID()
}

type ContainStolen interface {
	Stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case ContainCanUseFaceID:
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContainStolen:
		fmt.Printf("%T can be stolen\n", payMethod)
	}
}

func main() {
	print(new(Alipay))

	print(new(Cash))
}
