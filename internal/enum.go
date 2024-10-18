package internal

type TaskStatus string

const (
	TODO        TaskStatus = "To Do"
	IN_PROGRESS TaskStatus = "In-Progress"
	DONE        TaskStatus = "Done"
)

const (
	WINDOWS     = "LOCALAPPDATA"
	LINUX       = "HOME"
	MAC         = "HOME"
	BASE_FOLDER = "task-tracker"
	FILEDATA    = "data.json"
)

const (
	INVALID_INPUT     = "Your input ID is missing or not a number."
	MISSING_FILE_PATH = "File path is empty!"
	MISSING_CONTENT   = "Content is missing or invalid!"
)

const (
	ADD          = "add"
	UPDATE       = "update"
	REMOVE       = "rm"
	LIST         = "list"
	_DONE        = "done"
	_IN_PROGRESS = "ip"
	_TODO        = "todo"
	INIT         = "init"
)

type FilterFlagList string

const (
	FLAG_DONE        FilterFlagList = "--done"
	FLAG_IN_PROGRESS FilterFlagList = "--ip"
	FLAG_TODO        FilterFlagList = "--todo"
	FLAG_NONE        FilterFlagList = "none"
)
