# 将值对象改为引用对象

引用类型和值类型的区别在于唯一性，类是唯一的吗？如果是那就引用类型，如果不是可考虑值类型。

在示例中：
```
type customer struct {
	name string
	menu [10]string
}

func (c customer) order(name string) {
	for index, food := range c.menu {
		if food == "" {
			food = name
			c.menu[index] = food
			break
		}
	}
}

func main() {
	c := customer{name: "hxia"}
	c.order("cola chicken")
}
```

`customer` 类调用 `order` 点单，点单的菜品写到 `menu`，这里为防止 `menu` 执行引用类型复制，将 menu 设成数组用来存储菜品。

调用的结果显示 menu 为空：
```
{hxia [         ]}
```

这是因为方法调用执行的是对象值类型复制，而不是引用类型，函数 order 内的更新并不会作用到外层的 `customer` 对象。

将值方法调用改为引用类型调用：
```
func (c *customer) order(name string) {
	for index, food := range c.menu {
		if food == "" {
			food = name
			c.menu[index] = food
			break
		}
	}
}
```

发现改动作用在外层 `customer` 对象上：
```
{hxia [cola chicken         ]}
```

# 将引用对象改为值对象

构造值对象场景:
```
func New(name string) customer {
	return customer{name: name}
}
```

这里的工厂函数负责构造 `customer` 对象。这里函数返回的是值类型。

在这个示例中也可以返回引用类型没有问题。

为什么改值类型，进入 `New` 函数栈，在栈上构造 `customer` 结构体，如果函数返回的是引用类型，栈上的结构体数据会逃逸到堆上。而且，`customer` 类本身并不占多少空间，值类型复制并不会消耗太多空间。因此，这里考虑用值类型取代引用类型。


综上，可以看出。用值类型和引用类型取决于场景，有的场景需要值类型，有的场景需要引用类型，有的场景可能都行。这没问题，有问题的是在需要值类型的时候用了引用类型，需要引用类型的时候用了值类型，这是重构要做的，甚至在开发的时候就需要做了（否则，bug 会找上门。。）
