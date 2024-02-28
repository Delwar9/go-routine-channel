# Goroutine Example: Printing Even and Odd Numbers

This Go program demonstrates the use of goroutines and channels to print even and odd numbers in sequence. It utilizes channels to coordinate the execution between two goroutines, one for printing even numbers and the other for printing odd numbers.

## How It Works

- Two channels `evenCh` and `oddCh` are created with a buffer size of 1.
- A `sync.WaitGroup` called `wg` is used to coordinate the two goroutines.
- Two goroutines are started:
  - `even(evenCh, oddCh)`: Prints even numbers and signals the `oddCh` channel.
  - `odd(oddCh, evenCh)`: Prints odd numbers and signals the `evenCh` channel.
- The sequence starts by sending a signal to `evenCh` to start the process.
- Each goroutine waits for a signal from the other before printing its number and then signals the other to proceed.
- The main function waits for both goroutines to finish using `wg.Wait()`.

## Usage

1. Clone the repository:

   ```sh
   https://github.com/Delwar9/go-routine-channel.git

2. Navigate to the project directory:

   ```sh
   cd go-routine-channel

3. Run the Go program:
   ```sh
   go run main.go
   
You should see the output displaying the numbers 1 to 20 in order, with even numbers printed by the even goroutine and odd numbers printed by the odd goroutine.

## Requirements
Go 1.13 or higher
