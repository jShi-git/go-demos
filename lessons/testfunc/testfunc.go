package main

import . "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func myfunc(arg ...int) {
	for _, n := range arg {
		Printf("And the number is: %s\n", n)
	}
}

func main() {
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) //调用函数max(x, y)
	max_xz := max(x, z) //调用函数max(x, z)

	Printf("max(%d, %d) = %d\n", x, y, max_xy)
	Printf("max(%d, %d) = %d\n", x, z, max_xz)
	Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // 也可在这直接调用它

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	Printf("%d + %d = %d\n", x, y, xPLUSy)
	Printf("%d * %d = %d\n", x, y, xTIMESy)

	myfunc(1, 2, 3)

	for i := 0; i < 5; i++ {
		defer Printf("%d ", i)
	}
}
