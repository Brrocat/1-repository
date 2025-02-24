package taskService

func GetTasks() ([]Task, error) {
	return GetAllTasks()
}

func AddTask(task *Task) error {
	return CreateTask(task)
}
