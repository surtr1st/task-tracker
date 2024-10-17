package internal

type taskTracker struct{}

func UseTaskTracker() taskTracker {
	return taskTracker{}
}

func (tracker taskTracker) Add(description string) {}

func (tracker taskTracker) Update() {}

func (tracker taskTracker) List() {}
