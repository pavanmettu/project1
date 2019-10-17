package main

import "fmt"

func main() {
	iarr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	k := 2
	res := circle_arr(iarr, k)
	fmt.Printf("Position Saved := %d\n", res)
}

func circle_arr(carr []int, k int) int {
	clen := len(carr)
	index := 0
	rindex := 0
	tmpk := k
	n := 0
	fmt.Println(carr)
	for {
		for index < clen {
			if carr[index] == 1 {
				if tmpk > 1 {
					tmpk--
				} else {
					carr[index] = 0
					tmpk = k
				}
			}
			index++
		}
		fmt.Println(carr)
		index = 0
		val := 0
		for i := 0; i < clen; i++ {
			val += carr[i]
		}
		if val == 1 {
			break
		}
		n++
	}
	for i := 0; i < clen; i++ {
		if carr[i] == 1 {
			rindex = i
			break
		}
	}
	return rindex + 1
}
