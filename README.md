# 📝 Go ToDo App

A simple in-memory ToDo web application built with Go's standard library — no frameworks, no databases.

## Features

- Add new tasks via a web form
- View all current tasks
- Delete individual tasks
- Lightweight — uses only Go's `net/http` standard library

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher

## Getting Started

### 1. Clone or copy the source

```bash
git clone <your-repo-url>
cd todo-app
```

### 2. Run the app

```bash
go run main.go
```

### 3. Open in your browser

```
http://localhost:8080
```

## Project Structure

```
todo-app/
└── main.go       # All application logic
```

## API / Routes

| Route          | Method | Description                        |
|----------------|--------|------------------------------------|
| `/`            | GET    | Home page with task input form     |
| `/show-tasks`  | GET    | Lists all tasks with delete buttons|
| `/add-task`    | POST   | Adds a new task                    |
| `/delete-task` | POST   | Deletes a task by index            |

## Usage

1. Navigate to `http://localhost:8080`
2. Type a task into the input field and click **Add Task**
3. Click **View Tasks** to see your task list
4. Click **Delete** next to any task to remove it

## Limitations

- Tasks are stored **in-memory** — they are lost when the server restarts
- No user authentication or multi-user support
- Not recommended for production use without adding persistent storage (e.g. a database)

## Future Improvements

- Persist tasks using a file or database (e.g. SQLite, PostgreSQL)
- Add task editing / completion toggle
- Add a REST API with JSON responses
- Write unit tests for each handler

## Learnt from

Youtube and optimised it a little
