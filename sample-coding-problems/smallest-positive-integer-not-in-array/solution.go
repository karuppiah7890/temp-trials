package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    positiveIntegers := ignoreNegativeIntegers(A)

    if len(positiveIntegers) == 0 {
        return 1
    }

    return 0
}

func ignoreNegativeIntegers(A []int) []int {
    positiveIntegers := []int{}
    for _, value := range A {
        if value > 0 {
            positiveIntegers = append(positiveIntegers, value)
        }
    }
    return positiveIntegers
}
