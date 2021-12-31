package main

import (
	"bytes"
	"fmt"
)

func joinString(slist ...string) string {
	var b bytes.Buffer

	for _, s := range slist {
		b.WriteString(s)
	}

	return b.String()
}

func printTypeValue(slist ...interface{}) string {
	var b bytes.Buffer

	for _, s := range slist {
		str := fmt.Sprintf("%v", s)

		var typeString string

		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}

		b.WriteString("value: ")

		b.WriteString(str)

		b.WriteString(" type: ")

		b.WriteString(typeString)

		b.WriteString("\n")
	}
	return b.String()
}

func rawPrint(rawList ...interface{}) {
	for _, a := range rawList {
		fmt.Println(a)
	}
}

func print(slist ...interface{}) {
	// rawPrint(slist...)
	rawPrint("fmt", slist)
}

func main() {
	fmt.Println(joinString("pig ", "and ", "rat"))
	fmt.Println(joinString("hammer ", "mom ", "and ", "hawk"))

	fmt.Println(printTypeValue(100, "str", true))

	print(1, 2, 3)
}
