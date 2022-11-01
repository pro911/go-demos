package slices

import (
	"fmt"
	"sort"
)

func DemoBase() {
	//基于数组的到切片
	a := [5]int{1, 2, 3, 4, 5} //数组
	b := a[1:4]                //切片  切片key1:key2 前包含后舍去  2 3 4
	fmt.Printf("T:%T,b:%v\n", b, b)

	//切片再次切片
	c := b[0:len(b)]
	fmt.Printf("T:%T,c:%v\n", c, c)

	// make函数构造切片
	d := make([]int, 5, 10) //长度是5 容量是10的切片。
	fmt.Printf("T:%T,d:%v\n", d, d)
	fmt.Printf("len(d):%v,cap(d):%v\n", len(d), cap(d))

	//切片不能直接比较 只能与nil比较。
	var e []int
	if e == nil {
		//[] == nil
		fmt.Println("e是一个nil")
	}

	var f = []int{}
	if f == nil {
		fmt.Println("f是一个nil")
	}
	fmt.Println(f, len(f), cap(f))

	g := make([]int, 0, 0)
	if g == nil {
		fmt.Println("g是一个nil")
	}
	fmt.Println(g, len(g), cap(g))
	//判断一个切片是否为空 通常要用len来判断 而不是和nil比较。

	//切片的赋值拷贝
	h := make([]int, 3) //如果不传入容量值 默认容量=len
	i := h
	i[2] = 100
	fmt.Println(h, i)

	//切片的遍历
	j := []int{1, 2, 3, 4, 5}
	for x := 0; x < len(j); x++ {
		fmt.Printf("k:%v,v:%v\n", x, j[x])
	}
	fmt.Println()
	for index, value := range j {
		fmt.Printf("index:%v,value:%v\n", index, value)
	}

	//切片声明后需要初始化后才能被调用。
	var k []int //此时它并没有申请内存 a == nil 是真
	for x := 0; x < 10; x++ {
		k = append(k, x)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", k, len(k), cap(k), k)
	}
}

func DemoOne() {
	var a []string
	var b []int

	var c = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func DemoTwo() {
	var a = [...]int{3, 2, 8, 7, 9, 6, 4}
	sort.Ints(a[:])
	fmt.Println(a)
}
