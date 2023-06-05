#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_25) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_27)

&#10;

# Lesson 26

# concurrency - 3

# job que-que example

1. Introduction:

   - The provided code demonstrates the usage of channels for communication between goroutines in Go.
   - It showcases how channels can be used to distribute tasks among worker goroutines efficiently.

2. Understanding the Code:

   - The main function initializes two channels: `jobs` and `results` with a buffer size of 100.
   - The `jobs` channel is used to send task IDs, and the `results` channel is used to receive the corresponding results.
   - A worker goroutine is launched using the `worker` function, which takes the `jobs` and `results` channels as arguments.
   - In the main function, a loop is executed to send 100 tasks to the `jobs` channel using the `jobs <- i` syntax.
   - After sending all the tasks, the `jobs` channel is closed using the `close(jobs)` statement.
   - Another loop is executed to receive and print the results from the `results` channel using `<-results` syntax.

3. Worker Goroutine:

   - The `worker` function takes the `jobs` channel as a receive-only channel (`<-chan int`) and the `results` channel as a send-only channel (`chan<- int`).
   - It loops over the `jobs` channel using `for n := range jobs` syntax, which automatically receives values until the channel is closed.
   - For each received task ID, the `fib` function is called to calculate the Fibonacci number, and the result is sent to the `results` channel using `results <- fib(n)`.

4. Fibonacci Calculation:

   - The `fib` function recursively calculates the Fibonacci number for the given input.
   - If the input `n` is less than or equal to 1, the function returns the input value.
   - Otherwise, it recursively calls itself to calculate `fib(n-1)` and `fib(n-2)` and returns their sum.

5. Running the Example:

   - After understanding the code, you can run it to observe the output.
   - The program distributes the calculation of Fibonacci numbers among worker goroutines using channels.
   - Each worker receives a task ID from the `jobs` channel, calculates the Fibonacci number, and sends the result to the `results` channel.
   - The main function then receives and prints the results from the `results` channel.
   - You should see the Fibonacci numbers printed in the console.

6. Conclusion:
   - The example demonstrates how channels can be used to coordinate and distribute tasks among goroutines efficiently.
   - Channels enable safe and synchronized communication between goroutines, allowing for effective concurrent programming in Go.
