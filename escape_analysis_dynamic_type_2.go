package main

func Test2() []int {
	l := 20
	c := make([]int, 0, l) // 堆 动态分配不定空间 逃逸
	return c
}

func main() {
	Test2()
}