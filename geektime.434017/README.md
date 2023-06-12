# Tony Bai Go 语言第一课
课程很不错：来自极客时间专栏《》
https://time.geekbang.org/column/article/460666


## 入门知识点
* 安装 go 环境
* go 程序结构，go 语言执行次序
* go 项目布局
* go 包管理，go module
* 构建一个 web 应用程序


## 基础知识点
* 变量声明
* 代码块和作用域
* 数据类型
    * 基本数据类型
        * 整型
        * 浮点型
        * 字符串
        * bool 类型
    * 复合类型
        * 数组
        * 切片
        * map 类型
        * struct
* 控制结构
    * if
    * for、for range
    * switch、select
* 函数：输入、处理、输出；特定功能的代码块；可以拥有自定义类型
    * 参数传递：值拷贝
        * 传递数据的实际内容（开销成正比）：整型、数组、结构体
        * 传递数据内容的“描述符”（开销大小固定）：string、切片、map
    * 形参为接口类型
    * 形参为变长参数
    * 函数返回值
        * 普通函数返回值
        * 具名函数返回值
    * defer/setup/teardown
* 错误及异常
    * error
        * 普通错误
        * 自定义上下文错误
        * 自定义类型错误
    * panic/defer/recover






## 知识点
* 