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
}
