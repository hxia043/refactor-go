package performance

import (
	"fmt"
	"math"
)

type tragedy struct {
	performance map[string]interface{}
}

func (t *tragedy) Amount() int {
	result := 40000
	audience, ok := t.performance["audience"].(int)
	if !ok {
		return 0
	}

	if audience > 30 {
		result += 1000 * (audience - 30)
	}

	return result
}

func (t *tragedy) VolumeCredits() float64 {
	result := 0.0

	audience, ok := t.performance["audience"].(int)
	if !ok {
		fmt.Println("no audience in performance")
		return 0
	}

	result += math.Max(float64(audience-30), 0)
	return result
}

func NewTragedy(performance map[string]interface{}) *tragedy {
	return &tragedy{
		performance: performance,
	}
}
