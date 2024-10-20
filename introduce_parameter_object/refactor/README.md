# 引入参数对象

在 `combine_function_into_class` 中介绍了，当函数组团做事时，可以将这且函数组合成类。这里数据对于这些泥团函数来说重要亦不重要。这里的重点在于这些泥团函数绑定在一些，需要做一组事情。

可以将这些泥团函数提炼成包装函数，通过调用包装函数来调用泥团函数。也可以将这些泥团函数提炼成类方法，提炼成类方法的好处在于，数据和行为绑定到类中，高内聚了。

那么，这里讨论到数据。对于数据来说，如果一组数据成团出现，也可以将这些数据封装为对象。可以传递对象，更为重要的一点是，后续，使用到这些对象的函数可以被改造数据对象的方法。  

当然，这里的改造需要注意的是细粒度的问题，如示例:
```
func readingsOutsideRange(station station, r Range) {
	for _, reading := range station.readings {
		r.contain(reading.temp)
	}
}
```

改造的是提炼方法 `contian` 用来判断数据是否在 range 之间，而不是提炼的 `readingsOutsideRange`，如果用 `r.readingsOutsideRange` 那这个改造就过度了。变成了高耦合了，耦合的是 `station` 对象。
