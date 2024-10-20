# 分解条件表达式

在 `introduce_parameter_object` 中，重构的示例下有这么一段代码：  
```
if v < r.min || v > r.max {
    fmt.Println(v)
}
```

这段代码在条件判断中是可以被重构的，不一样的是这个示例还算清晰，且没有其它分支就保留了。

对于较为复杂的条件判断，重构就显得有必要了。如《重构》中的示例：
```
if !data.isBefore(plan.summerStart) && !data.isAfter(plan.summerEnd) {
    charge = quantity * plan.summerRate
} else {
    charge = quantity * plan.regularRate * plan.regularServiceCharge
}
```

这段代码里我们需要思考，需要看它在干嘛？原来是对于不是 summer 和 summer 做区分。那完全可以通过提炼函数，提炼出判断逻辑，这一好处在于清晰。

接着，对判断逻辑的各个分支应用提炼函数，if 的逻辑应该是相对的，提炼于分支函数，即可使得分支判断看的更清爽，干净。

注意，这里由于 `Javascript` 和 `Go` 语言层面的不同，这里的判断和分支函数被提取为对象 Plan 的方法。并且，`quantity` 总是和 plan 搭配在一起，进一步将 `quantity` 提到类中。

对于分解条件表达式来说属于形变实不变。
