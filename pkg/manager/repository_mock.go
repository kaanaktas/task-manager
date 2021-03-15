package manager

import (
	"database/sql"
	"strconv"
)

type mockRepository struct {
}

var tasks []TaskData

func (m mockRepository) newTask(name string) error {
	id := len(tasks) + 1
	task := TaskData{
		Id:     id,
		TaskId: strconv.Itoa(id),
		Name:   name,
		Done:   false,
	}
	tasks = append(tasks, task)

	return nil
}

func (m mockRepository) delete(taskId string) (int64, error) {
	for i, task := range tasks {
		if task.TaskId == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	return 1, nil
}

func (m mockRepository) taskDone(taskId string) (int64, error) {
	for i := range tasks {
		task := &tasks[i]
		if task.TaskId == taskId {
			task.Done = true
		}
	}

	return 1, nil
}

func (m mockRepository) findById(taskId string) (TaskData, error) {
	for _, task := range tasks {
		if task.TaskId == taskId {
			return task, nil
		}
	}

	return TaskData{}, sql.ErrNoRows
}

func (m mockRepository) findAll() ([]TaskData, error) {
	return tasks, nil
}

func MockRepository() Repository {
	tasks = tasks[:0]
	return &mockRepository{}
}
