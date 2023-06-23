package performance

type Performance interface {
	Amount() int
	VolumeCredits() float64
}

func CreatePerformanceCalculator(performanceData map[string]interface{}, playType string) Performance {
	switch playType {
	case "comedy":
		return NewComedy(performanceData)
	case "tragedy":
		return NewTragedy(performanceData)
	default:
		return nil
	}
}
