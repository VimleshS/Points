//https://blog.golang.org/slices
package main

import (
	"fmt"
)

func main() {
	//	fmt.Println("...")
	slice := []int{1, 2, 3}

	//BY DEFAULT THEY ARE VALUE AND REFER TO BLOG..
	//	fmt.Printf("%.10p %v\n", &slice, slice)
	//	check_1(slice)
	//	fmt.Printf("ot %.10p %v\n", &slice, slice)

	//WHEN PASSES AS A PONITER THEN CAN MODIFY>>

	fmt.Printf("%.10p %v\n", &slice, slice)
	check_1_pointer(&slice)
	fmt.Printf("ot %.10p %v\n", &slice, slice)

	//SLICE ARE PURE VALUE TYPES>>

	//	t := slice
	//	t = append(t, 999999)
	//	fmt.Printf("%.10p %v\n", &slice, slice)
	//	fmt.Printf("%.10p %v\n", &t, t)

	//MAP ARE REFERENCE TYPES>>

	//	h := map[int]int{1: 1, 2: 2}

	//	fmt.Printf("address %p\n", &h)
	//	check_2(h)
	//	fmt.Printf("address %p  %v \n", &h, h)

}

func check_1(s []int) {
	fmt.Printf("in  %.10p %v\n", &s, s)
	s[0] = 9999999
	s = append(s, 23)
}

func check_1_pointer(s *[]int) {
	ss := *s
	fmt.Printf("in  %.10p %v\n", &s, s)
	ss[0] = 9999999
	*s = append(ss, 23)
}

func check_2(h map[int]int) {
	fmt.Printf("passes address %p\n", &h)
	h[100] = 33333
	h[5000] = 555555555555
	h[5888888888888000] = 555555555555
	delete(h, 1)
}
