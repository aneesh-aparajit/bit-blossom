package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func main() {
	exists, nonExists := 500, 1000
	datasetExists, datasetNonExists := make([]string, 0), make([]string, 0)

	for i := 0; i < exists; i++ {
		datasetExists = append(datasetExists, uuid.NewString())
	}

	for i := 0; i < nonExists; i++ {
		datasetNonExists = append(datasetNonExists, uuid.NewString())
	}

	metrics := NewMetrics()

	for numHashFunctions := 1; numHashFunctions <= 20; numHashFunctions++ {
		for size := 100; size <= 10000; size += 100 {
			bloom := NewBloomFilter(size, numHashFunctions)
			for _, item := range datasetExists {
				bloom.Add(item)
			}

			truePositive, falsePositive := 0, 0
			for _, item := range datasetExists {
				if bloom.Exists(item) {
					truePositive++
				}
			}
			for _, item := range datasetNonExists {
				if bloom.Exists(item) {
					falsePositive++
				}
			}
			truePositivePercentage := (float32(truePositive) / float32(exists)) * 100
			falsePositivePercentage := (float32(falsePositive) / float32(nonExists)) * 100
			fmt.Println("[numHashFunctions:", numHashFunctions, "| size:", size, "]:\t", "truePositivePercentage:", truePositivePercentage, "% \tfalsePositivePercentage:", falsePositivePercentage, "%")
			metrics.Add(size, numHashFunctions, truePositivePercentage, falsePositivePercentage)
		}
	}

	v, err := metrics.Marshal()
	if err != nil {
		log.Panic(err.Error())
	}
	f, err := os.Create("metrics.json")
	if err != nil {
		log.Panic("couldn't create file")
	}
	if _, err := f.Write(v); err != nil {
		log.Panicln(err.Error())
	}
}
