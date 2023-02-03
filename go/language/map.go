package main

import "fmt"

// user defines a user in the program.
type user struct {
	name    string
	surname string
}

func main() {
	// ----------------------
	// Declare and initialize
	// ----------------------

	// Declare and make a map that stores values of type user with a key of type string.
	// 声明并用make函数创建一个 map(映射)，存储 user类型的值(value)与string类型的键(key) 
	users1 := make(map[string]user)

	// Add key/value pairs to the map.
	// 向 map 中添加键/值对.
	users1["Roy"] = user{"Rob", "Roy"}
	users1["Ford"] = user{"Henry", "Ford"}
	users1["Mouse"] = user{"Mickey", "Mouse"}
	users1["Jackson"] = user{"Michael", "Jackson"}

	// ----------------
	// Iterate over map
	// map的循环
	// ----------------

	fmt.Printf("\n=> Iterate over map\n")
	for key, value := range users1 {
		fmt.Println(key, value)
	}

	// ------------
	// Map literals
	// map 的遍历
	// ------------

	// Declare and initialize the map with values.
	users2 := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Iterate over the map.
	fmt.Printf("\n=> Map literals\n")
	for key, value := range users2 {
		fmt.Println(key, value)
	}

	// ----------
	// Delete key
	// 删除键
	// ----------

	delete(users2, "Roy")

	// --------
	// Find key
	// 查找键
	// --------

	// Find the Roy key.
	// If found is True, we will get a copy value of that type.
	// if found is False, u is still a value of type user but is set to its zero value.
	// 如果found为true，则获得一个那个key类型（user类型的值）的拷贝值
	// 如果found为false，则是获得user类型的零值
	u1, found1 := users2["Roy"]
	u2, found2 := users2["Ford"]

	// Display the value and found flag.
	fmt.Printf("\n=> Find key\n")
	fmt.Println("Roy", found1, u1)
	fmt.Println("Ford", found2, u2)

	// --------------------
	// Map key restrictions
	// Map 键约束
	// --------------------

	// type users []user
	// Using this syntax, we can define a set of users
	// This is a second way we can define users. We can use an existing type and use it as a base for
	// another type. These are two different types. There is no relationship here.
	// However, when we try use it as a key, like: u := make(map[users]int)
	// the compiler says we cannot use that: "invalid map key type users"
	// The reason is: whatever we use for the key, the value must be comparable. We have to use it
	// in some sort of boolean expression in order for the map to create a hash value for it.
	// 使用这种语法，我们可以定义一组 user
	// 这是定义users的第二种方式.我们可以使用现有的类型并将其作为另一种类型。
	// 这是两种不同的类型. 它们之间没有关系
	// 但是，当我们尝试使用它作为key时,比如：u: = make (map [ users ] int)
	// 编译器说我们不能使用"invalid map key type users"
	// 原因是：无论我们使用什么key，它的value都是可比较的，我们必须在某种布尔表达式中使用它，以便 map为其创建哈希值
}
