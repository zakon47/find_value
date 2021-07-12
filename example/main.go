package main

import (
	"fmt"
	"github.com/zakon47/find_value"
)

func main() {
	fmt.Println("=========> NUMBER:")
	fmt.Println(find_value.Number("lorem test text33txt"))
	fmt.Println(find_value.Number("33lorem test text"))
	fmt.Println()
	fmt.Println("=========> NumberReverse:")
	fmt.Println(find_value.NumberReverse("lorem test text33"))
	fmt.Println(find_value.NumberReverse("33lorem test text"))
	fmt.Println()
	fmt.Println("=========> String:")
	fmt.Println(find_value.String("lorem test text33txt"))
	fmt.Println(find_value.String("33lorem test text"))
	fmt.Println()
	fmt.Println("=========> StringReverse:")
	fmt.Println(find_value.StringReverse("lorem test text33txt"))
	fmt.Println(find_value.StringReverse("33lorem test text"))
}
