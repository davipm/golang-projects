package main

import (
	"fmt"
	"math"
	"time"
)

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

	var f = "apple"
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

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd", s)

	var c = make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy", c)

	var l = s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	var t = []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	var twoD = make([][]int, 3)

	for i := 0; i < 3; i++ {
		var innerLen = i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d", twoD)
}

func distance(name1 string, name2 string) int {
	if len(name1) != len(name2) {
		panic("Names shoeld be same lenght")
		return 0
	}

	var distance = 0

	for i := 0; i < len(name1); i++ {
		if name1[i] != name2[i] {
			distance += 1
		}
	}

	return distance
}

func factorial(number int) int {
	var result = 1

	for i := 2; i <= number; i++ {
		result *= i
	}

	return result
}

func mapExamples() {
	var m = make(map[string]int)

	m["k7"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1", v1)
	fmt.Println("len", len(m))

	delete(m, "k2")
	fmt.Println("map", m)
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)
}

func rangeExample() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	ksv := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range ksv {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range ksv {
		fmt.Println("key", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

// functions examples
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func functionsResults() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}

// Multiple Return Values
func vals() (int, int) {
	return 3, 7
}

// Variadic Functions
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

// Closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

// Pointers
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// Structs
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// Methods
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
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
	fmt.Println(distance("david", "deved"))
	fmt.Println("factorial result ", factorial(5))
	mapExamples()
	rangeExample()
	functionsResults()

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals()
	fmt.Println(c)

	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())

	fmt.Println(fact(7))

	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("zeroptr:", &i)

	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "bob", age: 20})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Fred", age: 40})
	fmt.Println(newPerson("Jhon"))
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)

	r := rect{width: 10, height: 5}
	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim())
}
