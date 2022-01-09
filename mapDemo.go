package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["American"]

	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	countryCapitalMap1 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	for country := range countryCapitalMap1 {
		fmt.Println(country, "首都是", countryCapitalMap1[country])
	}

	delete(countryCapitalMap1, "France")

	fmt.Println("删除后地图")

	for country := range countryCapitalMap1 {
		fmt.Println(country, "首都是", countryCapitalMap1[country])
	}
}
