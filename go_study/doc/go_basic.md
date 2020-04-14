### 变量

声明变量的一般形式是使用 var 关键字：

`var name type`

声明多个变量

`var a,b,c,d int`

还可以赋值

`var a,b,c,d = 1,'2',"3",false`

或者使用关键字 var 和括号，可以将一组变量定义放在一起

```
var (
   a int
   b string
   c bool
)
```

简化   名字 := 表达式  自动推导类型，但是后续不能修改变量的类型

`x:=10`



### 常量

常量是在编译时被创建的，通常是const name [type] = value

`const a = 1.5`

通常type可以省略，自动推导

const也可以和变量一样使用括号

```
const (
    e  = 2.7182818
    pi = 3.1415926
)
```

### 浮点数

这些浮点数类型的取值范围可以从很微小到很巨大。浮点数取值范围的极限值可以在 math 包中找到：

- 常量 math.MaxFloat32 表示 float32 能取到的最大数值，大约是 3.4e38；
- 常量 math.MaxFloat64 表示 float64 能取到的最大数值，大约是 1.8e308；
- float32 和 float64 能表示的最小值分别为 1.4e-45 和 4.9e-324。

`var f float32 = 1677216 // 1 << 24`

很小或很大的数最好用科学计数法书写，通过 e 或 E 来指定指数部分：

```
const Avogadro = 6.02214129e23  // 阿伏伽德罗常数const Planck   = 6.62606957e-34 // 普朗克常数
```

用 Printf 函数打印浮点数时可以使用“%f”来控制保留几位小数

### for

和C语言的for类似使用，下面举几个例子

```
package main

import "fmt"

func main() {
	s := "hello, world! 你好"

	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%c,", s[i])
	}

	fmt.Println()

	for _, r := range s { // rune
		fmt.Printf("%c,", r)
	}

	fmt.Println()

	n := len(s) - 1
	for n > 0 { // 类似while
		fmt.Printf("%c", s[n])
		n--
	}
}
```

### 数组

数组的基本形式：var 数组变量名 [元素数量]Type

`var a [3]int // 定义三个整数的数组`

初始化

```
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
```

如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算，因此，上面数组 q 的定义可以简化为：

`q:=[...]int{1,2,3}`

#### 比较两个数组是否相等

如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（`==`和`!=`）来判断两个数组是否相等，只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组，否则程序将无法完成编译。

```
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
fmt.Println(a == b, a == c, b == c) // "true false false"
d := [3]int{1, 2}
fmt.Println(a == d) // 编译错误：无法比较 [2]int == [3]int
```

#### 遍历数组——访问每一个数组元素

遍历数组也和遍历切片类似，代码如下所示：

```
var team [3]string
team[0] = "hammer"
team[1] = "soldier"
team[2] = "mum"

for k, v := range team {    
	fmt.Println(k, v)
}
```

代码输出结果：

```
0 hammer
1 soldier
2 mum
```

代码说明如下：

- 第 6 行，使用 for 循环，遍历 team 数组，遍历出的键 k 为数组的索引，值 v 为数组的每个元素值。
- 第 7 行，将每个键值打印出来。

### Range

range类似迭代器操作，返回 (索引, 值) 或 (键, 值)

可忽略不想要的返回值，或⽤用 "_" 这个特殊变量

```
package main

import "fmt"

func main() {

	s := "abc"
	for i := range s { // 忽略 2nd value，支持 string/array/slice/map
		fmt.Printf("%c", s[i])
	}
	fmt.Println()
	for _, c := range s { // 忽略 index
		fmt.Printf("%c", c)
	}
	fmt.Println()
	for range s { // 忽略全部返回值，仅迭代

	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m { // 返回 (key, value)
		fmt.Println(k, v)
	}
}
```

Range会复制对象，range的key, value都是从对象里复制出来的，因此不会干扰原来对象的数据，会额外占据内存

```go
package main

import "fmt"

func main() {

	a := [3]int{0, 1, 2}
	for i, v := range a { 	// index、value 都是从复制品中取出
		if i == 0 { 		// 在修改前，我们先修改原数组
			a[1], a[2] = 999, 999
			fmt.Println(a) 	// 确认修改有效，输出 [0, 999, 999]
		}
		a[i] = v + 100 		// 使⽤用复制品中取出的 value 修改原数组
	}
	fmt.Println(a) 			// 输出 [100, 101, 102]
}
```

改⽤用引⽤用类型，其底层数据不会被复制

```go
package main

import "fmt"

func main() {
	/**
	range在遍历的时候会复制对象的index和value，不会干扰原对象的数据
	因此会额外占据空间
	 */

	a := [3]int{0, 1, 2}
	for i, v := range a { 	// index、value 都是从复制品中取出
		if i == 0 { 		// 在修改前，我们先修改原数组
			a[1], a[2] = 999, 999
			fmt.Println(a) 	// 确认修改有效，输出 [0, 999, 999]
		}
		a[i] = v + 100 		// 使⽤用复制品中取出的 value 修改原数组
	}
	fmt.Println(a) 			// 输出 [100, 101, 102]

	fmt.Println("--------------------------------")

	/**
	改⽤用引⽤用类型，其底层数据不会被复制
	 */
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {	// 复制 struct slice { pointer, len, cap }
		if i == 0 {
			s = s[:3]		// 对 slice 的修改，不会影响 range
			s[2] = 100		// 对底层数据的修改
		}
		fmt.Println(i, v)
	}
}
```

### switch

分⽀支表达式可以是任意类型，不限于常量。可省略 break，默认⾃自动终⽌止。

如需要继续下⼀一分⽀支，可使⽤用 fallthrough，但不再判断条件

```go
package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	i := 2
	switch i {
	case x[1]:
		fmt.Println("a")
	case 1, 3:
		fmt.Println("b")
	default:
	}
	fmt.Println("c")

	// 如需要继续下⼀一分⽀支，可使⽤用 fallthrough，但不再判断条件
	x2 := 10
	switch x2 {
	case 10:
		fmt.Println("a")
		fallthrough
	case 0:
		fmt.Println("b")
	}
}
```

省略条件表达式，可当 if...else if...else 使⽤用

```go
// 省略条件表达式，可当 if...else if...else 使⽤用
	switch {
	case x[1] > 0:
		fmt.Println("a")
	case x[1] < 0:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
	switch i := x[2]; { // 带初始化语句
	case i > 0:
		fmt.Println("a")
	case i < 0:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
```

### append

Go语言的内建函数 append() 可以为切片动态添加元素

```
var a []int
a = append(a, 1) // 追加1个元素
a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
```

不过需要注意的是，在使用 append() 函数为切片动态添加元素时，如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变。

切片slice通常拥有type, len, cap [类型，长度，容量]，当len增长cap会动态变化适应len增长，在一定范围内是2倍增长，后续是1.25倍。

因为 append 函数返回新切片的特性，所以切片也支持链式操作，我们可以将多个 append 操作组合起来，实现在切片中间插入元素：

```
var a []inta = append(a[:i], append([]int{x}, a[i:]...)...) // 在第i个位置插入xa = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片
```

### map

map 是引用类型，可以使用如下方式声明：

`var mapname map[keytype]valuetype`

其中：

- mapname 为 map 的变量名。
- keytype 为键类型。
- valuetype 是键对应的值类型。

> 提示：[keytype] 和 valuetype 之间允许有空格。

在声明的时候不需要知道 map 的长度，因为 map 是可以动态增长的，未初始化的 map 的值是 nil，使用函数 len() 可以获取 map 中 pair 的数目。

```go
package main

import "fmt"

func main() {
	var mapLit map[string]int
	// var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}
```

示例中 mapLit 演示了使用`{key1: value1, key2: value2}`的格式来初始化 map ，就像数组和结构体一样。

上面代码中的 mapCreated 的创建方式`mapCreated := make(map[string]float)`等价于`mapCreated := map[string]float{} `。

mapAssigned 是 mapList 的引用，对 mapAssigned 的修改也会影响到 mapLit 的值。

注意：可以使用 make()，但不能使用 new() 来构造 map，如果错误的使用 new() 分配了一个引用对象，会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：

`mapCreated := new(map[string]float)`

接下来当我们调用`mapCreated["key1"] = 4.5`的时候，编译器会报错：

`invalid operation: mapCreated["key1"] (index of type *map[string]float).`

#### map遍历

map 的遍历过程使用 for range 循环完成

```
/**
	  map遍历
	*/
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
	}
	
	// 只遍历值
	for _, v := range scene {
		fmt.Println(v)
	}

	// 只遍历 key
	for k := range scene {
		fmt.Println(k)
	}
```

如果需要特定顺序的遍历结果，正确的做法是先排序，代码如下：

```
scene := make(map[string]int)
// 准备map数据
scene["route"] = 66
scene["brazil"] = 4
scene["china"] = 960
// 声明一个切片保存map数据
var sceneList []string
// 将map数据遍历复制到切片中
for k := range scene {
	sceneList = append(sceneList, k)
}
// 对切片进行排序
sort.Strings(sceneList)
// 输出
fmt.Println(sceneList)

// [brazil china route]
```

sort.Strings 的作用是对传入的字符串切片进行字符串字符的升序排列

### 函数

函数是Go里面的核心设计，它通过关键字`func`来声明，它的格式如下：

```Go
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2
}
```

上面的代码我们看出

- 关键字`func`用来声明一个函数`funcName`
- 函数可以有一个或者多个参数，每个参数后面带有类型，通过`,`分隔
- 函数可以返回多个值
- 上面返回值声明了两个变量`output1`和`output2`，如果你不想声明也可以，直接就两个类型
- 如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
- 如果没有返回值，那么就直接省略最后的返回信息
- 如果有返回值， 那么必须在函数的外层添加return语句

#### 多个返回值

Go语言比C更先进的特性，其中一点就是函数能够返回多个值。

```Go
package main

import "fmt"

//返回 A+B 和 A*B
func SumAndProduct(A, B int) (int, int) {
	return A+B, A*B
}

func main() {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
```

上面的例子我们可以看到直接返回了两个参数，当然我们也可以命名返回参数的变量，这个例子里面只是用了两个类型，我们也可以改成如下这样的定义，然后返回的时候不用带上变量名，因为直接在函数里面初始化了。

但如果你的函数是导出的(首字母大写)

官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差。

```go
func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A+B
	Multiplied = A*B
	return
}
```

**这里命名了返回值，因此返回的时候默认返回已经命名的值，只需要最后加上return语句即可，默认提取返回值进行返回**

#### 变参

Go函数支持变参，接受变参的函数是有着不定数量的参数的

```Go
func myfunc(arg ...int) {}
```

`arg ...int`告诉Go这个函数接受不定数量的参数。注意，这些参数的类型全部是`int`。在函数体中，变量`arg`是一个`int`的`slice`：

```Go
for _, n := range arg {
	fmt.Printf("And the number is: %d\n", n)
}
```

#### defer

Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。如下代码所示，我们一般写打开一个资源是这样操作的：

```Go
func ReadWrite() bool {
	file.Open("file")
// 做一些工作
	if failureX {
		file.Close()
		return false
	}

	if failureY {
		file.Close()
		return false
	}

	file.Close()
	return true
}
```

我们看到上面有很多重复的代码，Go的`defer`有效解决了这个问题。使用它后，不但代码量减少了很多，而且程序变得更优雅。在`defer`后指定的函数会在函数退出前调用。

```Go
func ReadWrite() bool {
	file.Open("file")
	defer file.Close()
	if failureX {
		return false
	}
	if failureY {
		return false
	}
	return true
}
```

如果有很多调用`defer`，那么`defer`是采用后进先出模式

```
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		} else {
			defer fmt.Printf("%d ", i)
		}
	}
}
```

输出结果： 0 2 4 3 1 

#### 函数作为值、类型

在Go中函数也是一种变量，我们可以通过`type`来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型

	type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])

函数作为类型到底有什么好处呢？那就是可以把这个类型的函数当做值来传递，请看下面的例子

```Go
package main

import "fmt"

type testInt func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main(){
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)    // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)  // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}
```

函数当做值和类型在我们写一些通用接口的时候非常有用，通过上面例子我们看到`testInt`这个类型是一个函数类型，然后两个`filter`函数的参数和返回值与`testInt`类型是一样的，但是我们可以实现很多种的逻辑，这样使得我们的程序变得非常的灵活。



### method

带有接收者的函数，我们称为`method`

`func (r ReceiverType) funcName(parameters) (results)`

下面的例子可以很好说明method的接收者不同而调用的函数也不同

```go
package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
```



### 并发

#### goroutine

goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过`go`关键字实现了，其实就是一个普通的函数。

```Go
go hello(a, b, c)
```

通过关键字go就启动了一个goroutine。

```go
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行
}
```

输出结果：

world
hello
hello
world
world
hello
world
hello
hello

#### channel

goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通信呢，Go提供了一个很好的通信机制channel。channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：

```Go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```

channel通过操作符`<-`来接收和发送数据

```Go
ch <- v    // 发送v到channel ch.
v := <-ch  // 从ch中接收数据，并赋值给v
```

我们把这些应用到我们的例子中来：

```Go
package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total  // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c  // receive from c

	fmt.Println(x, y, x + y)
}
```

默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。