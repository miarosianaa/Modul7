package main

import "fmt"

func main() {
	// Deklarasi variabel
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string
	var pertandingan int = 1

	// Input nama klub
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	fmt.Println("\nMasukkan skor untuk setiap pertandingan. Masukkan skor negatif untuk menghentikan.")

	// Loop untuk menerima skor dan menentukan pemenang
	for {
		fmt.Printf("Pertandingan %d - Skor %s: ", pertandingan, klubA)
		fmt.Scan(&skorA)
		fmt.Printf("Pertandingan %d - Skor %s: ", pertandingan, klubB)
		fmt.Scan(&skorB)

		// Kondisi untuk menghentikan program jika skor negatif
		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai.")
			break
		}

		// Menentukan hasil pertandingan
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}

		pertandingan++
	}

	// Menampilkan hasil pertandingan
	fmt.Println("\nHasil pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
}
