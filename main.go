package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello, World!")

	//str := "n/ah3air3k"
	str := "hello, world!"

	chars := []rune(str)
	var stringa string

	for i := len(chars) - 1; i >= 0; i-- {
		//fmt.Println(str)
		char := chars[i]

		// ascii 96 beetween 123 to lowercase characters
		if chars[i] > 96 && chars[i] < 123 {
			stringa += string(char)
		}

	}
	fmt.Println(stringa)

}
