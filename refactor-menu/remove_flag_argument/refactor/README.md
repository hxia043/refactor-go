# 移除标记参数

下面是一段让人抓狂的代码：
```
fmt.Println(call("xiaoxia", true))
fmt.Println(call("xiaoxia", false))
```

试想，如果这是客户端的话。客户端在调用 `call` 函数时发现需要这么调用 `call`。当他/她看到这个 `call` 估计会和我一样纳闷：这里的 `true` 和 `false` 是几个意思？

更让他/她抓狂的是，这里的 `call` 返回的结果还不一样。

这种代码非常符合坏代码的味道。

这里的 call 在干嘛？根据传入第二个参数判断老师心情好坏，如果心情好就喊 "xiaoming"，如果心情不好就喊传入的名字。

那么，我们可以消除这一心情判断的 flag。第一步是重构 call 函数：
```
/*
Option 1: decompose call with argument happyCall and unhappyCall
*/
func happyCall() string {
	return "xiaoming"
}

func unhappyCall(name string) string {
	return name
}
```

新增函数 `happyCall` 和 `unhappyCall` 取代 call。同时客户端代码可以更新为调用 `happyCall` 和 `unhappyCall`。  
这样更加直接，更加清晰。

如果客户端调用代码和实现代码处于不同模块中，并且更改不了客户端调用代码，那么需要在原来的 `call` 函数中加上 `descope` 标注，表明后续的函数调用可直接使用 `happyCall` 和 `unhappyCall`。

既然有了 happyCall 和 unhappyCall 这种更为直接的函数声明，那么可以在条件判断函数中替换分支为函数：
```
func call(name string, isHappy bool) string {
	if isHappy {
		return happyCall()
	} else {
		return unhappyCall(name)
	}
}
```

这里也让 call 的逻辑更清晰了，实现不需要关心 happy 的时候干什么，只需要调用 `happyCall` 即可。

《重构》一节介绍用这种方式重构：
```
func happyCall(name string) string {
    return call(name, true)
}

func unhappyCall(name string) string {
    return call(name, false)
}
```

老实讲我没 get 到这么重构的精髓，为什么要这么干？既然有了直接声明函数 `happyCall` 和 `unhappyCall` 为什么不让用户直接调呢？而在内部搞一层封装？没理解。。
