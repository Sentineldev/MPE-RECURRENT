package main

type NeuronInput struct {
	Value  float64
	Weight float64
}

func (input NeuronInput) WeightedInput() float64 {
	return input.Value * input.Weight
}
