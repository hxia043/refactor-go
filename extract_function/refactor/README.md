# 提炼函数

这里将 `printOwing()` 内代码根据要做的事情拆分为三个函数： `printBanner()`，`recordDateTo(invoice)` 和 `printDetails(invoice)`。

需要注意的是，拆分之后的函数虽只用 `invoice` 的部分属性，但给函数传递完整的结构体对象是有必要的。

同时，这里隐藏了 `calculateAmount(invoice)` 实现，而不是如《重构》示例中，将 `calculatexxx` 拿出来：
```
function printOwing(invoice) {
  printBanner();
  const outstanding = calculateOutstanding(invoice);
  recordDueDate(invoice);
  printDetails(invoice, outstanding);
}
```

个人觉得，这种区别取决于意图和实现。《重构》中的意图是想展示 outstanding 的计算过程，并且打印返回值。本例中，是想隐藏实际的计算过程，只关注于最终的结果。

