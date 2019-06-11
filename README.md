## 笔记说明 

由于`Go`语言语法相对简单，笔记中伪代码都以Go版本为演示，读者不具备Go语言基础也能看懂代码。  

[source](https://github.com/overnote/Algorithm/tree/master/sources)文件夹下拥有各类语言版本的实现案例，未来会不断涵盖Java，JavaScript，TypeScript，C，C++等各个版本。  

为了贴心，笔者列出了其他语言使用者不易理解的Go语法部分。 

如下是一个普通的函数`NewLineList()`：
```go
package list                                            // 包名，一般与文件夹名一致
func NewList(size int) *List {	                        // 名称大写代表公有
	return &List{
		MaxSize: size,
		Length: 0,
		Data: make([]interface{}, size),
	}
}
```

函数的调用：使用 包名.函数名 方式
```go
r := list.NewLineList(3)				// := 是变量类型自动推导
```

如下是go中的实例方法`IsEmpty()`：
```go
func (l *List) IsEmpty() bool {		                // `l`是实例对象的形参，相当于其他语言的`this`，`self`。
        if l.Length == 0 {
		return true
	}
	return false
}
```

go中实例方法的调用：
```go
r := list.NewList(3)                                    // 先使用上述第一节的函数生产一个实例对象
b := r.IsEmpty()                                        // 调用实例方法
```


## 算法资料

#### 数据结构

- [《大话数据结构》](https://book.douban.com/subject/6424904/)：C++编写，适合完全0基础人员对数据结构进行大致了解，书籍有不少代码错误。
- [《学习JavaScript数据结构与算法》第3版](https://book.douban.com/subject/26639401/)：JS编写，适合完全0基础人员对数据结构进行大致了解。
- [《数据结构》(邓俊辉)](https://book.douban.com/subject/25859528/)：C++编写，数据结构书籍集大成者，简洁，详细，深入，笔者力荐。

#### 算法入门

- [《算法图解》](https://book.douban.com/subject/26979890/)：图文并茂的算法书籍
- [《算法新解》](https://book.douban.com/subject/26931430/)：言简意赅，不错的算法入门书
- [《算法笔记》](https://book.douban.com/subject/26827295/)：涵盖了大部分常见算法，是难得的算法入门笔记
- [《算法设计与分析基础》第3版](https://book.douban.com/subject/26337727/)：Java书写的经典算法入门书籍
- [《数据结构与算法分析-**语言描述》](https://book.douban.com/subject/1139426/)：算法书籍的入门经典，笔者推荐C或者Java语言版
- [《编程珠玑》](https://book.douban.com/subject/3227098/)：为算法提供了精辟的解题思路，是算法思想学习的瑰宝

#### 算法进阶

- [《算法》](https://book.douban.com/subject/10432347/)：经典书籍，笔者认为最好的算法书籍，Java编写
- [《算法导论》](https://book.douban.com/subject/1885170/)：恢弘巨著，算法领域的代表作
- [《计算机程序设计艺术》](https://book.douban.com/subject/1130500/)：恢弘巨作，算法领域的里程碑

#### 算法习题

- [《编程之美》](https://book.douban.com/subject/3004255/)：微软面试指南集合
- [《算法竞赛入门经典》刘汝佳·第2版](https://book.douban.com/subject/25902102/)：算法题佳作之一
- [《算法竞赛进阶指南》李煜东 ](https://book.douban.com/subject/30136932/)：算法题佳作之一
- [《程序员代码面试指南》](https://book.douban.com/subject/26638586/)：面试指南之一
- [《剑指offer》](https://book.douban.com/subject/27008702/)：面试指南之一

#### 学习网址

- [牛客网](https://www.nowcoder.com/)：面向基础与面试的算法题库
- [力扣](https://leetcode.com/)：著名的算法题网站
- [算法导论](http://open.163.com/special/opencourse/algorithms.html)：麻省理工学院公开课
- [JULY博主](https://blog.csdn.net/v_july_v)：算法讲解私人博客
- [oRblt专栏](https://blog.csdn.net/orbit)：算法讲解私人博客

## 附录：博主自己的其他笔记

[推荐书籍](https://github.com/ruyuejun/polaris)：综合了所有技术对应的书籍与书评 https://github.com/ruyuejun/polaris  

[知识笔记汇总仓库](https://github.com/overnote)：所有技术笔记所在组织 https://github.com/overnote   

对上述仓库的分类索引：  
- [Server](https://github.com/overnote/server)：包含常用服务端技术：Nginx、mysql、redis、mongodb、linux系统等
- [数据结构与算法](https://github.com/overnote/algorithm)：数据结构与算法笔记，主讲Go版本，另附JS，Java版本，未来可能增加C版
- [JavaScript](https://github.com/overnote/javascript)：包含网页、JS、前端工程化、Node、vue、react等JS相关领域技术
- [Go](https://github.com/overnote/golang)：详尽的Go领域笔记，包括语法、并发编程、web编程、微服务等
- [Java](https://github.com/overnote/java)：整理中，Java篇幅过大，整理困难，可能会鸽
- [Python](https://github.com/overnote/python)：整理中
- [计算机考研](https://github.com/overnote/postgraduate)：整理中
