# Task Tracker - Roadmap.sh Project

## Table of Content
[Introduction](#introduction)
[Getting Started](#getting-started)
[Handle task operations](#handle-task-operations)
- [Add task](#add-task)
- [Update task](#update-task)
- [Delete task](#delete-task)
- [List tasks](#list-tasks)
    - [List all tasks that are done](#list-all-tasks-that-are-done)
    - [List all todo tasks only](#list-all-todo-tasks-only)
    - [List all tasks that are in progress](#list-all-tasks-that-are-in-progress)

## Introduction

Task Tracker is a beginner project governed by [Roadmap.sh](https://roadmap.sh/projects/task-tracker). This project's only purpose is for learning and practicing programming skills, getting into and becoming familiar with CLI tools.

## Getting Started


```bash
git clone https://github.com/surtr1st/task-tracker.git && cd task-tracker
go build ./cmd
```

## Handle task operations

### Add task

```bash
# ./cmd add "Shopping"
./cmd add <description>
```

### Update task


```bash
# ./cmd update 1 "Adding new feature"
./cmd update <id> <description>
```

### Delete task


```bash
# ./cmd rm 1
./cmd rm <id>
```

### List tasks


```bash
./cmd list
```

#### List all tasks that are done


```bash
./cmd list done
```

#### List all todo tasks only


```bash
./cmd list todo
```

#### List all tasks that are in progress


```bash
./cmd list ip
```

