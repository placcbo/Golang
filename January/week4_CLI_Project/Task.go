package week4cliproject

// Task represents a single todo item
type Task struct {
	ID    int
	Title string
	Done  bool
}

// Manager manages a collection of tasks
type Manager struct {
	tasks  []Task
	nextID int
}

// Add creates a new task and stores it
func (m *Manager) Add(title string) Task {
	task := Task{
		ID:    m.nextID,
		Title: title,
		Done:  false,
	}

	m.tasks = append(m.tasks, task)
	m.nextID++

	return task
}

// List returns all tasks
func (m *Manager) List() []Task {
	return m.tasks
}

// Complete marks a task as done by ID
func (m *Manager) Complete(id int) bool {
	for i, t := range m.tasks {
		if t.ID == id {
			m.tasks[i].Done = true
			return true
		}
	}
	return false

}
