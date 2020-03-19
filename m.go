package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//fmt.Println(os.Args, len(os.Args))
	fmt.Printf("%q\n", os.Args)

	fmt.Println(os.Args[1])
	defer file.Close()
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0666)

	if err != nil {
		log.Fatalf("Ups %q", os.Args)
	}

	/**
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hola")
	f.Close()
	/**
	input := make([]int, 100)
	output := make([]int, 100)
	for i := range input {
		input[i] = i
		output[i] = Sum(i, i)
	}

	fmt.Println(input)
	c1 != make (chan int)
	go func(){
		c1<-4
	}**/

}
