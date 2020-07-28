package main

import "fmt"

func main() {
	var example int     //name of variables tend to be written as nouns
	example = 1         //example is int type by implicity
	pointer := &example //pointer is inialized as the memory address that points where example store the content that is 1 now
	fmt.Println(example, *pointer)
	*pointer = 2
	fmt.Println(example, *pointer) //exit 2 2, as example and *pointer (that is the content of the address stored in pointer) go to the same content 2
	doSomething(&example)
	fmt.Println(example, *pointer)
}

//name of functions tend to be written as verbs
func doSomething(enter *int) {
	*enter = 3
	return
}
