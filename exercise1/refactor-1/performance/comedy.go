package performance

import (
	"fmt"
	"math"
)

type comedy struct {
	performance map[string]interface{}
}

func (t *comedy) Amount() int {
	result := 30000
	audience, ok := t.performance["audience"].(int)
	if !ok {
		return 0
	}

	if audience > 20 {
		result += 10000 + 500*(audience-20)
	}
	result += 300 * audience

	return result
}

func (t *comedy) VolumeCredits() float64 {
	result := 0.0
	audience, ok := t.performance["audience"].(int)
	if !ok {
		fmt.Println("no audience in performance")
		return 0
	}

	result += math.Max(float64(audience-30), 0) + math.Floor(float64(audience)/5)
	return result
}

func NewComedy(performance map[string]interface{}) *comedy {
	return &comedy{
		performance: performance,
	}
}
