package main

import (
	//"fmt"
	"fmt"
	"os"
)

const FilePath = "./TodoList.json"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func add(args []string) {
    file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_CREATE, 0644)
    // Ensure the file is closed after the function is completed
    defer file.Close()
    check(err)
    for _, arg := range args {
        // 0644 mean Owner read write, Everyone Read
        _, err := file.WriteString(arg + "\n")
        check(err)
    }
}

func list() {
    data, err := os.ReadFile(FilePath)
    check(err)
    fmt.Println(string(data))
}

func main() {
    switch os.Args[1] {
    case "clear":
        os.Create(FilePath)
    case "add":
        add(os.Args[2:])
    case "list":
        list()
    }
}