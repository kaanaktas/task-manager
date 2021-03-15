package manager

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository interface {
	newTask(name string) error
	delete(taskId string) (int64, error)
	taskDone(taskId string) (int64, error)
	findById(taskId string) (TaskData, error)
	findAll() ([]TaskData, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r repository) newTask(name string) error {
	taskId := uuid.New().String()
	_, err := r.db.Exec(`INSERT INTO task_table(task_id, task_name, task_done) VALUES ($1, $2, $3)`, taskId, name, false)
	if err != nil {
		return errors.WithMessage(err, "error in newTask")
	}

	return nil
}

func (r repository) delete(taskId string) (int64, error) {
	res, err := r.db.Exec(`DELETE FROM task_table WHERE task_id = $1`, taskId)
	if err != nil {
		return 0, errors.WithMessage(err, "error in delete")
	}

	if n, err := res.RowsAffected(); err != nil {
		return -1, errors.WithMessage(err, "error in delete")
	} else {
		return n, err
	}
}

func (r repository) taskDone(taskId string) (int64, error) {
	res, err := r.db.Exec(`UPDATE task_table SET task_done = true WHERE task_id = $1`, taskId)
	if err != nil {
		return 0, errors.WithMessage(err, "error in taskDone")
	}

	if n, err := res.RowsAffected(); err != nil {
		return 0, errors.WithMessage(err, "error in taskDone")
	} else {
		return n, err
	}
}

func (r repository) findById(taskId string) (TaskData, error) {
	var task TaskData
	err := r.db.Get(&task, "SELECT * FROM task_table WHERE task_id=$1", taskId)
	if err != nil {
		return TaskData{}, err
	}

	return task, nil
}

func (r repository) findAll() ([]TaskData, error) {
	var tasks []TaskData
	err := r.db.Select(&tasks, "SELECT * FROM task_table ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
