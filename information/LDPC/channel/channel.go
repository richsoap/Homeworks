package channel

import (
	"math"
	"math/rand"
	"time"
)

type Channel struct {
	SNRDB float64
}

func FromDB(val float64) float64 {
	return math.Pow(10, val/10)
}

func ToDB(val float64) float64 {
	return 10 * math.Log10(val)
}

func BuildChannel(snrdb float64) Channel {
	return Channel{snrdb}
}

func (c *Channel) SetSNR(snrdb float64) {
	c.SNRDB = snrdb
}

func (c *Channel) AWGN(data []float64) []float64 {
	noisePower := 1.0 / FromDB(c.SNRDB)
	noise := Noise(noisePower, len(data))
	for i := range data {
		data[i] += noise[i]
	}
	return data
}

func Measure(data *[]float64) float64 {
	result := 0.0
	for i := range *data {
		result += ((*data)[i] * (*data)[i])
	}
	return result / float64(len(*data))
}

func Noise(power float64, length int) []float64 {
	noise := make([]float64, length, length)
	sigma := math.Sqrt(power)
	rand.Seed(time.Now().UnixNano())
	for i := range noise {
		noise[i] = rand.NormFloat64() * sigma
	}
	return noise
}
