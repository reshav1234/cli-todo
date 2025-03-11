package main

func main() {
	todos := Todos{}

	storage := newStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlag := NewCmdFlags()
	cmdFlag.Execute(&todos)

	storage.Save(todos)
}
