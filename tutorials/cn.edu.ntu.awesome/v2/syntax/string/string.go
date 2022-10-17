package string

import "fmt"

func Traverse(str string) {

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	fmt.Println("----------")

	for _, v := range str {
		fmt.Printf("%c\n", v)
	}
}
