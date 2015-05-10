package main

import (
	"fmt"
	"github.com/user/stringutil"
	//"reflect"
	"runtime"
	"strings"
)

// fibonacci is a function that returns
// a function (closure) that returns an int.
func fibonacci() func() int {
	val1 := 1
	val2 := 0
	val3 := 1

	return func() int {
		r := val1 + val2
		val1 = val2
		val2 = val3
		val3 = val1 + val2
		return r
	}
}

func runsOn() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Printf("OS X, ")
	case "linux":
		fmt.Printf("Linux, ")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s, ", os)
	}
	fmt.Print("so Go on and play!\n")
}

func WordCount(s string) map[string]int {
	myfields := strings.Fields(s)
	l := len(myfields)
	m := make(map[string]int)

	for i := 0; i < l; i++ {
		m[myfields[i]]++
	}

	return m
}

func main() {
	fmt.Printf("\n")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	runsOn()
	fmt.Printf("\n")

	fmt.Printf("Let's test the WordCount() function\n(as suggested by https://tour.golang.org/moretypes/19):\n")
	s := "My dog is sad. My dog has fleas. My dog has no food."
	fmt.Printf("\tHere is a string:\n\t\t%s", s)
	fmt.Printf("\n\tHere is the string's word count as a map:\n\t\t")
	fmt.Println(WordCount(s))
	fmt.Printf("\n")

	var nfib int = 20
	fmt.Printf("Here are the first %d Fibonacci numbers implemented with a closure\n", nfib)
	fmt.Printf("(https://tour.golang.org/moretypes/22):\n\t")
	var f func() int
	f = fibonacci()
	//fmt.Println(reflect.TypeOf(f))
	for i := 0; i < nfib; i++ {
		fmt.Printf("%d ", f())
	}
	fmt.Printf("\n\n")
}
