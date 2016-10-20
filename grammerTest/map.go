package main

import "fmt"

/*
可以使用内建函数 make 也可以使用 map 关键字来定义 Map
 */
func main() {
	//var countryCapitalMap  map[string]string
	var countryCapitalMap = make(map[string] string)

	countryCapitalMap["France"] = "Pairs"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country := range countryCapitalMap{
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if(ok){
		fmt.Println("Capital of United States is", captial)
	}else {
		fmt.Println("Capital of United States is not present")
	}

	/* 删除元素 */
	delete(countryCapitalMap,"France");
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
}
