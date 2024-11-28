package main

import (
	"fmt"
	"math"
)

const NMAX int = 1000

func main() {
	var N int
	fmt.Scan(&N)

	var beratAnakKelinci [NMAX]float64

	for i := 0; i < N; i++ {
		fmt.Scan(&beratAnakKelinci[i])
	}

	var terkecil, terbesar float64

	terkecil = math.MaxFloat64
	terbesar = -math.MaxFloat64

	for i := 0; i < N; i++ {
		berat := beratAnakKelinci[i]
		if berat < terkecil {
			terkecil = berat
		}
		if berat > terbesar {
			terbesar = berat
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", terkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", terbesar)
}
