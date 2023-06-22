# 提炼变量

```
func (o *Order) getPrice() float64 {
	return o.quantity()*o.itemPrice() - math.Max(0, o.quantity()-500)*o.itemPrice()*0.05 + math.Min(o.quantity()*o.itemPrice()*0.1, 100)
}
```

我们看这段代码，最让人不理解的就是这个表达式在干嘛？
这里我们可以对每段逻辑进行拆分，赋值给变量，使得程序变得清晰易懂。

在拆分变量的时候要注意，这是类中的函数，那么可以将变量拆分到类中作为方法调用，这在对变量重复访问时将变得相当简洁。
