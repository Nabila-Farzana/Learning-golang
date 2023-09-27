// I function the code will be excuted line by line. But when we use defer keyword whatever the line we put for the execution it won't execute then , it will be executed at the very end of the function executed

// Each time a defer statement executes, the function value and parameters to the call are evaluated as usual and saved anew but the actual function is not invoked. Instead, defered functions are invoked immediately before the surrounding function returns, in the reverse order they were defered (LIFO)

package main
import "fmt"

func main() {
	defer fmt.Println("World") // the output will be Hello world again, as printing "World" is defered so it will be executed at the last 

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")// the output will be Hello, Three, Two, One, World. Because More than one defer will follow the LIFO``
	fmt.Println("Hello")
	dufferDefer() // output: Hello, 43210, Three, Two, One, World
}
//Defer stack: world, One, Two, Three
//Defer stack : 0, 1, 2, 3, 4
func dufferDefer(){
	for i:=0; i<5; i++{
		fmt.Print(i)
	}
}
