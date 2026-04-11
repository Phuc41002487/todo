#!/bin/bash
echo "Build todo app."
go build todo.go

# If even one test fails, the end result is failed
all_passed=true

# Clear task
printf "Clear task\n\n"
./todo.exe clear

# Helper to assert console output
assert_contains() {
  echo "$1" | grep -q "$2"
}

# helper to print status of test
pass() {
    printf "PASS\n\n"
}

fail() {
    printf "FAIL\n\n"
}


# Test list - Expect no task
echo "* Test list subcommand with empty list *"
if ./todo.exe list | grep -q "There is nothing to do. Chill"; then
    pass
else
    fail
    all_passed=false
fi

# Test add - add 3 task to todo list
# First add 1 task using by input 1 argument
# Second add 2 taks using by 2 arguments
echo "* Test add task to Todo list *"
./todo.exe add "learn Go" > /dev/null
./todo.exe add "learn Python" "learn C" > /dev/null
out=$(./todo list)
if assert_contains "$out" "^learn C | Undone$" &&
   assert_contains "$out" "^learn Python | Undone$" &&
   assert_contains "$out" "^learn Go | Undone$"; then
  pass
else
  fail
  all_passed=false
fi

# Test change status of task
echo "* Test change status of task *"
./todo.exe change "learn C" "Done" > /dev/null
out=$(./todo list)
if assert_contains "$out" "^learn C | Done$"; then
  pass
else
  fail
  all_passed=false
fi

# Test delete task by task name
echo "* Test delete task by task name *"
./todo.exe delete "learn C" > /dev/null
out=$(./todo list)
if assert_contains "$out" "learn C"; then
  fail
  all_passed=false
else
  pass
fi