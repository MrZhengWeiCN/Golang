package main

import (
	"fmt"
)

func main() {
	var countryCapitalMap map[string]string     //申明
	countryCapitalMap = make(map[string]string) //初始化

	countryCapitalMap["France"] = "法国"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	for key, value := range countryCapitalMap {
		fmt.Printf("%s 首都是： %s\n", key, value)
	}

	/*查看元素在集合中是否存在 */
	//captial, exis := countryCapitalMap["America"] /*如果确定是真实的,则存在,否则不存在 */
	/*if exis {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}*/

	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")
	//foreach
	for key, value := range countryCapitalMap {
		fmt.Printf("%s 首都是： %s\n", key, value)
	}

	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	a, ok := m[2][1]
	if !ok {
		//需要每次初始化map
		m[2] = make(map[int]string)
	}
	m[2][1] = "GOOD"
	a, ok = m[2][1]
	fmt.Println(a, ok)

	slice_map := make([]map[int]string, 5)
	for i := range slice_map {
		slice_map[i] = make(map[int]string)
		slice_map[i][2] = "value1"
		slice_map[i][1] = "value2"
		fmt.Println(slice_map[i])
	}
	fmt.Println(slice_map)
}
