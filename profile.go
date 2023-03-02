package main

import (
	"math/rand"
	"time"
)

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profile {
	// Array of Values
	healths := []int{50, 100, 150, 200, 250, 300, 350, 400}
	powers := []int{50, 100, 150, 200, 250, 300, 350, 400}
	exps := []int{0, 50, 100, 150, 200, 250, 300, 350}

	// Init global pseudo random generator
	rand.Seed(time.Now().Unix())

	return Profile{
		Name:   name,
		Health: randValues(healths),
		Power:  randValues(powers),
		Exp:    randValues(exps),
	}
}

// Rand values from slices
func randValues(slices []int) int {
	return slices[rand.Intn(len(slices))]
}
