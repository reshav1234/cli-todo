# Command Line Todo 
A simple command-line To-Do application written in Go that allows you to add, edit, delete, and toggle tasks as completed.

# Installation 
Ensure you have Go installed on your system. Then, clone the repository and build the project:

## Clone this repository
```git
git clone "https://github.com/reshav1234/cli-todo.git"
```

# Build the application
```bash
cd cli-todo
go build -o main.go*
```

# Usuage
Run the todo with following commands

1. Add a task
   ```go
   ./main.go -add "Add post function
   "```
   
3. Edit a task
   ```go
   ./main.go -edit "<index">: "Test post function"
   ```
   
5. Toggle a task
   ```go
   ./main.go -toggle <index>
   ```
6. Delete a task
   ```go
   ./main.go -del <index>
   ```
8. 
   ```go
   ./main.go -del <index>
   ```
9. List all the tasks
   ```go
   ./main.go -list
   ```


