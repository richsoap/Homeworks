package channel

import (
	"fmt"
	"math"
	"testing"
)

func approx(a, b float64) bool {
	return math.Abs(a-b) < 0.1
}

func checkResult(t *testing.T, name string, from, get, wanted float64) {
	if !approx(get, wanted) {
		t.Error(fmt.Sprintf("func: %v, input %v, get %v, wanted %v", name, from, get, wanted))
	}
}

func TestDBFunc(t *testing.T) {
	raws := []float64{1, 2, 3, 4, 5, 6, 7}
	dbs := []float64{0, 3, 4.8, 6, 7, 7.8, 8.5}
	for i := 0; i < len(raws); i++ {
		checkResult(t, "ToDB", raws[i], ToDB(raws[i]), dbs[i])
		checkResult(t, "FromDB", dbs[i], FromDB(dbs[i]), raws[i])
	}
}
