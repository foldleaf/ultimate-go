// Reference types: slice, map, channel, interface, function.
// Zero value of a reference type is nil.
// 引用类型(Reference types): slice, map, channel, interface, function. //切片、映射、通道、接口、函数
// 引用类型的零值是 nil.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ----------------------
	// Declare and initialize
	// 声明与初始化
	// ----------------------

	// Create a slice with a length of 5 elements.
	// make is a special built-in function that only works with slice, map and channel.
	// make creates a slice that has an array of 5 strings behind it. We are getting back a 3 word
	// data structure: the first word points to the backing array, second word is length and third
	// one is capacity.
	// 创建一个切片，长度为5个元素。
	// make 是一个特殊的内置函数，只能用于slice、map和channel。
	// make 创建了一个包含 5个字符串的数组. 我们获得 3 字的数据结构
	// 第一个字是指针 ，指向支撑数组, 第二个字是长度 ， 第三个是容量(capacity)
	//  -----
	// |  *  | --> | nil | nil | nil | nil | nil |
	//  -----      |  0  |  0  |  0  |  0  |  0  |
	// |  5  |
	//  -----
	// |  5  |
	//  -----

	// ------------------
	// Length vs Capacity
	// 长度与容量
	// ------------------

	// Length is the number of elements from this pointer position we have access to (read and write).
	// Capacity is the total number of elements from this pointer position that exist in the
	// backing array.
	// 长度是我们从指针的位置开始能够访问(读和写)的元素数量
	// 容量是从指针的位置开始存在的支撑数组的元素总数

	// Syntactic sugar -> looks like array
	// It also have the same cost that we've seen in array.
	// One thing to be mindful about: there is no value in the bracket []string inside the make
	// function. With that in mind, we can constantly notice that we are dealing with a slice, not
	// array.
	// 语法糖(Syntactic sugar) -> 类似数组
	// 它与数组有着同样的开销
	// 需要注意的一点是：make函数中的[]string括号里没有值。考虑到这一点，我们经常注意到我们处理的是切片，而不是数组。
	slice1 := make([]string, 5)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	// We can't access an index of a slice beyond its length.
	// Error: panic: runtime error: index out of range
	// slice1[5] = "Runtime error"
	// 我们不能访问超出切片长度的索引
	
	// We are passing the value of slice, not its address. So the Println function will have its
	// own copy of the slice.
	// 我们传递的是 slice 的值，而不是它的地址. 所以 Println函数会有slice的拷贝
	fmt.Printf("\n=> Printing a slice\n")
	fmt.Println(slice1)

	// --------------
	// Reference type
	// 引用类型
	// --------------

	// Create a slice with a length of 5 elements and a capacity of 8.
	// make allows us to adjust the capacity directly on construction of this initialization.
	// What we end up having now is a 3 word data structure where the first word points to an array
	// of 8 elements, length is 5 and capacity is 8.
	// 创建一个 slice ,长度为 5 个元素，容量为 8 个元素
	// make 允许我们在构造初始化时直接调整容量。
	// 我们现在得到的是一个 3 字段的数据结构，其中第一个字段指向8个元素的数组，长度为5，容量为8。
	//  -----
	// |  *  | --> | nil | nil | nil | nil | nil | nil | nil | nil |
	//  -----      |  0  |  0  |  0  |  0  |  0  |  0  |  0  |  0  |
	// |  5  |
	//  -----
	// |  8  |
	//  -----
	// It means that I can read and write to the first 5 elements and I have 3 elements of capacity
	// that I can leverage later.
	// 这意味着我可以读写前5个元素，还有3个元素的容量稍后可以利用。
	slice2 := make([]string, 5, 8)
	slice2[0] = "Apple"
	slice2[1] = "Orange"
	slice2[2] = "Banana"
	slice2[3] = "Grape"
	slice2[4] = "Plum"

	fmt.Printf("\n=> Length vs Capacity\n")
	inspectSlice(slice2)

	// --------------------------------------------------------
	// Idea of appending: making slice a dynamic data structure
	// append的作用:使 slice 成为动态数据结构
	// --------------------------------------------------------
	fmt.Printf("\n=> Idea of appending\n")

	// Declare a nil slice of strings, set to its zero value.
	// 3 word data structure: first one points to nil, second and last are zero.
	// 声明一个 nil 的字符串数组, 用零值初始化.
	// 3 字段的数据结构: 第一个指向 nil, 第二与第三个为 0.
	var data []string

	// What if I do data := string{}? Is it the same?
	// No because data in this case is not set to its zero value.
	// This is why we always use var for zero value because not every type when we create an empty
	// literal we have its zero value in return.
	// What actually happen here is that we have a slice but it has a pointer (as opposed to nil).
	// This is consider an empty slice, not a nil slice.
	// There is a semantic between a nil slice and an empty slice. Any reference type that set to
	// its zero value can be considered nil. If we pass a nil slice to a marshal function, we get
	// back a string that said null but when we pass an empty slice, we get an empty JSON document.
	// But where does that pointer point to? It is an empty struct, which we will review later.
	// 如果这样写 data := string{}
	// 一样吗？不一样，因为这样数据并没有使用零值初始化。
	// 这就是为什么我们总是对零值使用var，因为当我们创建一个空值时，并不是每个类型都会返回其零值。
	// 这里实际发生的是，我们有一个切片，但它有一个指针（而不是nil）。这是一个空(empty)切片，而不是nil切片。
	// 在nil切片和空切片之间存在语义（引用语义和值语义）上的区别。任何设置为零值的引用类型都可以被视为 nil。
	// 如果我们向marshal函数传递一个nil切片，我们会返回一个表示 null的字符串，但当我们传递一个空切片时，我们会得到一个空的JSON。
	// 但指针指向哪里？它是一个空的结构结构体，我们稍后将对此进行检查。

	// Capture the capacity of the slice.
	// 捕获 slice 的容量
	lastCap := cap(data)

	// Append ~100k strings to the slice.
	// 向切片 Append（追加） 100k 个字符串 .
	for record := 1; record <= 102400; record++ {
		// Use the built-in function append to add to the slice.
		// It allows us to add value to a slice, making the data structure dynamic, yet still
		// allows us to use that contiguous block of memory that gives us the predictable access
		// pattern from mechanical sympathy.
		// The append call is working with value semantic. We are not sharing this slice but
		// appending to it and returning a new copy of it. The slice gets to stay on the stack, not
		// heap.
		// 使用内置函数 append 添加到 slice.
		// 它允许我们为一个切片添加值，使数据结构动态化，但仍然允许我们使用连续的内存块，这为我们提供了机器能理解的可预测访问模式。
		// Append 的调用使用的是值语义 ，我们不是共享这个切片，而是附加给它并返回它的新副本。
		data = append(data, fmt.Sprintf("Rec: %d", record))

		// Every time append runs, it checks the length and capacity.
		// If it is the same, it means that we have no room. append creates a new backing array,
		// double its size, copy the old value back in and append the new value. It mutates its copy
		// on its stack frame and return us a copy. We replace our slice with the new copy.
		// If it is not the same, it means that we have extra elements of capacity we can use. Now we
		// can bring these extra capacity into the length and no copy is being made. This is very
		// efficient.
		// 没次 append 执行时, 都会检查长度和容量
		// 如果长度和容量一样, 意味着没有多余的空间了. append 会创建一个新的支撑数组,大小为原来的两倍，
		// 然后将原来的值复制过来并附加新值 . 它会在栈帧上对其副本进行转换并返回新的副本，我们使用新的副本来声明 slice
		// 如果长度和容量不一样，这意味着我们可以使用额外的容量元素
		// 现在我们可以将这些额外的容量添加到长度中，并且不需要复制。这很高效

		// Looking at the last column in the output, when the backing array is 1000 elements or
		// less, it doubles the size of the backing array for growth. Once we pass 1000 elements,
		// growth rate moves to 25%.
		// 查看输出中的最后一列，当支撑数组为1000个元素或更少时，它会将支撑数组的大小增加一倍。
		// 一旦我们超过了1000项元素，增长率将达到25%。

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {
			// Calculate the percent of change.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for capacity.
			lastCap = cap(data)

			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)
		}
	}

	// --------------
	// Slice of slice
	// 切片
	// --------------

	// Take a slice of slice2. We want just indexes 2 and 3.
	// The length is slice3 is 2 and capacity is 6.
	// Parameters are [starting_index : (starting_index + length)]
	// By looking at the output, we can see that they are sharing the same backing array.
	// These slice headers get to stay on the stack when we use these value semantics. Only the
	// backing array that needed to be on the heap.
	// 从slice2中截取切片. 我们只想要 索引 2 和 3.
	// slice3 的长度为 2 ，容量为 6.
	// 参数为 [起始索引: (起始索引 + 截取的切片长度)]
	// 观察输出, 我们可以看到它们共享同一个支撑数组
	// 当我们使用这些值语义时，这些切片头(slice headers)会处于在栈上。 只有支撑数组处于堆上
	slice3 := slice2[2:4]

	fmt.Printf("\n=> Slice of slice (before)\n")
	inspectSlice(slice2)
	inspectSlice(slice3)

	// When we change the value of the index 0 of slice3, who are going to see this change?
	// 当我们更改slice3的索引0的值时，谁会看到这种变化？
	slice3[0] = "CHANGED"

	// The answer is both.
	// We have to always to aware that we are modifying an existing slice. We have to be aware who
	// are using it, who is sharing that backing array.
	// 答案是所有(切片)都可以.
	// 我们必须始终意识到我们正在修改一个现有的切片. 我们必须知道到谁在使用它，谁在共享支撑数组
	fmt.Printf("\n=> Slice of slice (after)\n")
	inspectSlice(slice2)
	inspectSlice(slice3)

	// How about slice3 := append(slice3, "CHANGED")?
	// Similar problem will occur with append if the length and capacity is not the same.
	// Instead of changing slice3 at index 0, we call append on slice3. Since the length of slice3
	// is 2, capacity is 6 at the moment, we have extra rooms for modification. We go and change
	// the element at index 3 of slice3, which is index 4 of slice2. That is very dangerous.
	// 如果长度和容量不相同，append也会出现类似的问题。
	// 我们不在索引0处更改 slice3，而是在slice3调用 append 。
	// 由于现在slice3的长度为2，容量为6，我们有额外的空间去修改。
	// 我们更改slice3的索引3处的元素，也是slice2的索引4处的元素，这是非常危险的
	
	// So, what if the length and capacity is the same? Instead of making slice3 capacity 6, we set
	// it to 2 by adding  another parameter to the slicing syntax like this: slice3 := slice2[2:4:4]
	// When append looks at this slice and see that the length and capacity is the same, it wouldn't
	// bring in the element at index 4 of slice2. It would detach.
	// slice3 will have a length of 2 and capacity of 2, still share the same backing array.
	// On the call to append, length and capacity will be different. The addresses are also different.
	// This is called 3 index slice. This new slice will get its own backing array and we don't
	// affect anything at all to our original slice.
	// 那么如果长度和容量一样呢？
	// 我们不将slice的容量设置为6，而是通过向切片语法添加另一个参数，将容量设置为2: slice3 := slice2[2:4:4]
	// 当 append 查看这个切片，发现长度和容量是相同的时候，它不会引入 slice2索引4处的元素。它们会分离开
	// Slice3的长度为2，容量为2，仍然共享相同的支撑数组。
	// append 的调用中，长度与容量不同，地址也不同
	// 这是3个索引切片。这个新切片将有自己的支撑数组，我们对原始切片没有任何影响。


	// ------------
	// Copy a slice
	// Copy 切片
	// ------------

	// copy only works with string and slice only.
	// Make a new slice big enough to hold elements of original slice and copy the values over using
	// the built-in copy function.
	// copy函数只对字符串和切片有效
	// 创建一个足够大的新切片来保存原始切片的元素，并使用内置的copy函数拷贝它们的值。
	slice4 := make([]string, len(slice2))
	copy(slice4, slice2)

	fmt.Printf("\n=> Copy a slice\n")
	inspectSlice(slice4)

	// -------------------
	// Slice and reference
	// 切片与引用
	// -------------------

	// Declare a slice of integers with 7 values.
	x := make([]int, 7)

	// Random starting counters.
	for i := 0; i < 7; i++ {
		x[i] = i * 100
	}

	// Set a pointer to the second element of the slice.
	twohundred := &x[1]

	// Append a new value to the slice. This line of code raises a red flag.
	// We have x is a slice with length 7, capacity 7. Since the length and capacity is the same,
	// append doubles its size then copy values over. x now points to different memory block and
	// has a length of 8, capacity of 14.
	// 在切片上append一个新值。这行代码会出现一个红色标志
	// x是一个长度为7，容量也为7的切片。由于长度与容量相等，append会扩大容量为原来的两倍并拷贝原值，
	// x现在指向不同的内存块，且长度为8，容量为14
	x = append(x, 800)

	// When we change the value of the second element of the slice, twohundred is not gonna change
	// because it points to the old slice. Everytime we read it, we will get the wrong value.
	// 当我们变更切片第二个元素的值，twohundred 不会改变
	// 因为它指向的是旧切片，每次我们读取它都会得到错误的值
	x[1]++

	// By printing out the output, we can see that we are in trouble.
	fmt.Printf("\n=> Slice and reference\n")
	fmt.Println("twohundred:", *twohundred, "x[1]:", x[1])

	// -----
	// UTF-8
	// -----
	fmt.Printf("\n=> UTF-8\n")

	// Everything in Go is based on UTF-8 character sets.
	// If we use different encoding scheme, we might have a problem.

	// Declare a string with both Chinese and English characters.
	// For each Chinese character, we need 3 byte for each one.
	// The UTF-8 is built on 3 layers: bytes, code point and character. From Go perspective, string
	// are just bytes. That is what we are storing.
	// In our example, the first 3 bytes represents a single code point that represents that single
	// character. We can have anywhere from 1 to 4 bytes representing a code point (a code point is
	// a 32 bit value) and anywhere from 1 to multiple code points can actually represent a
	// character. To keep it simple, we only have 3 bytes representing 1 code point representing 1
	// character. So we can read s as 3 bytes, 3 bytes, 1 byte, 1 byte,... (since there are only 2
	// Chinese characters in the first place, the rest are English)
	s := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune -> maximum number of bytes we need to
	// represent any code point is 4.
	// Rune is its own type. It is an alias for int32 type. Similar to type byte we are using, it
	// is just an alias for uint8.
	var buf [utf8.UTFMax]byte

	// When we are ranging over a string, are we doing it byte by byte or code point by code point or
	// character by character?
	// The answer is code point by code point.
	// On the first iteration, i is 0. On the next one, i is 3 because we are moving to the next
	// code point. Then i is 6.
	for i, r := range s {
		// Capture the number of bytes for this rune/code point.
		rl := utf8.RuneLen(r)

		// Calculate the slice offset for the bytes associated with this rune.
		si := i + rl

		// Copy rune from the string to our buffer.
		// We want to go through every code point and copy them into our array buf, and display
		// them on the screen.
		// "Every array is just a slice waiting to happen." - Go saying
		// We are using the slicing syntax, creating our slice header where buf becomes the backing
		// array. All of them are on the stack. There is no allocation here.
		copy(buf[:], s[i:si])

		// Display the details.
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}

// inspectSlice exposes the slice header for review.
// Parameter: again, there is no value in side the []string so we want a slice.
// Range over a slice, just like we did with array.
// While len tells us the length, cap tells us the capacity
// In the output, we can see the addresses are aligning as expected.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}
