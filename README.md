# Task Tracker CLI

A simple command-line task tracker written in Go. It lets you add, update, delete, list, and mark tasks as in progress or done.

## Tags

- [Go](https://go.dev/)
- [CLI](https://en.wikipedia.org/wiki/Command-line_interface)
- [Task Manager](https://en.wikipedia.org/wiki/Task_management)
- [JSON](https://www.json.org/json-en.html)
- [Terminal App](https://en.wikipedia.org/wiki/Terminal_emulator)
-[Roadmap.sh project description] (https://roadmap.sh/projects/task-tracker)

## Features

- Add tasks
- Update task descriptions
- Delete tasks
- List tasks
- Filter tasks by status
- Mark tasks as in progress or done
- Persist tasks in a JSON file

## Requirements

- Go 1.20 or newer

## Installation

Clone the repository and run:

```bash
go run .
```

## Usage

### Add a task

```bash
go run . add "Buy groceries"
```

### Update a task

```bash
go run . update 1 "Buy groceries and milk"
```

### Delete a task

```bash
go run . delete 1
```

### List tasks

```bash
go run . list
```

### List tasks by status

```bash
go run . list todo
```

### Mark a task as in progress

```bash
go run . mark-in-progress 1
```

### Mark a task as done

```bash
go run . mark-done 1
```

## Storage

Tasks are stored in a JSON file named `tasks.json` in the project directory.

## Development

Run tests:

```bash
go test ./...
```
