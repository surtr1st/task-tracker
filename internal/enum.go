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
	INVALID_INPUT           = "Your input ID is missing or not a number!"
	MISSING_FILE_PATH       = "File path is empty!"
	MISSING_CONTENT         = "Content is missing or invalid!"
	MISSING_TASK_ID         = "Missing specify task ID!"
	MISSING_REQUIRED_VALUES = "Missing `id` and new update value!"
	REMOVE_ALL_TASKS        = "Cleared all tasks!"
	ADDED_TASK              = "Added new task!"
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
	HELP         = "help"
	FLUSH        = "flush"
)

type FilterStatus string

const (
	FILTER_DONE        FilterStatus = "done"
	FILTER_IN_PROGRESS FilterStatus = "ip"
	FILTER_TODO        FilterStatus = "todo"
	FILTER_NONE        FilterStatus = "none"
)

type FilterUpdateProperty string

const (
	UPDATE_DESCRIPTION FilterUpdateProperty = "desc"
	UPDATE_STATUS      FilterUpdateProperty = "status"
)

var (
	INEXISTENCE_TASK = "Task with ID '%d' does not exist"
	UPDATED_TASK     = "Updated task with ID: %d"
	REMOVED_TASK     = "Removed task `%s` with ID: %d"
	SET_TASK_TO      = "Task with ID: %d - `%s` has set to '%s'"
)
