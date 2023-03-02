package main

import "fmt"

func main() {
	narto := MakeProfile("Naruto")
	fmt.Println("Name : ", narto.Name)
	fmt.Println("Health : ", narto.Health)
	fmt.Println("Power : ", narto.Power)

	powerUp := PowerUp(narto, 2)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)
}
