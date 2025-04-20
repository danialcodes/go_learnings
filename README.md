# Basic Learning and hello world projects

## This branch contains basic learning and hello world projects for the go programming languages.

 - The entry point for the project is `main.go` file.
 - Every go file should have a package declaration at the top, like `package main`
 - There are builtin packages like `fmt`, `math`, `time`, etc. that can be imported using the `import` statement. It's kind of like `#include` in C/C++.
 - To run a go program, we can use the `go run` command followed by the file name. For example, `go run main.go`. It's similar to `python main.py` in Python or `npm start` in Node.js.

# Functions

## There are different types of functions in Go:
- Standard or Named Function:
    - A function with a name that can be called from anywhere in the program.
    - Example:
      ```go
      func add(a int, b int) int {
            return a + b
      }
      ```
- Init Function:
    - A special function that is called automatically when the package is initialized.
    - It is used to initialize variables or perform setup tasks.
    - Takes no parameters and returns no values.
    - It is executed before the main function and is not called explicitly.
    - Example:
      ```go
      func init() {
            fmt.Println("This is the init function")
      }
      ```
- Anonymous Function:
    - A function without a name that can be defined and called inline.
    - It can be assigned to a variable or passed as an argument to another function.
    - Example:
      ```go
      func() {
            fmt.Println("This is an anonymous function")
      }()
      ```
- First-Order Function:
    - Normal functions that can be assigned to variables, passed as arguments, or returned from other functions.
    - Example:
      ```go
      func add(a int, b int) int {
            return a + b
      }
      func main() {
            result := add(2, 3)
            fmt.Println(result) // Output: 5
      }
      ```
- Higher-Order Function:
    - A function that takes another function as an argument or returns a function as a result.
    - Example:
      ```go
      func applyFunction(f func(int, int) int, a int, b int) int {
            return f(a, b)
      }
      ```

- Closure
    - A function that captures the lexical scope in which it was defined, allowing it to access variables from that scope even after the scope has exited.
    - Example:
      ```go
      func makeCounter() func() int {
            count := 0
            return func() int {
                  count++
                  return count
            }
      }
      func main() {
            counter := makeCounter()
            fmt.Println(counter()) // Output: 1
            fmt.Println(counter()) // Output: 2
      }
      ```
- Receiver Function:
    - A function that is associated with a specific type (struct) and can be called on instances of that type.
    - It allows you to define methods for your custom types.
    - Example:
      ```go
      type Circle struct {
            radius float64
      }
      func (c Circle) Area() float64 {
            return math.Pi * c.radius * c.radius
      }
      func main() {
            circle := Circle{radius: 5}
            fmt.Println(circle.Area()) // Output: 78.53981633974483
      }
      ```
- Variadic Function:
    - A function that can accept a variable number of arguments of a specific type.
    - It is defined using the `...` syntax before the parameter type.
    - Example:
      ```go
      func sum(numbers ...int) int {
            total := 0
            for _, num := range numbers {
                  total += num
            }
            return total
      }
      func main() {
            fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15
      }
      ```
# Slice
### A slice is a dynamically-sized, flexible view into the elements of an array. It is a reference type that provides a way to work with a portion of an array without copying the entire array.
### Slices are more powerful than arrays in Go because they can grow and shrink in size, and they provide built-in functions for manipulation. They are often used to represent collections of data.
- Slices are created using the `make` function or by slicing an existing array or slice. They can be passed to functions, returned from functions, and used in various operations.
- Slices are defined by a pointer to the underlying array, a length, and a capacity. The length is the number of elements in the slice, while the capacity is the maximum number of elements that can be stored in the underlying array without reallocating memory.
 ### Example:
```go
package main
import "fmt"
func main() {
    // Creating a slice using the make function
    slice1 := make([]int, 5) // Length: 5, Capacity: 5
    fmt.Println(slice1)       // Output: [0 0 0 0 0]

    // Creating a slice from an array
    arr := [5]int{1, 2, 3, 4, 5}
    slice2 := arr[1:4] // Slicing from index 1 to index 4 (exclusive)
    fmt.Println(slice2) // Output: [2 3 4]

    // Appending elements to a slice
    slice3 := []int{1, 2, 3}
    slice3 = append(slice3, 4, 5)
    fmt.Println(slice3) // Output: [1 2 3 4 5]

    // Slicing a slice
    slice4 := slice3[1:4] // Slicing from index 1 to index 4 (exclusive)
    fmt.Println(slice4)   // Output: [2 3 4]
}
```