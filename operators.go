package main

func main(){
	// arithmetic operators
	a := 10
	b := 5
	c:= a + b // addition
	d := a - b // subtraction
	e := a * b // multiplication
	f := a / b // division
	g := a % b // modulus

	println("Addition:", c)
	println("Subtraction:", d)
	println("Multiplication:", e)
	println("Division:", f)
	println("Modulus:", g)

	// comparison operators
	println("Equal to:", a == b) // false
	println("Not equal to:", a != b) // true
	println("Greater than:", a > b) // true
	println("Less than:", a < b) // false
	println("Greater than or equal to:", a >= b) // true
	println("Less than or equal to:", a <= b) // false

	// logical operators
	x := true
	y := false

	println("Logical AND:", x && y) // false
	println("Logical OR:", x || y) // true	
	println("Logical NOT:", !x) // false
	
}