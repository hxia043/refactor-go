# 以命令取代函数和以函数取代命令

函数和类的关系相当于什么？个人理解，函数相当于机器，可以做任务的机器，把一件任务做好，只做一件任务的机器。类是对象，是人，是使用工具的对象。

在工程里，函数过多，都是机器，没有对象操作不好，容易混乱。都是对象，没有机器，干活的少了。

怎么平衡函数和类呢？  
个人以为，如果只关注一件事，且不需要过多的参数，对象参与，可以将这件事交给函数实现。如果事情和事情之间有联系，可以组合函数成类对象，将这些做事情的函数提炼成类方法。  
类尽量短小精悍，这符合开闭原则，高内聚，低耦合，没问题。如果类只做一件事，且并不复杂，那么可以将这样的类转换为函数，更为简洁。

这就是《重构》里以命令对象取代函数，以函数取代命令对象的适用场合。
