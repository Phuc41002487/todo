package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const FilePath = "./TodoList.json"

type Todos struct {
    Todos []Todo `json:"todos"`
}

type Todo struct {
    Name string `json:"name"`
    Status string `json:"status"`
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
    fileInfo, err := os.Stat(FilePath)
    var todos Todos

    if fileInfo.Size() !=0 {
        fmt.Println("go here")
        byteValue, err := io.ReadAll(file)
        check(err)
        err = json.Unmarshal(byteValue, &todos)
    }

    file, err = os.Create(FilePath)
    check(err)
    
    for _, arg := range args {
        todo := Todo {
            Name: arg,
            Status: "Undone",
        }
        todos.Todos = append(todos.Todos, todo)
        dat, err := json.MarshalIndent(todos, "", "  ")
        check(err)
        fmt.Println(string(dat))
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