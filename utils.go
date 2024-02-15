package main

import (
	"math"
	"math/rand"
	"time"
)

func Sigmoide(x float64) float64 {

	return 1 / (1 + math.Exp(-x))
}

func GenerateRandomNumber() float64 {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	value := r.Float64()

	return value
}
