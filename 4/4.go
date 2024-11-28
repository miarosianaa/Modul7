package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array dengan input dari user
func isiArray(t *tabel, n *int) {
	var char rune
	fmt.Println("Masukkan karakter (akhiri dengan '.'): ")
	for *n = 0; *n < NMAX; *n++ {
		fmt.Scanf("%c", &char)
		if char == '.' {
			break
		}
		t[*n] = char
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan urutan isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Fungsi untuk memeriksa apakah isi array membentuk palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// Mengisi array tab dengan memanggil prosedur isiArray
	isiArray(&tab, &m)

	// Mencetak isi array sebelum dibalik
	fmt.Print("Isi array sebelum dibalik: ")
	cetakArray(tab, m)

	// Membalikkan isi array tab dengan memanggil balikanArray
	balikanArray(&tab, m)

	// Mencetak isi array setelah dibalik
	fmt.Print("Isi array setelah dibalik: ")
	cetakArray(tab, m)

	// Memeriksa apakah isi array membentuk palindrom
	if palindrom(tab, m) {
		fmt.Println("Array membentuk palindrom.")
	} else {
		fmt.Println("Array tidak membentuk palindrom.")
	}
}
