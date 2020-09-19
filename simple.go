package main

import (
	"fmt"
	"strconv"
)

func impar(a int, b int) []int {
	var list []int
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			list = append(list, i)
		}
	}
	return list
}

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(impar(12, 500))
	fmt.Println(" ----------- Matriz -----------")
	var mat [][]int = toMath([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	for i := 0; i < len(mat); i++ {
		fmt.Println(mat[i])
	}
	fmt.Print("Fibo(20): " + strconv.Itoa(fibo(20)))
}

func fibo(n int) int {
	if n < 2 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func toMath(li []int, col int) [][]int {
	var mat [][]int
	if len(li)%col != 0 {
		return mat
	}
	for i := 0; i < len(li)/col; i++ {
		var fil []int
		for j := 0; j < col; j++ {
			//fmt.Print(strconv.Itoa(i) + " " + strconv.Itoa(j))
			fil = append(fil, li[i*col+j])
		}
		mat = append(mat, fil)

	}
	return mat

}
