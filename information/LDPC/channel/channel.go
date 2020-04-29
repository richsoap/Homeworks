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
	signalPower := Measure(&data)
	noisePower := signalPower / FromDB(c.SNRDB)
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
	for i := range noise {
		noise[i] = NormRand(0, power)
	}
	return noise
}

func NormRand(average, sig float64) float64 {
	s := 0.0
	v1 := 0.0
	v2 := 0.0
	rand.Seed(time.Now().UnixNano())
	for s >= 1 || s == 0 {
		v1 = 2*rand.Float64() - 1
		v2 = 2*rand.Float64() - 1
		s = v1*v1 + v2*v2
	}
	return v1*math.Sqrt(-2*math.Log(s)/s)*math.Sqrt(sig) + average
}
