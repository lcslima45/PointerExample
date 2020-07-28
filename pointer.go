package main

import "fmt"

func main() {
	var example int     //names of variables tend to be written as nouns
	example = 1         //example is int type by implicity attribution
	pointer := &example //pointer is inialized as the memory address that points where example store the content that is 1 now
	fmt.Println(example, *pointer)
	*pointer = 2
	fmt.Println(example, *pointer) //exit 2 2, as example and *pointer (that is the content of the address stored in pointer) go to the same content 2
	doSomething(&example)          //now that you passed the reference of the variable to the function, the content stored has changed
	fmt.Println(example, *pointer) //now example and *pointer will get exit 3 3
}

//names of functions tend to be written as verbs
func doSomething(enter *int) {
	*enter = 3
	return
}
