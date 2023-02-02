// ---------
// CPU CACHE
// CPU 缓存
// ---------

// Cores DO NOT access main memory directly but their local caches.
// What store in caches are data and instruction.
// 内核不能直接访问主存，只能访问它们的本地缓存
// 存储在缓存中的是数据和指令

// Cache speed from fastest to slowest: L1 -> L2 -> L3 -> main memory.
// Scott Meyers: "If performance matter then the total memory you have is the total amount of
// caches" -> access to main memory is incredibly slow; practically speaking it might not even be there.
// 缓存速度从快到慢: L1-> L2-> L3-> 主存。
// Scott Meyers: "如果性能很重要，那么您拥有的总内存就是缓存” 
// -> 访问主存的速度非常缓慢; 甚至可能不存在（主存）.

// How do we write code that can be sympathetic with the caching system to make sure that
// we don't have a cache miss or at least, we minimize cache misses to our fullest potential?
// 我们如何编写能够与缓存系统产生共鸣的代码，以确保我们不会发生缓存丢失，或者至少尽可能减少缓存丢失？


// Processor has a Prefetcher. It predicts what data is needed ahead of time.
// There are different granularity depending on where we are on the machine.
// Our programming model uses a byte. We can read and write to a byte at a time. However, from the
// caching system POV, our granularity is not 1 byte. It is 64 bytes, called a cache line. All
// memory is junked up in this 64 bytes cache line.
// 处理器(Processor) 有一个预取器 (Prefetcher). 它会提前预测需要哪些数据.
// 根据我们在机器的位置不同有粒度也不同.
// 我们的程序模型使用一个字节。我们可以一次读写一个字节. 
// 然而，在缓存系统POV看来, 我们的粒度不只1字节，而是64 字节, 被称为缓存行. 
// 所有内存都在64字节的缓存行内被抛弃(没理解什么意思)

// Since the caching mechanism is complex, Prefetcher tries to hide all the latency from us.
// It has to be able to pick up on predictable access pattern to data.
// -> We need to write code that creates predictable access pattern to data
// 由于缓存机制很复杂，所以预取器Prefetcher试图向我们隐藏所有延迟。
// 它只能够识别可预测的数据访问模式。
// -> 我们需要编写能够创建可预测的数据访问模式的代码

// One easy way is to create a contiguous allocation of memory and to iterate over them.
// The array data structure gives us ability to do so.
// From the hardware perspective, array is the most important data structure.
// From Go perspective, slice is. Array is the backing data structure for slice (like Vector in C++).
// Once we allocate an array, whatever it size, every element is equal distant from other element.
// As we iterate over that array, we begin to walk cache line by cache line. As the Prefetcher see
// that access pattern, it can pick it up and hide all the latency from us.
// 一种简单的方法是创建一个连续的内存分配，并对它们进行迭代。
// 数组(array)数据结构使我们能够这样做
// 从硬件的角度来看, 数组是最重要的数据结构。
// 从 Go 语言的角度来看, 切片(slice) 最重要. 数组是切片的支撑数据结构 (像 Vector 在 C++ 中).
// 一旦我们分配了一个数组，不管其大小，每个元素与其他元素的距离都是相等的。
// 当我们迭代那个数组时，我们开始遍历缓存行. 
// 当预取器看到这种访问模式时，它可以获取取并向我们隐藏所有的延迟。

// For example, we have a big nxn matrix. We do LinkedList Traverse, Column Traverse, and Row Traverse
// and benchmark against them.
// Unsurprisingly, Row Traverse has the best performance. It walk through the matrix cache line
// by cache line and create a predictable access pattern.
// Column Traverse does not walk through the matrix cache line by cache line. It looks like random
// access memory pattern. That is why is the slowest among those.
// However, that doesn't explain why the LinkedList Traverse's performance is in the middle. We
// just think that it might perform as poorly as the Column Traverse.
// -> This leads us to another cache: TLB - Translation lookaside buffer. Its job is to maintain
// operating system page and offset to where physical memory is.
// 例如，我们有一个大的nxn矩阵。我们进行LinkedList（链表）遍历、Column（列）遍历和Row（行）遍历，并对它们进行基准测试。
// 不出所料，行遍历具有最好的性能。它遍历矩阵缓存行并创建一个可预测的访问模式
// 列遍历不会遍历矩阵缓存行。它看起来像随机存取存储器模式。这就是它在其中最慢的原因。
// 然而，这并不能解释为什么LinkedList遍历的性能处于中间。我们本来认为它的表现可能和列遍历一样糟糕。
// -> 这指出了另一个缓存: TLB - 转译后备缓冲器(Translation lookaside buffer). 
// 它的任务是维护操作系统页面和物理内存所在位置的偏移量。

// ----------------------------
// Translation lookaside buffer (TLB)
// ----------------------------

// Back to the different granularity, the caching system moves data in and out the hardware at 64
// bytes at a time. However, the operating system manages memory by paging its 4K (traditional page
// size for an operating system).
// TLB: For every page that we are managing, let's take our virtual memory addresses that we use
// (softwares run virtual addresses, its sandbox, that is how we use/share physical memory)
// and map it to the right page and offset for that physical memory.
// 回到之前说到的不同粒度，缓存系统每次以64字节的速度将数据移入和移出硬件. 
// 而操作系统通过分页4K（操作系统的传统页面大小）来管理内存
// TLB：对于我们正在管理的每个页面，让我们获取我们使用的虚拟内存地址
//（软件运行虚拟地址，它的沙盒(sandbox)，这就是我们使用/共享物理内存的方式），并将其映射到正确的页面和物理内存的偏移量。

// A miss on the TLB can be worse than just the cache miss alone.
// The LinkedList is somewhere in between is because the chance of multiple nodes being on the same
// page is probably pretty good. Even though we can get cache misses because cache lines aren't
// necessary in the distance that is predictable, we probably don't have so many TLB cache misses.
// In the Column Traverse, not only we have cache misses, we probably have a TLB cache miss on
// every access as well.
// TLB 上的缺失(miss)可能比仅缓存缺失更糟糕。
// LinkedList介于两者之间，因为多个节点位于同一页面上的可能性很高。 
// 尽管 由于在可预测的距离内不需要缓存行导致会发生缓存缺失，但我们可能没有那么多TLB缓存缺失。
// 在列遍历中，我们不仅有缓存缺失，而且可能每次访问都有TLB缓存1缺失。

// Data-oriented design matters.
// It is not enough to write the most efficient algorithm, how we access our data can have much
// more lasting effect on the performance than the algorithm itself.
// 面向数据的设计问题
// 仅仅编写最有效的算法是不够的，我们如何访问数据对性能的影响比算法本身的影响更持久。


package main

import "fmt"

func main() {
	// -----------------------
	// Declare and initialize
	// 声明与初始化
	// -----------------------

	// Declare an array of five strings that is initialized to its zero value.
	// Recap: a string is a 2 word data structure: a pointer and a length
	// Since this array is set to its zero value, every string in this array is also set to its
	// zero value, which means that each string has the first word pointed to nil and
	// second word is 0.
	// 声明一个由五个字符串组成的数组，用零值(zero value)初始化该数组
	// 简介: 字符串是一个 2字段 (word) 的数据结构: 一个指针 pointer 和一个长度 length
	// 由于此数组被设置为零值，因此此数组中的每个字符串也被设置为其零值，这意味着每个字符串的第一个字段指向 nil， 第二个字段是 0。
	//  -----------------------------
	// | nil | nil | nil | nil | nil |
	//  -----------------------------
	// |  0  |  0  |  0  |  0  |  0  |
	//  -----------------------------
	var strings [5]string

	// At index 0, a string now has a pointer to a backing array of bytes (characters in string)
	// and its length is 5.
	// 在索引 0 处，一个字符串现在有一个指针指向支撑字节数组(字符串中的字符)，这个数组的长度为 5

	// -----------------
	// What is the cost?
	// -----------------

	// The cost of this assignment is the cost of copying 2 bytes.
	// We have two string values that have pointers to the same backing array of bytes.
	// Therefore, the cost of this assignment is just 2 words.
	// 以下情形内存会分配给其复制 2 个字节的开销。
	// 我们有两个字符串值，它们指向相同的支撑字节数组。
	// 因此，此任务的成本仅为2个字

	//  -----         -------------------
	// |  *  |  ---> | A | p | p | l | e | (1)
	//  -----         -------------------
	// |  5  |                  A
	//  -----                   |
	//                          |
	//                          |
	//     ---------------------
	//    |
	//  -----------------------------
	// |  *  | nil | nil | nil | nil |
	//  -----------------------------
	// |  5  |  0  |  0  |  0  |  0  |
	//  -----------------------------
	strings[0] = "Apple"
	strings[1] = "Orange"
	strings[2] = "Banana"
	strings[3] = "Grape"
	strings[4] = "Plum"

	// ---------------------------------
	// Iterate over the array of strings
	// 遍历字符串数组
	// ---------------------------------

	// Using range, not only we can get the index but also a copy of the value in the array.
	// fruit is now a string value; its scope is within the for statement.
	// In the first iteration, we have the word "Apple". It is a string that has the first word
	// also points to (1) and the second word is 5.
	// So we now have 3 different string value all sharing the same backing array.
	// 使用 range ，我们不仅可以得到索引，还可以得到数组中值的拷贝。
	// fruit 现在是一个字符串值; 它的作用域在 for 语句中。
	// 在第一次循环中，我们有 "Apple" 这个单词，这是一个字符串，第一个字指向(1)(见上图)，第二个字为 5
	// 现在我们有 3 个不同的字符串值共享同一个支撑数组。	——fruit、Apple、5

	// What are we passing to the Println function?
	// We are using value semantic here. We are not sharing our string value. Println is getting
	// its own copy, its own string value. It means when we get to the Println call, there are now
	// 4 string values all sharing the same backing array.
	// 我们要传递什么给 Println 函数？
	// 我们在这里使用了值语义。我们没有共享我们的字符串值。Printf 获取它自己的拷贝和它的字符串值
	// 4个字符串值都共享相同的支撑数组

	// We don't want to take an address of a string.
	// We know the size of a string ahead of time.
	// -> it has the ability to be on the stack
	// -> not creating allocation
	// -> not causing pressure on the GC
	// -> the string has been designed to leverage value mechanic, to stay on the stack, out of the
	// way of creating garbage.
	// -> the only thing that has to be on the heap, if anything is the backing array, which is the
	// one thing that being shared
	//我们不想采用字符串的地址。
	//我们提前知道字符串的大小。
	//-> 它具有在栈上的能力
	//-> 不创建分配
	//-> 不产生垃圾收集的压力
	//-> 字符串被设计成利用值机制，停留在栈上，避免产生垃圾。
	//-> 堆上唯一需要的东西(如果有的话)，就是支撑数组，它是被共享的
	fmt.Printf("\n=> Iterate over array\n")
	for i, fruit := range strings {
		fmt.Println(i, fruit)
	}

	// Declare an array of 4 integers that is initialized with some values using literal syntax.
	// 使用文本语法声明一个由4个整数组成的数组，该数组使用某些值进行初始化
	numbers := [4]int{10, 20, 30, 40}

	// Iterate over the array of numbers using traditional style.
	// 使用传统方法对数字数组进行循环。
	fmt.Printf("\n=> Iterate over array using traditional style\n")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	// ---------------------
	// Different type arrays
	// 不同类型的数组
	// ---------------------

	// Declare an array of 5 integers that is initialized to its zero value.
	// 声明一个由5个整数组成的数组，该数组初始化为它的零值。
	var five [5]int

	// Declare an array of 4 integers that is initialized with some values.
	// 声明一个由 4个整数组成的数组，该数组用一些值进行初始化。
	four := [4]int{10, 20, 30, 40}

	fmt.Printf("\n=> Different type arrays\n")
	fmt.Println(five)
	fmt.Println(four)

	// When we try to assign four to five like so five = four, the compiler says that
	// "cannot use four (type [4]int) as type [5]int in assignment"
	// This cannot happen because they have different types (size and representation).
	// The size of an array makes up its type name: [4]int vs [5]int. Just like what we've seen
	// with pointer. The * in *int is not an operator but part of the type name.
	//当我们尝试将 four 赋值给 five 时，比如 five = four，编译器会说
	//"cannot use four (type [4]int) as type [5]int in assignment"//“在赋值中不能使用 four (type [4] int)作为 type [5] int”
	//这不允许发生，因为它们具有不同的类型(大小和表示形式)。
	//数组的大小构成了它的类型名: [4] int 与 [5] int。就像在指针里面，* 在 *int 中不是运算符，而是类型名称的一部分。
	

	// Unsurprisingly, all array has known size at compiled time.
	// 这并不奇怪，所有数组在编译时都已知大小。

	// -----------------------------
	// Contiguous memory allocations
	// 连续内存分配
	// -----------------------------

	// Declare an array of 6 strings initialized with values.
	// 声明一个由 6 个字符串组成的数组，使用值进行初始化。
	six := [6]string{"Annie", "Betty", "Charley", "Doug", "Edward", "Hoanh"}

	// Iterate over the array displaying the value and address of each element.
	// By looking at the output of this Printf function, we can see that this array is truly a
	// contiguous block of memory. We know a string is 2 word and depending on computer
	// architecture, it will have x byte. The distance between two consecutive IndexAddr is exactly
	// x byte.
	// v is its own variable on the stack and it has the same address every single time.
	//通过查看这个 Printf 函数的输出，我们可以看到这个数组实际上是一个连续内存块。
	//我们知道字符串是2个字(4字节)，根据计算机架构，它将有x个字节。两个连续IndexAddr之间的距离正好是x字节。
	//v是栈中它自己的变量，每次都有相同的地址。
	fmt.Printf("\n=> Contiguous memory allocations\n")
	for i, v := range six {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &six[i])
	}
}
