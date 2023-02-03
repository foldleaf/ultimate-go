// ----------------------------------
// https://cn.linkedin.com/in/icyfenix
// Everything is about pass by value.
// 一切都是关于值的传递
// ----------------------------------

// Pointers serve only 1 purpose: sharing.
// Pointers share values across the program boundaries.
// There are several types of program boundaries. The most common one is between function calls.
// We can also have a boundary between Goroutines which we will discuss later.
// 指针只有一个用途: 共享。
// 指针跨程序的边界共享值。
// 程序边界有几种类型，最常见的是函数调用之间的边界。
// 我们也可以在 Goroutines 之间划一条边界，稍后我们将讨论这个问题。


// When this program starts up, the runtime creates a Goroutine.
// Every Goroutine is a separate path of execution that contains instructions that needed to be executed by
// the machine. We can also think of Goroutines as lightweight threads.
// This program has only 1 Goroutine: the main Goroutine.
// 当这个程序启动时，运行时创建一个 Goroutine。
// 每个Goroutine都是一个单独的执行路径，其中包含需要由机器执行的指令。我们也可以将Goroutine视为轻量级线程。
// 本程序只有1个 Goroutine:  main Goroutine.

// Every Goroutine is given a block of memory, called the stack.
// The stack memory in Go starts out at 2K. It is very small. It can change over time.
// Every time a function is called, a piece of stack is used to help that function run.
// The growing direction of the stack is downward.
// 每个 Goroutine 都有一个内存块，称为栈（内存中的栈）。
// Go 中的栈内存从2K 开始，它非常小，可以随时间变化。
// 每次调用函数时，都会使用一段栈来帮助该函数运行。
// 栈(stack)的生长方向是向下的。(从高地址向低地址增长)
// 栈生长方向指的就是执行push、pop(入栈、出栈)命令后，堆栈指针sp(esp)所指向的地址是增大还是减小。向上就是增大，向下就是减小

// Every function is given a stack frame, memory execution of a function.
// The size of every stack frame is known at compile time. No value can be placed on a stack
// unless the compiler knows its size ahead of time.
// If we don't know the size of something at compile time, it has to be on the heap.
// 每个函数都有一个栈帧，栈帧是函数在内存中的执行(区域)。
// 每个栈帧的大小在编译时就已确定。除非编译器事先知道某个值的大小，否则这个值不能放在栈上。
// 如果我们在编译时不知道某个东西的大小，那么它必须在堆(heap)上。

// Zero value enables us to initialize every stack frame that we take.
// Stacks are self cleaning. We clean our stack on the way down.
// Every time we make a function, zero value initialization cleans the stack frame.
// We leave that memory on the way up because we don't know if we would need that again.
// 零值使我们能够初始化所获取的每个栈帧
// 栈是自动清理的.在向下（入栈）的过程中我们会清理栈
// 每次我们创建一个函数时，零值初始化都会清理栈帧。
// 我们在向上(出栈)的过程中保留下了这些内存，因为我们不知道我们是否会再次需要这些内存。

package main

// user represents an user in the system.
type user struct {
	name  string
	email string
}

func main() {
	// -------------
	// Pass by value
	// 值传递
	// -------------

	// Declare variable of type int with a value of 10.
	// This value is put on a stack with a value of 10.
	// 声明int类型的变量，值为10。
	// 这个值(10)被放在栈上
	count := 10

	// To get the address of a value, we use &.
	// 我们用 & 获取值的地址
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" count.
	// 传递 count 的值.
	increment1(count)

	// Printing out the result of count. Nothing has changed.
	// 打印 count 的结果。 没有变化.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	// This is still considered pass by value, not by reference because the address itself is a value.
	// 传递 count 的地址(值).
	// 这仍然被认为是通过值传递的，而不是通过引用传递的，因为地址本身是一个值
	increment2(&count)

	// Printing out the result of count. count is updated.
	// 打印 count 的结果. count 更新了.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// ---------------
	// Escape analysis
	// 逃逸分析
	// ---------------

	stayOnStack()
	escapeToHeap()
}

func increment1(inc int) {
	// Increment the "value of" inc.
	inc++
	println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

// increment2 declares count as a pointer variable whose value is always an address and points to
// values of type int.
// The * here is not an operator. It is part of the type name.
// Every type that is declared, whether you declare or it is predeclared, you get for free a pointer.
// Increment2将 count 声明为一个指针变量，其值总是一个地址并指向 int类型 的值。
// 这里的 * 不是运算符。它是类型名称的一部分。
// 每个类型被声明时，无论是声明的还是预声明的，都可以免费获得一个指针。
func increment2(inc *int) {
	// Increment the "value of" count that the "pointer points to".
	// The * is an operator. It tells us the value of the pointer points to.
	// 增加指针指向的 count 的值
	// 这里的 * 是一个运算符。它告诉我们指针指向的值。
	*inc++
	println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}

// stayOnStack shows how the variable does not escape.
// Since we know the size of the user value at compile time, the compiler will put this on a stack
// frame.
// StayOnStack 展示了变量如何不逃逸(escape)。
// 由于我们知道编译时user值的大小，编译器将把它放在栈帧上
func stayOnStack() user {
	// In the stayOnStack stack frame, create a value and initialize it.
	// 在 stayOnStack 栈帧中，创建一个值并初始化它
	u := user{
		name:  "Hoanh An",
		email: "hoanhan101@gmail.com",
	}

	// Take the value and return it, pass back up to main stack frame.
	// 获取该值并返回它，然后传递回主栈帧。
	return u
}

// escapeToHeap shows how the variable escape.
// This looks almost identical to the stayOnStack function.
// It creates a value of type user and initializes it. It seems like we are doing the same here.
// However, there is one subtle difference: we do not return the value itself but the address
// of u. That is the value that is being passed back up the call stack. We are using pointer
// semantic.
// EseToHeap 显示了变量如何逃逸。
// 这看起来几乎与 stayOnStack 函数一样
// 它创建了一个user类的值并初始化它。这看起来也一样
// 然而，有一个微妙的区别：我们不返回值本身，而是返回u的地址。
// 这是传递回调用栈的值。我们使用的是指针语义。

// You might think about what we have after this call is: main has a pointer to a value that is
// on a stack frame below. If this is the case, then we are in trouble.
// Once we come back up the call stack, this memory is there but it is reusable again. It is no
// longer valid. Anytime now main makes a function call, we need to allocate the frame and
// initialize it.
// 你可能会想，调用后我们得到的是：main有一个指针指向下面栈帧上的值。如果是这样的话，那么问题就来了。
// 一旦我们回到调用栈，这块内存就在那里，但它可以再次使用，但是是失效的效。每当main调用函数时，我们都需要分配帧并初始化它。

// Think about zero value for a second here. It enables us to initialize every stack frame that
// we take. Stacks are self cleaning. We clean our stack on the way down. Every time we make a
// function call, zero value, initialization, we are cleaning those stack frames. We leave that
// memory on the way up because we don't know if we need that again.
// 在这里考虑一下零值。它使我们能够初始化我们获取的每个栈帧。栈是自动清理的。我们在向下（入栈）的过程中清理栈。
// 每次我们进行函数调用、零值初始化时，我们都会清理这些栈帧。
// 我们在向上（出栈）的过程中保留内存，因为不知道是否会再次用到

// Back to the example. It is bad because it looks like we take the address of user value, pass it
// back up to the call stack giving us a pointer which is about to get erased.
// However, that is not what will happen.
// 回到示例。这很糟糕，因为它看起来像是我们获取了user值的地址，将其传递回调用栈，给我们一个即将被擦除的指针。
// 然而，这并不会发生。

// What is actually going to happen is escape analysis.
// Because of the line "return &u", this value cannot be put inside the stack frame for this function
// so we have to put it out on the heap.
// Escape analysis decides what stays on the stack and what does not.
// In the stayOnStack function, because we are passing the copy of the value itself, it is safe to
// keep these things on the stack. But when we SHARE something above the call stack like this,
// escape analysis said this memory is no longer valid when we get back to main, we must put it
// out there on the heap. main will end up having a pointer to the heap.
// In fact, this allocation happens immediately on the heap. escapeToHeap is gonna have a pointer
// to the heap. But u is gonna base on value semantic.
// 真正会发生的是逃逸分析(escape analysis)。
// 因为 "return &u" 这行，这个值不能放在这个函数的栈帧中，所以我们必须把它放到堆(heap)中。
// 逃逸分析决定哪些内容保留在栈上，哪些不保留。
// 在 stayOnStack 函数中，因为我们传递的是值本身的副本，所以把这些东西放在栈上。
// 但是，当我们像这样在调用栈上共享一些东西时，逃逸分析表明，当我们返回main时，这块内存不再有效，
// 我们必须将它放在堆上。main最终将有一个指向堆的指针。
// 实际上，这种分配会立即在堆上进行。
// escapeToHeap将有一个指向堆的指针。但u将是属于值语义的。
func escapeToHeap() *user {
	u := user{
		name:  "Hoanh An",
		email: "hoanhan101@gmail.com",
	}

	return &u
}

// ----------------------------------
// What if we run out of stack space?
// 如果栈空间用完了会发生什么？
// ----------------------------------

// What happens next is during that function call, there is a little preamble that asks "Do we have
// enough stack space for this frame?". If yes then no problem because at compile time we know
// the size of every frame. If not, we have to have bigger frame and these values need to be copied
// over. The memory on that stack moves. It is a trade off. We have to take the cost of this copy
// because it doesn't happen a lot. The benefit of using less memory in any Goroutine outweighs the
// cost.
// 接下来发生的是在函数调用期间,在这之前有个小问题，我们有足够的栈空间给这个帧吗？
// 如果有,那么没有问题,因为在编译时我们就知道每帧的大小
// 如果没有，我们需要更大的帧并将这些值拷贝过来。栈上的内存发生移动
// 我们必须承担这个拷贝的开销，因为它发生的次数不多。
// 在任何Goroutine中使用较少内存的好处大于这种开销

// Because the stack can grow, no Goroutine can have a pointer to some other Goroutine stack.
// There would be too much overhead for the compiler to keep track of every pointer. The latency will
// be insane.
// -> The stack for a Goroutine is only for that Goroutine. It cannot be shared between Goroutines.
// 因为栈可以增长，所以没有 Goroutine 可以拥有指向其他 Goroutine 栈的指针。
// 对于编译器来说，跟踪每个指针会有太多的开销,延迟会很吓人
// -> Goroutine 的栈只适用于那个 Goroutine，不能在 Goroutines 之间共享

// ------------------
// Garbage collection
// 垃圾收集
// ------------------

// Once something is moved to the heap, Garbage Collection has to get in.
// The most important thing about the Garbage Collector (GC) is the pacing algorithm.
// It determines the frequency/pace that the GC has to run in order to maintain the smallest possible t.
// 一旦有东西被移动到堆中，垃圾收集就必须跟进。
// 关于垃圾收集器(GC)最重要的是pacing算法。
// 它确定GC必须运行的frequency/pace，以保持最小的可能时间t

// Imagine a program where you have a 4 MB heap. GC is trying to maintain a live heap of 2 MB.
// If the live heap grows beyond 4 MB, we have to allocate a larger heap.
// The pace the GC runs at depends on how fast the heap grows in size. The slower the
// pace, the less impact it is going to have. The goal is to get the live heap back down.
// 想象一下有一个4MB 堆的程序。GC 试图维护一个2MB 的活动堆。
// 如果活动堆增长超过4MB，我们必须分配一个更大的堆。
// GC 运行的速度取决于堆大小增长的速度,速度越慢，影响越小。其目标是减小活动堆


// When the GC is running, we have to take a performance cost so all Goroutines can keep running
// concurrently. The GC also has a group of Goroutines that perform the garbage collection work.
// 当 GC 运行时，我们必须考虑性能成本，这样所有 Goroutines 都可以保持同时运行
// GC还有一组执行垃圾收集工作的Goroutine。
// 它给自己使用了我们可用 CPU 容量的25% 。
// 更多关于GC和节奏算法的细节见
// https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/README.md
// It uses 25% of our available CPU capacity for itself.
// More details about GC and pacing algorithm can be find at:
// https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/README.md
