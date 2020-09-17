# Day 1 notes

So I found that we can run input in Golang with 2 Modules/Libraries. fmt and bufio. However it was better with bufio as fmt methods were not helping. I am sure I am missing something but I am not sure what.

1. [fmt](https://golang.org/pkg/fmt/)
   
   - We have 3 function/methods we can use in fmt to read input from terminal.
     
     - [Scan](https://golang.org/pkg/fmt/#Scan)- takes text input and stores space-separated values into successive arguments.
       
       ```go
       var input1 string
       var input2 string
       fmt.Scan(&input1 &input2) // input is "Hello World"
       fmt.Println(input1) // output is "Hello"
       ```
     
     - [Scanf](https://golang.org/pkg/fmt/#Scanf) - Works like ftm.PrintF allowing you to format input. I found the paramenter "%q" should ignore spaces however when I tested it failed. I am still checking what should be the correct method.
       
       ```go
       var input1 string
       fmt.Scanf("%q", &input1) // input is "Hello World"
       fmt.Println(input1) // output is "Hello"
       ```
     
     - [Scanln](https://golang.org/pkg/fmt/#Scanln) - Scan input until newline however it also stops at next space instead of end of line
       
       ```go
       var input1 string
       fmt.Scanln(&input1) // input is "Hello World"
       fmt.Println(input1) // output is "Hello"
       ```

2. [Bufio](https://golang.org/pkg/bufio/) 
   
   Create a reader to read terminal input with the bufio module using the scanner function.
   
   ```go
   var input string
   
   scanner := bufio.NewScanner(os.Stdin)
   if scanner.Scan() {
       input = scanner.Text() // input "Hello World"
   }
   fmt.Println(input) // output is "Hello World"
   ```

After reading and looking up I am not sure if there is a bug in fmt.Scanln function however it is recommended to use bufio module as it has better performance.



one thing to note this was tested on hackerank and on my local windows 10 machine with latest golang v go1.15.2
