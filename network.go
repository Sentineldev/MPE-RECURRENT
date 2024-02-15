package main

type Network struct {
	Layers []Layer
}

func (network *Network) Network(layerNeuron []int64) {

	var layers []Layer

	for _, numOfNeurons := range layerNeuron {
		var layer Layer
		layer.Layer(numOfNeurons)
		layers = append(layers, layer)
	}
	network.Layers = layers
}

func (network *Network) Train(Inputs []float64) []float64 {
	current := Inputs
	var layers []Layer
	for index, layer := range network.Layers {
		if index == 0 {
			continue
		}
		layer.Inputs = current
		current = layer.CalculateOutput()
		layers = append(layers, layer)
	}

	network.Layers = layers
	return current
}
