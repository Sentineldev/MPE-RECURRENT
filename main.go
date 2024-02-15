package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	learningRate = 0.5
)

func GetFloats(sample Sample) []float64 {

	var floats []float64
	for _, value := range sample.Body {
		result, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal("Can't convert input to float64", err)
		}

		floats = append(floats, result)
	}
	return floats
}

func main() {
	samples := LoadSamples()

	var network Network
	network.Network([]int64{784, 10})
	output := network.Train(GetFloats(samples[0]))

	expected := []float64{
		1,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}

	outputLayer := network.Layers[len(network.Layers)-1]
	fmt.Printf("Antes de la actualizacion\n")
	for _, layer := range network.Layers {
		for _, neuron := range layer.Neurons {
			for index, input := range neuron.Inputs {
				if index == 3 {
					break
				}
				fmt.Printf("%+v\n", input)
			}
		}
	}
	outputLayer.NodeCosts(output, expected)

	fmt.Printf("Despues de la actualizacion\n")
	for _, neuron := range outputLayer.Neurons {
		for index, input := range neuron.Inputs {
			if index == 3 {
				break
			}
			fmt.Printf("%+v\n", input)
		}
	}
}
