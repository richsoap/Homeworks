package modem

type Modem struct{}

func (m *Modem) Modulate(data []int) []float64 {
	n := len(data)
	result := make([]float64, n, n)
	for i := 0; i < n; i++ {
		if data[i] == 0 {
			result[i] = 1
		} else {
			result[i] = -1
		}
	}
	return result
}

func (m *Modem) Demodulate(data []float64) []int {
	n := len(data)
	result := make([]int, n, n)
	for i := 0; i < n; i++ {
		if data[i] < 0 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}
