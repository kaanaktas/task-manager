package manager

type TaskData struct {
	Id     int
	TaskId string `db:"task_id"`
	Name   string `db:"task_name"`
	Done   bool   `db:"task_done"`
}

type Task struct {
	TaskId string `json:"id"`
	Name   string `json:"name"`
	Done   bool   `json:"done"`
}
