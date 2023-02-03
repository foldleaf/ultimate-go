package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// user is a struct type that declares user information.
type user struct {
	ID   int
	Name string
}

// updateStats provides update stats.
type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {
	// Retrieve the user profile.
	u, err := retrieveUser("Hoanh")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the user profile
	// Since the returned u is an address, use * to get the value.
	fmt.Printf("%+v\n", *u)

	// Update user name. Don't care about the update stats.
	// This _ is called blank identifier.
	// Since we don't need anything outside the scope of if, we can use the compact syntax.
	// 更新 user的 name. 不需要关系 stats的更新.
	// 这个 _ 是空标识符
	// 因为我们不需要if作用域之外的东西，所以可以用紧凑的语法.
	if _, err := updateUser(u); err != nil {
		fmt.Println(err)
		return
	}

	// Display the update was successful.
	fmt.Println("Updated user record for ID", u.ID)
}

// retrieveUser retrieves the user document for the specified user.
// It takes a string type name and returns a pointer to a user type value and bool type error.
// 它接收一个 string 类型的形参 name，返回一个指向 user 类型的指针以及bool类型的 error.
func retrieveUser(name string) (*user, error) {
	// Make a call to get the user in a json response.
	// 调用 getUser获取 user ，得到 json 响应（string类型）.
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	// Goal: Unmarshal the json document into a value of the user struct type.
	// Create a value type user.
	// 目的: 将 json 文档解析为 user 结构体类型的值.
	var u user

	// Share the value down the call stack, which is completely safe so the Unmarshal function can
	// read the document and initialize it.
	// 在调用栈中共享值，这是完全安全的，因此Unmarshal函数可以读取JSON文档并初始化它。
	err = json.Unmarshal([]byte(r), &u)

	// Share it back up the call stack.
	// Because of this line, we know that this create an allocation.
	// The value is the previous step is not on the stack but on the heap.
	// 将其共享回调用堆栈。
	// 这一行，我们知道进行了一次分配。
	// 该赋值为前一步，不是在栈上而是在堆上
	return &u, err
}

// GetUser simulates a web call that returns a json
// document for the specified user.
func getUser(name string) (string, error) {
	response := `{"ID":101, "Name":"Hoanh"}`
	return response, nil
}

// updateUser updates the specified user document.
// GetUser 模拟 Web 调用返回指定user的 JSON文档
func updateUser(u *user) (*updateStats, error) {
	// response simulates a JSON response.
	response := `{"Modified":1, "Duration":0.005, "Success" : true, "Message": "updated"}`

	// Unmarshal the json document into a value of the userStats struct type.
	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	// Check the update status to verify the update is successful.
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}
