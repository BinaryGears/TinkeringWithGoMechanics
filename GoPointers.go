/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import "fmt"

//A function defined and declared
func pointerChangeValue(i int) {
	p := &i
	*p = *p + *p
	fmt.Println("Printing from inside the function the value of i", i)
}

func main() {

	//Declare a variable with explicit type
	var i int = 1

	//Declare a pointer to variable i
	p := &i

	fmt.Println("i variable before the change: ", i)
	fmt.Println("The address of the i variable: ", p)
	fmt.Println("The value of the i variable obtained by the pointer to i: ", *p)

	//Set the value of the variable i from the pointer p;
	*p = 2

	fmt.Println("i variable after the change", i)
	fmt.Println("The address of the i variable: ", p)
	fmt.Println("The value of the i variable obtained by the pointer to i: ", *p)

	//Function call
	pointerChangeValue(i)

}
