package main

import (
	"fmt"
)

func hitungHariBerikutnya(N int, frekuensi []int) int {
	lcm := 1
	for i := 0; i < N; i++ {
		lcm = (lcm * frekuensi[i]) / gcd(lcm, frekuensi[i])
	}
	return lcm
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Membaca jumlah pedagang
	var N int
	fmt.Scan(&N)

	// Membaca frekuensi kunjungan pedagang
	frekuensi := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&frekuensi[i])
	}

	// Menghitung jumlah hari berikutnya untuk Pasar Rakyat
	hasil := hitungHariBerikutnya(N, frekuensi)

	// Menampilkan hasil
	fmt.Println(hasil)
}
