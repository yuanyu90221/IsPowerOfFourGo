package power

import (
	"math"
)

func IsPowerOfFour(n int64) bool {
	if n <= 0 {
		return false
	}
	result := math.Log(float64(n)) / math.Log(4)
	return int(math.Floor(result)) == int(math.Ceil(result))
}
