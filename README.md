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

函数的调用：
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

推荐书籍：
```
数据结构
        《学习JavaScript数据结构与算法》第3版            - 入门书籍：使用JavaScript编写
        《数据结构》(邓俊辉)                            - 入门书籍：使用C++编写  
        《数据结构与算法分析》                           - 经典书籍：推荐C，Java版本

算法初级
        《算法图解》                                    - 入门书籍：适合大多数人算法零基础入门
        《算法新解》                                    - 入门书籍：思路清晰，语言得当
        《算法设计与分析基础》                            - 经典书籍：Java版算法思路总结，极力推荐
        《编程珠玑》                                    - 经典书籍：算法书的有口皆碑的佳作

算法进阶
        《算法》                                       - 经典书籍：Java版算法大部头
        《算法导论》                                    - 恢弘巨制：算法领域经典代表作
        《计算机程序设计艺术》                            - 恢弘巨制：计算机领域里程碑书籍 

算法习题
        《编程之美》
        《算法竞赛入门经典》刘汝佳·第2版
        《算法竞赛进阶指南》李煜东 
        《程序员代码面试指南》
        《剑指offer》                               
```

学习网址：
- [牛客网](https://www.nowcoder.com/)：面向基础与面试的算法题库
- [力扣](https://leetcode.com/)：著名的算法题网站
- [算法导论](http://open.163.com/special/opencourse/algorithms.html)：麻省理工学院公开课
- [JULY博主](https://blog.csdn.net/v_july_v)：算法讲解私人博客
- [oRblt专栏](https://blog.csdn.net/orbit)：算法讲解私人博客
 

## 附录：笔者自己的其他笔记

[推荐书籍](https://github.com/ruyuejun/polaris)：综合了所有技术对应的书籍与书评

[知识笔记汇总仓库](https://github.com/overnote)：所有技术笔记所在组织  

各类技术分类索引：
- [算法](https://github.com/overnote/Algorithm)：涉及数据结构、基础算法、算法进阶等内容，有Go、JS两个版本，未来增加C，Java
- [服务端](https://github.com/overnote/Server)：涉及服务端常用技术，如：nginx、mysql、redis等
- [JavaScript](https://github.com/overnote/JavaScript)：涉及网页开发、JS学习、Node、vue、react等
- [Go](https://github.com/overnote/Golang)：
- [Java]()：整理中
- [Python]()：整理中
