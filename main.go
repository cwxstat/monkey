package main

import (
	"fmt"

	"time"
)

func MyOut(s string) {
	fmt.Printf("this is out: %s\n", s)

}

func main() {

	fmt.Println("here")
	time.Sleep(1 * time.Second)
}
