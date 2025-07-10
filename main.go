package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")

	// var (
	// 	nama string
	// 	umur int
	// )
	// const pi = 3.14

	// nama = "Bryan"
	// umur = 20
	// fmt.Println("nama:", nama)
	// fmt.Println("umur:", umur)
	// fmt.Println("pi:", pi)

	// fmt.Println("===Program Input Output===")
	// var nama string
	// fmt.Print("Masukkan nama: ")
	// fmt.Scanln(&nama)
	// fmt.Println(nama)

	// var age int
	// fmt.Print("Masukkan umur: ")
	// fmt.Scanln(&age)
	// fmt.Println(age)

	// fmt.Println("Summary")
	// fmt.Printf("Halo nama saya adalah %s dan umur saya adalah %d", nama, age)

	// var numbers []int
	// numbers = append(numbers, 10)
	// numbers = append(numbers, 20)
	// fmt.Println(numbers)

	// nilai := 80
	// if nilai >= 80 {
	// 	fmt.Println("Nilai A")
	// } else {
	// 	fmt.Println("Nilai B")
	// }
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	var buah = []string{"apel", "jeruk", "mangga", "pisang", "anggur"}
	for _, value := range buah {
		fmt.Println(value)
	}
}
