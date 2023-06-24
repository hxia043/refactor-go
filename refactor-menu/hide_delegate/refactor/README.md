# 隐藏委托关系

如果两个类对象不需要直接访问，那么可以在类对象之间加上第三者减少耦合，也即为通过朋友，访问朋友的朋友。这是有效的，有效的前提是基于这两个类没有那么强的关系，如果两个类关系强，那就让它们直接访问，如 `remove_middle_man` 小节所述。

当然，如果两个类关系很强，强到可以绑定在一起，那就应该重构为一个类，让他们组成一个新家庭。

隐藏委托关系重构在于，当两个类关系没那么强，服务类变动不需要被客户类感知，那完全可以加入一层中介做第三方。

```
type Person struct{}

type Department struct {
	manager string
}

func (p *Person) department() Department {
	return Department{}
}

func main() {
	p := Person{}
	_ = p.department().manager
}
```

示例中，`Person` 要访问部门的经理，不过 `Person` 好像和 `Department` 没有那么强的关联关系，他不需要知道这个 `Department` 的信息，只要告诉他 `Person` 的 `manager` 是谁。  

基于这种类关系，对其进行重构：
```
func (p *Person) manager() string {
	return Department{}.manager
}

func main() {
	p := Person{}
	_ = p.manager()
}
```

在 `Person` 里加入中介 `manager`，注意中介可以是函数也可以是类。要找部门的 `manager`，直接找 `Person` 的 中介就好了。减少耦合。这也遵循迪米特法则。
