package main

import "fmt"

func main() {
	// --------------
	// Built-in types
	// 内置类型
	// --------------

	// Type provides integrity and readability.
	// - What is the amount of memory that we allocate?
	// - What does that memory represent?
	// 类型提供完整性和可读性。
	// - 分配的内存大小?
	// - 分配的内存指代了什么?

	// Type can be specific such as int32 or int64.
	// For example,
	// - uint8 contains a base 10  number using one byte of memory
	// - int32 contains a base 10 number using 4 bytes of memory.
	// 类型可以是特定的，例如 int32或 int64。
	// 例如,
	// - Uint8包含一个以10为基数的数字，使用一个内存字节
	// - Int32包含一个以10为基数的数字，占用4个字节的内存.

	// When we declare a type without being very specific, such as uint or int, it gets mapped
	// based on the architecture we are building the code against.
	// On a 64-bit OS, int will map to int64. Similarly, on a 32 bit OS, it becomes int32.
	// 当我们声明一个没有非常具体的类型（如uint或int）时，它会根据我们构建代码所依据的体系结构进行映射。
	// 在64位操作系统上，int 会映射到 int64，类似地，在32位操作系统上，它会变成 int32。
	

	// The word size is the number of bytes in a word, which matches our address size.
	// For example, in 64-bit architecture, the word size is 64 bit (8 bytes), address size is 64
	// bit then our integer should be 64 bit.
	// 字段size是指一个字中的字节数，它与我们的地址大小相匹配。
	//例如，在64位体系结构中，字大小为 64 位( 8 字节) ，地址大小为 64 位，那么我们的 integer 应该是 64 位。

	// ------------------
	// Zero value concept
	// 零值的概念
	// ------------------

	// Every single value we create must be initialized. If we don't specify it, it will be set to
	// the zero value. The entire allocation of memory, we reset that bit to 0.
	// - Boolean false
	// - Integer 0
	// - Floating Point 0
	// - Complex 0i
	// - String "" (empty string)
	// - Pointer nil
	// 我们创建的每个值都必须初始化。如果我们不指定它，它将被设置为零值。对于整个内存分配，我们将该位重置为 0。
	// 各种内置类型数据的零值：
	// - Boolean false
	// - Integer 0
	// - Floating Point 0
	// - Complex 0i
	// - String "" (空字符串)
	// - Pointer nil

	// Strings are a series of uint8 types.
	// A string is a two word data structure: first word represents a pointer to a backing array, the
	// second word represents its length.
	// If it is a zero value then the first word is nil, the second word is 0.
	// 字符串是一系列 uint8类型。
	// 字符串是有两个字段的数据结构: 第一个字段表示指向支撑数组的指针，第二个字段表示它的长度。
	// 如果字符串是一个零值，那么它第一个字段是 nil，第二个字段是0。

	// ----------------------
	// Declare and initialize
	// 声明与初始化
	// ----------------------

	// var is the only guarantee to initialize a zero value for a type.
	// Var 是初始化类型的零值的唯一保证
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Using the short variable declaration operator, we can define and initialize at the same time.
	// 使用短变量声明运算符，我们可以在定义的同时初始化
	aa := 10
	bb := "hello" // 1st word points to a array of characters, 2nd word is 5 bytes // 第一个字段指向一个字符数组，第二个字段是5个字节
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// ---------------------
	// Conversion vs casting
	// ---------------------

	// Go doesn't have casting, but conversion.
	// Instead of telling a compiler to pretend to have some more bytes, we have to allocate more
	// memory.
	// Go 没有 casting(强制转换), 但有 conversion(显式转换).
	// 我们必须分配更多的内存，而不是告诉编译器假装有更多的字节。

	// Specify type and perform a conversion.
	// 指定类型并执行转换
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
	
	// 区别：
	// casting: 假设有一个8bytes的变量转换成32bytes时会直接在当时内存地址上向前多取24bytes当做转换后的数据，
	// conversion: 先开辟32bytes的空间，然后将原来的8bytes数据复制到新开辟的空间上
}
