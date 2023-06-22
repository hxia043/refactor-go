# 内联变量

```
func main() {
	price := getPrice()
	fmt.Println("price: ", price)
}
```

这段代码已经告诉我们 `getPrice` 获取到的是 price，在将其赋给 `price` 变量则会显得累赘，不够清晰。

可以使用内联变量取决局部变量，使得逻辑更清晰。
