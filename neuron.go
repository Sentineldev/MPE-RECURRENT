package main

type Neuron struct {
	Inputs   []NeuronInput
	Function ActivationFunc
	Bias     float64
	Cost     float64
}

func (neuron *Neuron) Neuron(function ActivationFunc, bias float64) {
	neuron.Function = function
	neuron.Bias = bias
}

func (neuron *Neuron) SetInputs(values []float64) {

	var inputs []NeuronInput
	if len(neuron.Inputs) == 0 {
		neuron.Bias = GenerateRandomNumber()
		for _, input := range values {
			inputs = append(inputs, NeuronInput{
				Value:  input,
				Weight: GenerateRandomNumber(),
			})
		}
	} else {
		for index, input := range values {
			inputs = append(inputs, NeuronInput{
				Value:  input,
				Weight: neuron.Inputs[index].Weight,
			})
		}
	}
	neuron.Inputs = inputs
}

func (neuron Neuron) PonderateSum() float64 {
	var sum float64
	for _, input := range neuron.Inputs {
		sum += input.WeightedInput()
	}
	return sum + neuron.Bias
}
func (neuron Neuron) Activation() float64 {
	sum := neuron.PonderateSum()

	return neuron.Function(sum)
}
func (neuron Neuron) CalculateOutput() float64 {
	return neuron.Activation()
}
func (neuron *Neuron) UpdateWeights(cost float64) {
	var inputs []NeuronInput
	for _, input := range neuron.Inputs {
		newWeight := (cost * input.Value * learningRate) + input.Weight
		inputs = append(inputs, NeuronInput{
			Weight: newWeight,
		})
	}
	neuron.Inputs = inputs
}
