#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_21) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_23)

&#10;

# Lesson 22

- The `defer` keyword in Go is used to schedule a function call to be executed later, typically when the surrounding function exits.
- The deferred function call is placed on a stack and executed in a last-in-first-out (LIFO) order.
- Any arguments to the deferred function are evaluated immediately at the time the `defer` statement is encountered.
- The `defer` statement is often used to ensure that certain cleanup tasks or finalization steps are executed, regardless of whether the function completes normally or panics.
- Common use cases for `defer` include closing files, releasing resources, unlocking locks, or logging function entry and exit.
- It is important to note that deferred functions are not executed immediately when encountered, but rather when the surrounding function has finished its execution.
- If a function contains multiple `defer` statements, they will be executed in the reverse order in which they were defined.
- `defer` statements can be placed anywhere within a function and can even be nested within control flow statements like loops or conditionals.
- When a function panics, the deferred functions will still be executed before the panic is propagated up the call stack.
- It is possible to modify the return values of a function within a deferred function, but it is generally not recommended as it can lead to confusion and unexpected behavior.

Remember, the `defer` keyword provides a convenient way to ensure that certain actions are always taken, even in the presence of errors or panics, making your code more robust and maintainable.
