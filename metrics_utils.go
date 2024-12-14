package main

import "encoding/json"

type Metric struct {
	size             int
	numHashFunctions int
	truePositive     float32
	falsePositive    float32
}

type Metrics struct {
	Metrics []Metric `json:"metrics"`
}

func NewMetrics() *Metrics {
	return &Metrics{
		Metrics: make([]Metric, 0),
	}
}

func (m *Metrics) Add(size, numHashFunctions int, truePositives, falsePositives float32) {
	m.Metrics = append(m.Metrics, Metric{
		size:             size,
		numHashFunctions: numHashFunctions,
		truePositive:     truePositives,
		falsePositive:    falsePositives,
	})
}

func (m *Metrics) Marshal() ([]byte, error) {
	return json.MarshalIndent(m, "", "\t")
}

func (m *Metrics) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}
