# 以卫语句取代嵌套条件表达式

在 `decompose conditional` 一节中介绍了如何清晰重构判断逻辑。对于判断逻辑来说，当判断的分支有两种，相等和不等。

相等意味着重要性，对于相等分支，如果分支过多，可使用 switch 匹配，亦或用 `if... else if ... else 结构`。并对其中的判断函数，分支函数做封装，整体层次是相当清晰的。

对于不等，意味着某些分支的优先级要高于其它分支，那么这些分支应该以卫语句的方式呈现。卫语句结构更加清晰（其位置隐含了重要性这一点）。

在分支的结构里，也可以通过反转条件手法使得逻辑更加简洁，清晰。比如：
```
func adjustedCapital(anInstrument Instrument) int {
	result := 0
	if anInstrument.capital > 0 {
		if anInstrument.interestRate > 0 && anInstrument.duration > 0 {
			result = (anInstrument.income / anInstrument.duration) * anInstrument.adjustmentFactor
		}
	}

	return result
}
```

这个示例中，嵌套的条件判断如果用条件反转，合并条件表达式整合判断逻辑之后会清晰不少：
```
func adjustedCapital(anInstrument Instrument) int {
	result := 0

	if anInstrument.capital <= 0 || anInstrument.interestRate <= 0 || anInstrument.duration <= 0 {
		return result
	}

	return (anInstrument.income / anInstrument.duration) * anInstrument.adjustmentFactor
}
```

当然，这里可以将判断逻辑提炼成函数，更加清晰简洁。

有意思的是，这个示例中 result 不是承担了一种责任。如果满足条件是一种 result，如果不满足是另一种 result。可以直接返回，去掉 result 来解脱它身上的责任。
