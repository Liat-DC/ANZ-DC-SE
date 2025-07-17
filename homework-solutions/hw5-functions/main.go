// This program prints a triangle in 4 different ways

package main

import (
	"fmt"
)

func print_line(l int, c string) {
	for i := 0; i < l; i++ {
		fmt.Print(c)
	}
}

func print_triangle_1(b int, c string) {
	fmt.Println()
	for i := 1; i <= b; i++ {
		print_line(i, c)
		fmt.Println()
	}
}

func print_triangle_2(b int, c string) {
	fmt.Println()
	for i := b; i >= 1; i-- {
		print_line(i, c)
		fmt.Println()
	}
}

func print_triangle_3(b int, c string) {
	fmt.Println()
	for i := 1; i <= b; i++ {
		print_line(b-i, " ")
		print_line(i, c)
		fmt.Println()
	}
}

func print_triangle_4(b int, c string) {
	fmt.Println()
	for i := 1; i <= b; i += 2 {
		print_line((b-i)/2, " ")
		print_line(i, c)
		fmt.Println()
	}
}

func main() {

	var base int
	var char string

	fmt.Println("Please enter base length and a single character:")
	fmt.Scan(&base, &char)

	print_triangle_1(base, char)
	print_triangle_2(base, char)
	print_triangle_3(base, char)
	print_triangle_4(base, char)

}
