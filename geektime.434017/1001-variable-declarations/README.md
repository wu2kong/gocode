# 变量声明方法

### 最通用的声明方法

~~~
var age int = 18
~~~

### 多变量声明

~~~go
var (
    name string = "wu2kong"
    age int = 18
    gendar rune = 'f'
    availiable bool = true
)

var a, b, c int = 5, 6, 7
var (
    age, point int = 18, 32
    c, d, e rune = 'C', 'D', 'E'
)
~~~

### 语法糖1："var varName = initExpression"

省略类型信息声明的"语法糖"仅适用于在变量声明的同时显式赋予变量初值的情况

~~~go
// 编译器会根据右边的类型自动加上类型
var b = 13
// 或者，可以显示给定类型
var b = int32(13)
// 所以以下也是被允许的
var a, b, c = 12, 'A', "hello"
~~~

### 语法糖2：短变量声明

最简化的变量声明形式：短变量声明。使用短变量声明时，我们甚至可以省去 var 关键字以及类型信息，它的标准范式是"varName := initExpression"。

~~~go
a := 12
b := 'A'
c := "hello"


a, b, c := 12, 'A', "hello"
~~~

短变量声明中的变量类型也是由 Go 编译器自动推导出来的。

### Go 语言的两类变量

* 包级变量 (package varible)
* 局部变量 (local varible)

#### 包级变量

> 包级变量只能使用带有 var 关键字的变量声明形式，不能使用短变量声明形式，但在形式细节上可以有一定灵活度
一般使用语法糖1

~~~go
// 第一类：声明并同时显式初始化。

// $GOROOT/src/io/io.go
var ErrShortWrite = errors.New("short write")
var ErrShortBuffer = errors.New("short buffer")
var EOF = errors.New("EOF")

//第一种：
plain
var a = 13 // 使用默认类型
var b int32 = 17  // 显式指定类型
var f float32 = 3.14 // 显式指定类型

//第二种：更推荐这种
var a = 13 // 使用默认类型
var b = int32(17) // 显式指定类型
var f = float32(3.14) // 显式指定类型

//第二类：声明但延迟初始化。
var a int32
var f float64
~~~

> 声明聚类与就近原则，提高代码可读性

~~~go
// $GOROOT/src/net/net.go
var (
    netGo  bool 
    netCgo bool 
)

var (
    aLongTimeAgo = time.Unix(1, 0)
    noDeadline = time.Time{}
    noCancel   = (chan struct{})(nil)
)
~~~

### 局部变量

第一类：对于延迟初始化的局部变量声明，我们采用通用的变量声明形式

第二类：对于声明且显式初始化的局部变量，建议使用短变量声明形式

~~~go
// 接受默认类型的变量，我们使用下面这种形式：
a := 17
f := 3.14
s := "hello, gopher!"

// 不接受默认类型的变量，我们依然可以使用短变量声明形式，只是在":="右侧要做一个显式转型，以保持声明的一致性：
a := int32(17)
f := float32(3.14)
s := []byte("hello, gopher!")
~~~
