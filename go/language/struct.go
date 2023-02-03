package main

import "fmt"

// example represents a type with different fields.
// example表示具有不同字段的类型。
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// ----------------------
	// Declare and initialize
	// ----------------------

	// Declare a variable of type example set to its zero value.
	// How much memory do we allocate for example?
	// a bool is 1 byte, int16 is 2 bytes, float32 is 4 bytes
	// Putting together, we have 7 bytes. However, the actual answer is 8.
	// That leads us to a new concept of padding and alignment.
	// The padding byte is sitting between the bool and the int16. The reason is because of
	// alignment.
	// 声明一个类型为 example 的变量设置为其零值。
	// 我们应该分配多少内存给 example ?
	// bool 是1字节，int16是2字节，float32是4字节
	// 放在一起就是7个字节，但是实际的答案是8。
	// 这里我们将引入新的概念：填充(padding)和对齐 (alignment)。
	// 填充的字节位于bool和int16之间。原因是对齐。

	// The idea of alignment: It is more efficient for this piece of hardware to read memory on its
	// alignment boundary. We will take care of the alignment boundary issues so the hardware
	// people don't.
	// 总之对齐就是为了提高硬件读取内存的效率

	// Rule 1:
	// Depending on the size a particular value, Go determines the alignment we need. Every 2 bytes
	// value must fall on a 2 bytes boundary. Since the bool value is only 1 byte and start at
	// address 0, then the next int16 must start on address 2. The byte at address that get skipped
	// over becomes a 1 byte padding. Similarly, if it is a 4 bytes value then we will have a 3
	// bytes padding value.
	// 规则 1:
	// 根据特定值的大小，Go 确定我们需要的对齐方式。
	// 根据特定值的大小，Go决定我们需要的对齐方式。每2个字节的值必须落在2个字节边界上。
	// 由于布尔值仅为1字节，并且从地址0开始，因此下一个int16必须从地址2开始。
	// 被跳过的地址处的字节变为1字节填充。类似地，如果它是一个4字节的值，那么我们将有一个3字节的填充值。
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Rule 2:
	// The largest field represents the padding for the entire struct.
	// We need to minimize the amount of padding as possible. Always lay out the field
	// from highest to smallest. This will push any padding down to the bottom.
	// 规则:
	// 最大的字段表示整个结构的padding.
	// 我们需要尽可能减少padding. 所以需要将字段按内存从大到小排列放置

	// In this case, the entire struct size has to follow a 8 bytes value because int64 is 8 bytes.
	// 在这种情况下，整个结构大小必须以8字节为准，因为 int64是8字节。
	// type example struct {
	//     counter int64
	//     pi      float32
	//     flag    bool
	// }

	// 声明 example 类型的变量并使用结构体语法初始化.
	// 每行必须以逗号","结尾
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	// Declare a variable of an anonymous type and init using a struct literal.
	// This is one time thing.
	// 声明一个匿名(anonymous)类型的变量并使用结构体语法初始化
	// 这只需要一个步骤
	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi", e3.pi)

	// ---------------------------
	// Name type vs anonymous type
	// 非匿名类型与匿名类型
	// ---------------------------

	// If we have two name type identical struct, we can't assign one to another.
	// For example, example1 and example2 are identical struct, var ex1 example1, var ex2 example2.
	// ex1 = ex2 is not allowed. We have to explicitly say that ex1 = example1(ex2) by performing a
	// conversion.
	// However, if ex is a value of identical anonymous struct type (like e3 above), then it is possible to
	// assign ex1 = ex
	// 如果有两个名称类型相同的结构体，则不能将其中一个分配给另一个。
	// 例如，example1和 example2是相同的结构体，var ex1 example1，var ex2 example2。
	// ex1 = ex2 这是不允许的. 我们必须使用类型转换清晰地说明 ex1 = example1(ex2) 
	// 但是，如果 ex 是相同的匿名结构体类型的值(如上面的e3)，那就可以赋值ex1 = ex
	var e4 example
	e4 = e3

	fmt.Printf("%+v\n", e4)
}
