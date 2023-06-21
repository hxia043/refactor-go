# 内联函数

```
func reportLines(aCustomer Customer) []string {
	lines := []string{}
	gatherCustomerData(lines, aCustomer)
	return lines
}

func gatherCustomerData(out []string, aCustomer Customer) {
	out = append(out, "name: "+aCustomer.name)
	out = append(out, "location: "+aCustomer.location)
}
```

从 `reportLines` 看，通过调用 `gatherCustomerData` 隐藏了具体实现。好的一点是对于不需要关注实现的函数来说，隐藏这部分是有用的。不好的一点是，这部分没有返回值，意味着，不知道里面做了什么，返回了什么。这是让人有点担心的点，再细看实际上做的内容是将 `Customer` 的数据添加到列表里。那么，我们不需要通过层层调用，可以直接将这部分实现拿出来。

层层调用是基于一个函数关注一件事的基础上，如果多个函数关注一件事，反而不用层层包装，把逻辑搞复杂。
