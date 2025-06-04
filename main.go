package main

import "fmt"

func HelloWorld(name string) string {
    return fmt.Sprintf("Hi, %s!", name)
}

func main() {
    fmt.Println(HelloWorld("World"))
}
