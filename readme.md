#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_23) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_25)

&#10;

# Lesson 24

# Concurrency in go (go routines)

1. Introduction to Goroutines:

   - Goroutines are lightweight concurrent functions or "green threads" in Go.
   - They allow concurrent execution of functions independently of the main program flow.
   - Goroutines are an essential part of Go's concurrency model, enabling efficient concurrent programming.

2. Creating Goroutines:

   - To create a goroutine, you need to prefix a function or method call with the `go` keyword.
   - The `go` statement launches a new goroutine to execute the function concurrently.
   - The function is executed in the background, allowing the program to continue its execution.
   - Goroutines are extremely lightweight, and Go can create thousands or even millions of goroutines efficiently.

3. Example of Using Goroutines:

   - Let's consider an example where we calculate the sum of numbers concurrently using goroutines.

   ```go
   package main

   import (
       "fmt"
       "sync"
   )

   func calculateSum(numbers []int, wg *sync.WaitGroup) {
       defer wg.Done()

       sum := 0
       for _, num := range numbers {
           sum += num
       }

       fmt.Printf("Sum of numbers %v: %d\n", numbers, sum)
   }

   func main() {
       numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

       var wg sync.WaitGroup
       wg.Add(2)

       go calculateSum(numbers[:len(numbers)/2], &wg)
       go calculateSum(numbers[len(numbers)/2:], &wg)

       wg.Wait()
   }
   ```

   In this example:

   - We define a function `calculateSum` that takes a slice of integers and calculates their sum.
   - The function receives a `sync.WaitGroup` pointer `wg` to notify when it finishes its execution.
   - In the `main` function, we create a slice of numbers and initialize a `sync.WaitGroup`.
   - We split the numbers slice into two halves and launch two goroutines, each calculating the sum of its respective half.
   - We pass the respective number slices and the `WaitGroup` pointer to the goroutines.
   - After launching the goroutines, we call `wg.Wait()` to wait for both goroutines to complete.
   - Finally, the main program waits until both goroutines finish their execution.

4. Running the Example:

   - After understanding the code, you can run it to observe the output.
   - The program calculates the sum of numbers concurrently using goroutines.
   - The two goroutines work on different halves of the numbers slice concurrently, speeding up the calculation.
   - The `sync.WaitGroup` ensures that the main program waits for both goroutines to finish before exiting.
   - You should see the sum of each half printed as output.

5. Conclusion:
   - Goroutines are a powerful feature in Go that allows concurrent programming with ease.
   - They enable efficient utilization of system resources and simplify concurrent tasks.
   - By leveraging goroutines, you can write scalable, responsive, and efficient concurrent programs in Go.
