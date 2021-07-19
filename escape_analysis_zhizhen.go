package main

type User struct {
	Name 	string
}

func GetUserInfo1() User {
	user := User{
		Name: "wangkun",
	}
	return user
}

func GetUserInfo2() *User {
	user := User{
		Name: "wangkun",
	}
	return &user
}
func GetUserInfo3() *User {
	user := &User{
		Name: "wangkun",
	}
	return user
}

func main() {
	u1 := GetUserInfo1()
	u2 := GetUserInfo2()
	u3 := GetUserInfo3()
	println(u1.Name)
	println(u2.Name)
	println(u3.Name)
}
