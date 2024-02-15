package main

type Layer struct {
	Neurons []Neuron
	Inputs  []float64
}

func (layer *Layer) Layer(numOfNeurons int64) {
	var neurons []Neuron
	for i := 0; i < int(numOfNeurons); i++ {
		var neuron Neuron
		neuron.Neuron(
			Sigmoide,
			1.00,
		)
		neurons = append(neurons, neuron)
	}
	layer.Neurons = neurons
}
func (layer *Layer) CalculateOutput() []float64 {
	var output []float64
	var neurons []Neuron
	for _, neuron := range layer.Neurons {
		neuron.SetInputs(layer.Inputs)
		neurons = append(neurons, neuron)
		output = append(output, neuron.CalculateOutput())
	}

	layer.Neurons = neurons

	return output
}

func (layer Layer) NodeCost(outputActivation float64, expectedOutput float64) float64 {

	value := outputActivation - expectedOutput
	return value * value
}

func (layer Layer) NodeCosts(output []float64, expectedOutput []float64) {
	var costs []float64
	for index, value := range output {
		costs = append(costs, layer.NodeCost(value, expectedOutput[index]))
	}

	for index, neuron := range layer.Neurons {
		neuron.UpdateWeights(costs[index])
	}
}
