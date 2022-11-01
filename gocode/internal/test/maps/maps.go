package maps

import "fmt"

// DemoBase map(映射)
func DemoBase() {
	// 光声明map类型 但是没有初始化 a就是初始值nil
	var a map[string]int
	fmt.Println(a == nil)

	//map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	//map中添加键值对
	a["name"] = 100
	a["sex"] = 200
	fmt.Printf("a:%#v,T:%T\n", a, a)

	//声明map的同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v,T:%T\n", b, b)
}
