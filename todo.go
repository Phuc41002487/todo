package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const FilePath = "./TodoList.json"
const Undone = "Undone"
const Done = "Done"

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

func CheckFile(ops string) {
    _, err := os.Stat(FilePath)
    if !os.IsNotExist(err) {
        switch ops {
        case "list":
            fmt.Println("There is nothing to do. Chill")
            os.Exit(0)
        case "change":
            fmt.Println("Error: There is no task to change.")
            os.Exit(1)
        }
    }
}

// GetTodos get the current todo list from json file
func GetTodos(todos *Todos) {
    fileInfo, err := os.Stat(FilePath)
    if os.IsNotExist(err) {
        os.Create(FilePath)
        fileInfo, err = os.Stat(FilePath)
        check(err)
    } else {
        check(err)
    }
    file, err := os.Open(FilePath)
    check(err)
    if fileInfo.Size() !=0 {
        byteValue, err := io.ReadAll(file)
        check(err)
        err = json.Unmarshal(byteValue, &todos)
        check(err)
    }
}

// Write todo list to json file
func WriteTodos(todos Todos) {
    file, err := os.Create(FilePath)
    check(err)
    // Ensure the file is closed after the function is completed
    defer file.Close()
    dat, err := json.MarshalIndent(todos, "", "  ")
    check(err)
    fmt.Println(string(dat))
    _, err = file.WriteString(string(dat) + "\n")
    check(err)
}

// Add function add all the task enter in the CLI to todo list.
// At the end print the complete list of task in the console
func Add(args []string) {
    // Get existing task
    var todos Todos
    GetTodos(&todos)
    Outer:
    for _, arg := range args {
        for i := range todos.Todos {
            if todos.Todos[i].Name == arg {
                fmt.Println("Task " + arg + " already existed.")
                continue Outer
            }
        }
        todo := Todo {
            Name: arg,
            Status: "Undone",
        }
        todos.Todos = append(todos.Todos, todo)
    }
    WriteTodos(todos)
}

// Change function change the status of task
// The function accept two parameter:
//   - taskname: used to identify the task to change status
//   - status: status to changed to 
func Change(args []string) {
    if len(args) != 2 {
        fmt.Println(`Error: change task status only support two arguments.
        First arg is task name and Second is the status`)
        os.Exit(1)
    } else {
        CheckFile("change")
        name := args[0]
        status := args[1]
        var todos Todos
        GetTodos(&todos)
        for i := range todos.Todos {
            if todos.Todos[i].Name == name {
                if ((status == Undone) || (status == Done)) {
                    todos.Todos[i].Status = status
                } else {
                    fmt.Printf("Error: status can only be %s or %s.\n", Undone, Done)
                    os.Exit(1)
                }
            }
        }
        WriteTodos(todos)
    }
}

//Delete function delete task specify by the cli arguments
func Delete(args []string) {
    var todos Todos
    GetTodos(&todos)
    if len(todos.Todos) == 0 {
        fmt.Println("There is no task to delete. Chill")
    } else {
        for _, arg := range args {
            result := todos.Todos[:0]
            for _, todo := range todos.Todos {
                if todo.Name != arg {
                    result = append(result, todo)
                }
            }
            todos.Todos = result
        }
        WriteTodos(todos)
    }
}

// List function print the current todo list to console
func List() {
    CheckFile("list")
    data, err := os.ReadFile(FilePath)
    check(err)
    fmt.Println(string(data))
}

// Help function show the information about the app and how to use it
func Help() {
    fmt.Println(`todo is a cli app used to manage task that you want to do.
Currently the todo app support below command
  - help: show information about the app and supported subcommand
  - list: show the current tasks and and their status.
  - add: add task to todo list. task will be added with undone status.
      If your task name include comprised of multiple words. Please enclosed them in double quote.
      You can add multiple task at a time. Please separate the task by using space.
  - change: change the status of the task.
          Please enter your task name first and then the status you want to change it to
          Task status can only Done or Undone.
  - delete: delete a task by task name.`)
}

func main() {
    if len(os.Args) == 1 {
        Help()
    } else {
        switch os.Args[1] {
        case "clear":
            os.Create(FilePath)
        case "add":
            Add(os.Args[2:])
        case "list":
            List()
        case "change":
            Change(os.Args[2:])
        case "delete":
            Delete(os.Args[2:])
        case "help":
            Help()
        default:
            fmt.Println("Invalid Command was enter. " + 
            "Please use help subcommand to know how to use the app")
        }
    }
}