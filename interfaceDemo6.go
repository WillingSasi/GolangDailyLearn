package main

import (
	"fmt"
	"sort"
)

// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool
// 	Swap(i, j int)
// }

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	name := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Sort(name)

	for _, v := range name {
		fmt.Printf("%s\n", v)
	}

	names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Sort(names)

	for _, v := range names {
		fmt.Printf("%s\n", v)
	}

	namess := []string{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Strings(namess)

	for _, v := range namess {
		fmt.Printf("%s\n", v)
	}
}
