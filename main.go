package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{} // initially empty

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.HandleFunc("/add-task", addTask)

	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	if len(taskItems) == 0 {
		fmt.Fprintln(writer, "No tasks yet. Add one!")
		return
	}

	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	html := `
	<h1>Welcome to ToDo App</h1>
	<form action="/add-task" method="POST">
		<input type="text" name="task" placeholder="Enter new task" required>
		<button type="submit">Add Task</button>
	</form>
	<br>
	<a href="/show-tasks">View Tasks</a>
	`
	fmt.Fprintln(writer, html)
}

func addTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		task := request.FormValue("task")
		if task != "" {
			taskItems = append(taskItems, task)
		}
	}
	http.Redirect(writer, request, "/show-tasks", http.StatusSeeOther)
}
