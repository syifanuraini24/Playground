// Sort array terlebih dahulu, kemudian rotasi ke kiri sesuai dengan nilai yang telah ditentukan.
//
// Contoh Sort array:
// Input: [4,5,2,1,3]
// Output: [1,2,3,4,5]

//Contoh RotateLeft:
//Input: 4, [1,2,3,4,5]
//Output: [5,1,2,3,4]

// Explanation RotateLeft:
// untuk melakukan rotasi kiri dengan nilai 4, array mengalami urutan perubahan berikut:
// [1,2,3,4,5] -> [2,3,4,5,1] -> [3,4,5,1,2] -> [4,5,1,2,3] -> [5,1,2,3,4]

package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{4, 5, 2, 1, 3}
	arrSorted := Sort(nums)
	fmt.Println(arrSorted)
	rotateLeft := RotateLeft(2, arrSorted)
	fmt.Println(rotateLeft)
}

func Sort(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	return arrCopy
}

func RotateLeft(d int, arr []int) []int {
	if d < 1 || len(arr) < 1 {
		return arr
	}
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	loop := d % len(arr)
	for i := 0; i < loop; i++ {
		first := copyArr[0]
		copyArr = append(copyArr[1:], first)
	}
	return copyArr
}
