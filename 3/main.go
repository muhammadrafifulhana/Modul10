package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64, jumlah int) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 0; i < jumlah; i++ {
		berat := arrBerat[i]
		if berat < *bMin {
			*bMin = berat
		}
		if berat > *bMax {
			*bMax = berat
		}
	}
}

func rerata(arrBerat arrBalita, jumlah int) float64 {
	var total float64
	var count int

	for i := 0; i < jumlah; i++ {
		berat := arrBerat[i]
		if berat != 0 {
			total += berat
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return total / float64(count)
}

func main() {
	var N int
	var arr arrBalita

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	var bMin, bMax float64

	hitungMinMax(arr, &bMin, &bMax, N)

	avg := rerata(arr, N)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", avg)
}
