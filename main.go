package main

import (
	"fmt"
	"math"
)

func main() {

	// Buatkan fungsi untuk memeriksa apakah sebuah bilangan adalah bilangan prima.
	fmt.Println(CheckPrimeNumber(4))
	//Buatkan fungsi untuk menentukan bilangan terbesar dari sederet bilangan dalam array berikut: $bilangan = array(11,6, 31, 201, 99, 861, 1, 7, 14, 79). Tanpa menggunakan built in fungsi array;
	fmt.Println(CheckBigestNumber([]int{
		11, 6, 31, 201, 99, 861, 1, 7, 14, 79,
	}))

	/**
	 * Buatkan fungsi yang dapat menghasilkan format berikut:
		1
		1 2
		1 2 3
		1 2 3 4
		1 2 3 4 5
		1 2 3 4 5 6
		1 2 3 4 5 6 7
		1 2 3 4 5 6 7 8
	*/
	GeneratePatern(8)

	/**
	 * Dengan menggunakan teknik bubble sorting, urutkan bilangan-bilangan berikut:
		$bilangan = array(99, 2, 64, 8, 111, 33, 65, 11, 102, 50);
	*/
	fmt.Println(BubbleSort(([]int{99, 2, 64, 8, 111, 33, 65, 11, 102, 50})))

	/**
	 * Buatkan fungsi untuk menghasilkan output berikut:
		1 5 9
		2 6 10
		3 7 11
		4 8 12
	*/
	GeneratePatern2(4, 3)

}

func CheckPrimeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	sqrt := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrt; i++ {
		// Jika n dapat dibagi habis oleh i, maka bukan bilangan prima
		if number%i == 0 {
			return false
		}
	}
	return true
}

func CheckBigestNumber(numbers []int) int {
	// Pastikan array tidak kosong
	if len(numbers) == 0 {
		return 0 // atau return nilai yang sesuai untuk kondisi Anda
	}

	largest := numbers[0]

	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}

	return largest

}

func GeneratePatern(n int) {

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func BubbleSort(arr []int) []int {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap {
			break
		}

	}
	return arr

}

func GeneratePatern2(rows, column int) {
	// book matriks
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, column)
	}
	number := 1

	for i := 0; i < column; i++ {
		for j := 0; j < rows; j++ {
			matrix[j][i] = number
			number++
		}
	}

	// print matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < column; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}
