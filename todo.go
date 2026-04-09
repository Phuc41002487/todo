package main

import (
	"fmt"
	"os"
    "encoding/json"
)

const FilePath = "./TodoList.json"

type Todo struct {
    Name string
    Status string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func add(args []string) {
    // 0644: Owner read write, Everyone read
    file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_CREATE, 0644)
    // Ensure the file is closed after the function is completed
    defer file.Close()
    check(err)
    
    for _, arg := range args {
        todo := Todo {
            Name: arg,
            Status: "Undone",
        }
        dat, err := json.Marshal(todo)

        _, err = file.WriteString(string(dat) + "\n")
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