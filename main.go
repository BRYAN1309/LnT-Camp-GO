package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var (
		nama string
		umur int
	)
	const pi = 3.14

	nama = "Bryan"
	umur = 20
	fmt.Println("nama:", nama)
	fmt.Println("umur:", umur)
	fmt.Println("pi:", pi)

	fmt.Println("===Program Input Output===")
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)
	fmt.Println(nama)

	var age int
	fmt.Print("Masukkan umur: ")
	fmt.Scanln(&age)
	fmt.Println(age)

	fmt.Println("Summary")
	fmt.Printf("Halo nama saya adalah %s dan umur saya adalah %d", nama, age)

	var numbers []int
	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	fmt.Println(numbers)

	nilai := 80
	if nilai >= 80 {
		fmt.Println("Nilai A")
	} else {
		fmt.Println("Nilai B")
	}
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	var buah = []string{"apel", "jeruk", "mangga", "pisang", "anggur"}
	for _, value := range buah {
		fmt.Println(value)
	}

	var rumahAnjing *string
	fmt.Println(rumahAnjing)

	anjing := "puppy"
	fmt.Println(anjing) // Output: <nil>

	rumahAnjing = &anjing
	fmt.Println(rumahAnjing)

	fmt.Println(*rumahAnjing) // Output: 0xc00000a080
	type Alamat struct {
		Jalan   string
		Kota    string
		KodePos string
	}
	mahasiswa := struct {
		Nama   string
		Umur   int
		Hobi   []string
		Alamat Alamat
		Aktif  bool
	}{
		Nama:  "Bryan",
		Umur:  20,
		Hobi:  []string{"membaca", "bermain game", "bersepeda"},
		Aktif: true,
		Alamat: Alamat{
			Jalan:   "jalan damai",
			Kota:    "Bandung",
			KodePos: "12345"},
	}
	fmt.Println(mahasiswa)
	fmt.Println("Nama:", mahasiswa.Nama)
	fmt.Println("Umur:", mahasiswa.Umur)
	fmt.Println("Hobi:", mahasiswa.Hobi)
	fmt.Println("Aktif:", mahasiswa.Aktif)
	fmt.Println("Alamat:", mahasiswa.Alamat.Jalan)
	fmt.Println("Kota", mahasiswa.Alamat.Kota)
	fmt.Println("Kode Pos:", mahasiswa.Alamat.KodePos)
}
