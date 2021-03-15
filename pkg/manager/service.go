package manager

type Service interface {
	newTask(name string) error
	taskDone(taskId string) (int64, error)
	delete(taskId string) (int64, error)
	findById(taskId string) (Task, error)
	findAll() ([]Task, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s service) newTask(name string) error {
	return s.repo.newTask(name)
}

func (s service) taskDone(taskId string) (int64, error) {
	return s.repo.taskDone(taskId)
}

func (s service) delete(taskId string) (int64, error) {
	return s.repo.delete(taskId)
}

func (s service) findById(taskId string) (Task, error) {
	if taskData, err := s.repo.findById(taskId); err != nil {
		return Task{}, err
	} else {
		return Task{
			TaskId: taskData.TaskId,
			Name:   taskData.Name,
			Done:   taskData.Done,
		}, nil
	}
}

func (s service) findAll() ([]Task, error) {
	if tasks, err := s.repo.findAll(); err != nil {
		return nil, err
	} else {
		taskList := make([]Task, len(tasks))
		for i, t := range tasks {
			task := Task{
				TaskId: t.TaskId,
				Name:   t.Name,
				Done:   t.Done,
			}
			taskList[i] = task
		}

		return taskList, nil
	}
}
