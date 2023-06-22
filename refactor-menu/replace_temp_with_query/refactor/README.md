# 查询取代变量

查询取代变量，取代的是什么变量？

```
以查询取代临时变量（178）手法只适用于处理某些类型的临时变量：那些只被计算一次且之后不再被修改的变量。最简单的情况是，这个临时变量只被赋值一次，但在更复杂的代码片段里，变量也可能被多次赋值——此时应该将这些计算代码一并提炼到查询函数中。并且，待提炼的逻辑多次计算同样的变量时，应该能得到相同的结果。 - 《重构》
```

分析 `getPrice` 的函数，对于函数来说，它关注于一件事，并且足够短小精悍。可以说，不用重构了。
```
func (b *bill) getPrice() float64 {
	var basePrice = b._quantity * b._item.price
	var discountFactor = 0.98
	if basePrice > 1000 {
		discountFactor -= 0.03
	}

	return float64(basePrice) * discountFactor
}
```

不过，这里的临时变量可以进一步被重构，提炼成两个处理 basePrice 和 discountFactor 的函数用于返回局部变量。  

最终，得到重构后的函数：
```
func (b *bill) getPrice() float64 {
	return b.basePrice() * b.discountFactor()
}
```

可以发现，函数更加清晰明了，值得重构。
