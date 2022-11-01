package arrays

import "fmt"

// DemoOne
//作业求数组[1, 3, 5, 7, 8]所有元素的和
//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func DemoOne() {

	var listArray = [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range listArray {
		sum = sum + v
	}
	fmt.Println(sum)

	for k, v := range listArray {
		for i := k + 1; i < len(listArray); i++ {
			if v+listArray[i] == 8 {
				fmt.Println(k, i)
			}
		}
	}
}
