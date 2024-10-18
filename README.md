# Task Tracker - Roadmap.sh Project

## Table of Content
- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Handling task operations](#handling-task-operations)
    - [Addding a task](#adding-a-task)
    - [Updating a task](#updating-atask)
    - [Deleting a task](#deleting-a-task)
    - [Listing tasks](#listing-tasks)
    - [Flushing tasks](#flushing-tasks)

## Introduction

Task Tracker is a beginner project governed by [Roadmap.sh](https://roadmap.sh/projects/task-tracker). This project's only purpose is for learning and practicing programming skills, getting into and becoming familiar with CLI tools.

## Getting Started

Cloning and building the project:

```bash
git clone https://github.com/surtr1st/task-tracker.git && cd task-tracker
go build ./cmd
```

Using `init` to initialize the data file in your machine. The file data will usually auto-create whenever you're using any commands. But if any unexpected errors occur, you may use this just in case.

```bash
./cmd init
```
This tool does not have autocomplete, so you may use the `help` subcommand to show an available list of subcommands.

```bash
./cmd help
```

## Handling task operations

Commands for handling task operations should be simple and intuitive.

### Adding a task

```bash
# ./cmd add "Shopping"
./cmd add <description>
```

### Updating a task

```bash
# ./cmd update 1 "Adding new feature"
./cmd update <id> <description>
```

### Deleting a task

```bash
# ./cmd rm 1
./cmd rm <id>
```

### Listing tasks

This subcommand also has the `help` subcommand in order to list all subcommands of the `list` subcommand.

```bash
./cmd list help
```

Listing all tasks without any status filter

```bash
./cmd list
```

Listing all tasks that are done

```bash
./cmd list done
```

Listing all todo tasks only

```bash
./cmd list todo
```

Listing all tasks that are in progress

```bash
./cmd list ip
```

### Flushing tasks

This subcommand is generally used for testing, but you may use it as you please.

```bash
./cmd flush
```