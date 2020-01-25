package main


import (
	"fmt"
	"github.com/tcnksm/go-input"
	"os"
	"strconv"
)


func fibonacci(n int) int{

	if n <=1 {
		return n

	}else {

		return fibonacci(n - 1 ) + fibonacci(n - 2)
	}

}

func main(){


	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}


	query := "Please enter an integer"

	number, err := ui.Ask(query, &input.Options{
		Required: true,
		Loop:     false,
	})

	n, err := strconv.Atoi(number)
	if err == nil {
		fmt.Println(fibonacci(n))
	}

}