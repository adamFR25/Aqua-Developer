package main

import "fmt"

func main() {
	n := [...]int{50000, 25000, 75000, 10000, 20000, 20000, 5000, 30000, 10000, 10000}
	x := [...]string{
		"Benih lele", "Pakan lele Cap Menara", "Probiotik A", "Probiotik Nila B", "Pakan Nila",
		"Benih Nila", "Cupang", "Benih Nila", "Benih Cupang", "Probiotik B",
	}

	isDone := false

	for !isDone {
		isDone = true
		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				x[i], x[i+1] = x[i+1], x[i]
				isDone = false
			}
		}
	}

	fmt.Println("=====Pembelian 100K=====")
	var total int = 0
	for i := 0; i < len(n); i++ {
		if total >= 100000 {
			break
		}
		fmt.Println(x[i], ":", n[i])
		total = n[i] + total

	}
	fmt.Println("Total : ", total)

	fmt.Println("====Termurah & Termahal=====")
	fmt.Println("Termurah :", x[0])
	fmt.Println("Termahal :", x[len(x)-1])

	sepuluh(x, n)

}

func sepuluh(nama [10]string, harga [10]int) {

	fmt.Println("======Barang 10k=========")
	for i := 0; i < len(nama); i++ {
		if harga[i] == 10000 {
			fmt.Println("Barang 10K : ", nama[i])
		}
	}

}
