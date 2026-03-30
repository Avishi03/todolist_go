package main

import (
	"fmt"
	"net/http"
	"strconv" // added for delete
)

var taskItems = []string{}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.HandleFunc("/add-task", addTask)
	http.HandleFunc("/delete-task", deleteTask)
	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	if len(taskItems) == 0 {
		fmt.Fprintln(writer, "<p>No tasks yet. <a href='/'>Add one!</a></p>")
		return
	}

	fmt.Fprintln(writer, "<h2>Your Tasks</h2><ul>")
	for i, task := range taskItems {
		// ← render a delete form next to each task
		fmt.Fprintf(writer, `<li>%s
			<form action="/delete-task" method="POST" style="display:inline">
				<input type="hidden" name="index" value="%d">
				<button type="submit">Delete</button>
			</form>
		</li>`, task, i)
	}
	fmt.Fprintln(writer, "</ul><a href='/'>Back</a>")
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

func deleteTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		indexStr := request.FormValue("index")
		index, err := strconv.Atoi(indexStr)
		if err == nil && index >= 0 && index < len(taskItems) {
			taskItems = append(taskItems[:index], taskItems[index+1:]...)
		}
	}
	http.Redirect(writer, request, "/show-tasks", http.StatusSeeOther)
}
