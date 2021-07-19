package main

func Slice1() {
	s := make([]int, 10000, 10000)

	for index, _ := range s {
		s[index] = index
	}
}

func main() {
	Slice1()
}