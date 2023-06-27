/*
Test Go project

[Resources]
+ Golang Book : https://www.golang-book.com/public/pdf/gobook.pdf
*/
package main // Define package name; similar to Java

// Import dependencies/external libraries
import (
    // 'Format' package, similar to stdio for I/O operations
    "fmt"
)

/*
SCOPE: Global Variables vs Local Variables
*/
// Define global variables
var global_var string = "Hello World Global"

// Test Functions for Learning
func print_standard_output() {
    // Similar to Java or C, Println prints to standard output
    fmt.Println("Hello World") // Print string
    fmt.Println("Hello World"[1]) // Get the 2nd element in the string and print the index byte
    fmt.Println("Hello " + "World") // Concatenating string
}

func print_truth_table() {
    /*
    Truth Table

    [AND]
    ======================
    TRUE  | TRUE  | TRUE
    TRUE  | FALSE | FALSE
    FALSE | TRUE  | FALSE
    FALSE | FALSE | FALSE
    ======================

    [OR]
    ======================
    TRUE  | TRUE  | TRUE
    TRUE  | FALSE | TRUE
    FALSE | TRUE  | TRUE
    FALSE | FALSE | FALSE
    ======================
    */
    fmt.Println(true && true) // TRUE
    fmt.Println(true && false) // FALSE
    fmt.Println(true || true) // TRUE
    fmt.Println(true || false) // TRUE
    fmt.Println(!true)        // False
}

func print_Variables() {
    /*
    Golang requires the use of keywords to define the data object that is going to be created 
        + 'var' keyword to define that it is a variable, followed by defining the variable name and the type
        + 'const' keyword to create a constant/immutable object that cannot be changed during runtime after compile time
    Golang also supports dynamic-typed language syntax, compared to static typing
        - this means golang will interpret the context and infer the typing based on the literal provided, similar to python and Javascript
    Golang can also define multiple variables at the same time through a 'key-value' syntax

    [Keywords]
    const [variable-name] [data-type]
    var [variable-name] [data-type] : To create a variable
    */
    // Define/Declare variables
    var str_var string // Statically define language
    const const_var string = "Constant Value" // Constant variable

    // Defining multiple variables
    var (
        // [variable-name] = [value]
        a = 5
        b = 10
        c = 15
        d = "Hello Multiple Variables"
        e = 0.5
    )

    // Assign values/variables, Process
    str_var = "Hello Static World"
    dynamic := "Hello Dynamic World" // Defining a variable with a literal; Data typing will be inferred based on the value provided
    // const_var = "New Value" // Trying to change a constant value will result in an error

    // Print output
    fmt.Println(str_var) // Print statically-defined variable
    fmt.Println(dynamic) // Print dynamically-defined variable
    fmt.Println(global_var) // Print global variable
    fmt.Println(const_var) // Constant value, cannot be changed

    // Print multi-key variable
    fmt.Println(a, ", ", b, ", ", c, ", ", d, ", ", e)
}

func get_user_Input() {
    /*
    Getting user input

    Similar to using C or Java, 
        - use 'fmt.Scanf("data-type-format", &[reference-variable])' to get user input and put the value into the variable
    */
    // Define variables
    var input_Float float64

    // Get float value from user input and put the result into the variable 'float64' via reference
    fmt.Print("Please Enter a number: ")
    fmt.Scanf("%f", &input_Float)

    // Process input
    output := input_Float * 2 // Multiply result by 2

    // Output result
    fmt.Println(output)
}

func for_loop() {
    /*
    For loop used to run an interation 

    [Notes]
    - Unlike python,c/c++ or any language,
        - Golang does not have a while loop, do-while, while-do etc
            - The for loop is the loop objects
    */

    // Declare variables
    max_iter := 10 // Maximum Iteration

    fmt.Println("===========")
    fmt.Println("While loop ")
    fmt.Println("===========")
    // While Loop 10 times
    i := 1  // Starting Index
    for i <= max_iter {
        fmt.Println(i)
        i = i+1
    }

    fmt.Println("==============")
    fmt.Println("'For' for-loop")
    fmt.Println("==============")
    // "For" for-loop
    for i:=0; i < max_iter; i++ {
        fmt.Println(i)
    }

    fmt.Println("====================")
    fmt.Println(" Key-Value For Loop ")
    fmt.Println("====================")
    var total float64 = 0
    // Dictionary-like, Key-value for loop
    arr := []float64{98,93,77,82,83} // Initialize an array/list of type <float64>
    // NOTE: 
    //  - Replace 'i' with '_' to tell the compiler that you dont need this
    //  - 'value' is the same as 'arr[i]'
    for _, value := range arr {
        // Loop it [arr] number of times
        // Value
        // Increment total by the value
        total += value
        fmt.Println("Total:", total)
    }
}

func if_condition() {
    /*
    Boolean comparison operator
    */
    fmt.Println("=================")
    fmt.Println("Check even or odd")
    fmt.Println("=================")
    // Check even or odd
    for i:=1; i<= 10; i++ {
        if i % 2 == 0 {
            // Even
            fmt.Println(i, "is even.")
        } else {
            // Odd
            fmt.Println(i, "is odd.")
        }
    }

    fmt.Println("=====================")
    fmt.Println("Check if is divisible")
    fmt.Println("=====================")
    // Check even or odd
    target := 5
    for i:=1; i <= 10; i++ {
        if target % i == 0 {
            // Even
            fmt.Println(target, "is divisible by", i)
        } else {
            // Odd
            fmt.Println(target, "is not divisible by", i)
        }
    }
}

func switch_case() {
    /*
    switch-case option chooser
    */

    fmt.Println("=====================")
    fmt.Println(" Switch-case Integer ")
    fmt.Println("=====================")
    var result_int int = 0
    switch result_int {
    case 0: 
        fmt.Println("Zero")
    case 1:
        fmt.Println("One")
    default: fmt.Println("Invalid Result.")
    }

    fmt.Println("=====================")
    fmt.Println(" Switch-case String  ")
    fmt.Println("=====================")
    var result_str string = "Hello"
    switch result_str {
    case "Hello":
        fmt.Println("World")
    default: fmt.Println("Invalid String.")
    }
}

func arrays() {
    /*
    Arrays, Slices and Maps

    [Arrays]
    Structure: var variable_name [array-size]<data-type>
    */

    fmt.Println("============")
    fmt.Println(" Arrays     ")
    fmt.Println("============")
    // Declare variables
    var arr [5]int // Create an integer array of size 5

    // Initialize and Assigning value
    for i:=0; i<5; i++ {
        arr[i] = i+1 // Assigning the 5th element in the array with the value '100'
    }

    // Output
    fmt.Println(arr)

    // Slice
    fmt.Println("============")
    fmt.Println(" Slice      ")
    fmt.Println("============")
    // Declare Variables
    arr_2 := [5]float64{1,2,3,4,5}

    // To create a slice: array[lower_bound:upper_bound]
    var lower_bound int = 0
    var upper_bound int = 2
    x := arr_2[lower_bound:upper_bound] // Get the 1st element to the 6th element of the array

    fmt.Println("Array Result: ", x)
}

// Declare Function
func main() {
    fmt.Println("=====================")
    fmt.Println(" 1. Standard Output  ")
    fmt.Println("=====================")
    print_standard_output()

    fmt.Println("=====================")
    fmt.Println(" 2. Truth Table      ")
    fmt.Println("=====================")
    print_truth_table()

    fmt.Println("=====================")
    fmt.Println(" 3. Using Variables  ")
    fmt.Println("=====================")
    print_Variables()

    fmt.Println("=====================")
    fmt.Println(" 4. Standard Input   ")
    fmt.Println("=====================")
    get_user_Input()

    fmt.Println("=====================")
    fmt.Println(" 5. For loop         ")
    fmt.Println("=====================")
    for_loop()

    fmt.Println("=====================")
    fmt.Println(" 6. If condition     ")
    fmt.Println("=====================")
    if_condition()

    fmt.Println("=====================")
    fmt.Println(" 7. Switch Case      ")
    fmt.Println("=====================")
    switch_case()

    fmt.Println("============================")
    fmt.Println(" 8. Arrays, Slices and Maps ")
    fmt.Println("============================")
    arrays()
}
