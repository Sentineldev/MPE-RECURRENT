package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Sample struct {
	Label string
	Body  []string
}

func (sample Sample) String() string {
	return fmt.Sprintf("Etiqueta: %s\nTamaño de la muestra: %d", sample.Label, len(sample.Body))
}

func LoadSamples() []Sample {

	file, err := os.Open("mnist_train.csv")

	var samples []Sample

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error while reading records", err)
	}

	for _, eachrecord := range records {

		label := eachrecord[0:1]
		sample := eachrecord[1:]

		samples = append(samples, Sample{
			Label: label[0],
			Body:  sample,
		})
	}

	return samples[1:]
}
