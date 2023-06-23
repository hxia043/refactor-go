# 以多态取代条件逻辑

在条件逻辑判断中，判断的层级是同级的。对于同级的逻辑可以抽取为多态实现条件逻辑判断，更清晰。

在实现多态时，通过工厂函数做进一步封装使得逻辑更加清晰：
```
func AirSpeedVelocity(bird Bird) int {
	switch bird.Kind {
	case "EuropeanSwallow":
		return NewEuropeanSwallow(bird).airSpeedVelocity()
	case "AfricanSwallow":
		return NewAfricanSwallow(bird).airSpeedVelocity()
	case "NorwegianBlueParrot":
		return NewNorwegianBlueParrot(bird).airSpeedVelocity()
	default:
		return 0
	}
}
```

示例中，不够清晰，对象的构建过程是可见的。使用工厂函数做进一步封装：
```
func amountFor(performanceData map[string]interface{}) int {
	playType, ok := playFor(performanceData)["type"].(string)
	if !ok {
		return 0
	}

	if calculator := performance.CreatePerformanceCalculator(performanceData, playType); calculator != nil {
		return calculator.Amount()
	}

	return 0
}
```

通过工厂函数做一层封装，清晰不少，不要暴露那么多。
