package main

import "fmt"

const NMAX int = 1000

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var beratIkan [NMAX]float64
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var wadahBerat []float64
	totalIkan := 0

	for i := 0; i < x; i += y {
		var wadahTotal float64
		for j := i; j < i+y && j < x; j++ {
			wadahTotal += beratIkan[j]
		}
		wadahBerat = append(wadahBerat, wadahTotal)
		totalIkan++
	}

	for _, total := range wadahBerat {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	var totalBerat float64
	for _, total := range wadahBerat {
		totalBerat += total
	}
	rataRata := totalBerat / float64(totalIkan)
	fmt.Printf("%.2f\n", rataRata)
}
