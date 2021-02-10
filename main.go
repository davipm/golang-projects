package main

import (
	"fmt"
	"math"
	"time"
)

func helloTest() string {
	return "Hello World!"
}

func helloWorld() {
	fmt.Println("Hello World!")
}

func valuesExample() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variablesExample() {
	var a = "initial"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

const s string = "constant"

func constantsExample() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

func forExample() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func ifElseExample() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func switchExample() {
	var i = 2
	fmt.Println("Write", i, "as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's weekday")
	}
}

func arrayExample() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("set", a[4])

	fmt.Println("len", len(a))

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println("dlc", b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d", twoD)
}

func sliceExample() {
	var s = make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("set", s[2])

	fmt.Println("len:", len(s))
}

// main application
func main() {
	helloWorld()
	valuesExample()
	variablesExample()
	constantsExample()
	forExample()
	ifElseExample()
	switchExample()
	arrayExample()
	sliceExample()
}
