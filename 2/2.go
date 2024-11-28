package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung rata-rata
func rataRata(arr []int) float64 {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi
func standarDeviasi(arr []int) float64 {
	mean := rataRata(arr)
	var sumSquares float64
	for _, val := range arr {
		sumSquares += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(sumSquares / float64(len(arr)))
}

// Fungsi untuk menghitung frekuensi elemen
func frekuensi(arr []int, elemen int) int {
	count := 0
	for _, val := range arr {
		if val == elemen {
			count++
		}
	}
	return count
}

func main() {
	var N, x, index, elemen int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	arr := make([]int, N)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:", arr)

	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan bilangan kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := x - 1; i < N; i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < N {
		arr = append(arr[:index], arr[index+1:]...)
	}
	fmt.Println("Isi array setelah dihapus:", arr)

	fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata(arr))
	fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi(arr))

	fmt.Print("Masukkan elemen untuk menghitung frekuensinya: ")
	fmt.Scan(&elemen)
	fmt.Printf("Frekuensi elemen %d: %d kali\n", elemen, frekuensi(arr, elemen))
}
