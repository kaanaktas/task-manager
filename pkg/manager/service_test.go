package manager

import (
	"reflect"
	"testing"
)

func Test_service_newTask_empty_list(t *testing.T) {
	taskName := "buy some milk"
	type fields struct {
		repo Repository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedLen int
		wantErr     bool
	}{
		{
			"new_task",
			fields{repo: MockRepository()},
			args{name: taskName},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			if err := s.newTask(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("newTask() error = %v, wantErr %v", err, tt.wantErr)
			}

			tasks, _ := s.findAll()
			if len(tasks) != tt.expectedLen {
				t.Errorf("newTask() error = %v, wantErr %v", len(tasks), tt.expectedLen)
			}
			if tasks[0].Name != taskName {
				t.Errorf("newTask() error = %v, wantErr %v", tasks[0].Name, taskName)
			}
		})
	}
}

func Test_service_newTask_with_item(t *testing.T) {
	taskPreviousName := "buy some milk"
	taskName := "enjoy the assigment"
	type fields struct {
		repo Repository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedLen int
		wantErr     bool
	}{
		{
			"new_task_with_item",
			fields{repo: MockRepository()},
			args{name: taskName},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			_ = s.newTask(taskPreviousName)

			if err := s.newTask(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("newTask() error = %v, wantErr %v", err, tt.wantErr)
			}

			tasks, _ := s.findAll()
			if len(tasks) != tt.expectedLen {
				t.Errorf("newTask() error = %v, wantErr %v", len(tasks), tt.expectedLen)
			}
			if tasks[0].Name != taskPreviousName {
				t.Errorf("newTask() error = %v, wantErr %v", tasks[0].Name, taskPreviousName)
			}
			if tasks[1].Name != taskName {
				t.Errorf("newTask() error = %v, wantErr %v", tasks[1].Name, taskName)
			}
		})
	}
}

func Test_service_delete_empty_list(t *testing.T) {
	taskName := "rest for a while"
	type fields struct {
		repo Repository
	}
	type args struct {
		taskId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			"delete_item",
			fields{repo: MockRepository()},
			args{taskId: "1"},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			_ = s.newTask(taskName)
			tasks, _ := s.findAll()
			if len(tasks) != 1 {
				t.Errorf("delete() error = %v, wantErr %v", len(tasks), 1)
			}
			got, err := s.delete(tt.args.taskId)
			if (err != nil) != tt.wantErr {
				t.Errorf("delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("delete() got = %v, want %v", got, tt.want)
			}
			tasks, _ = s.findAll()
			if len(tasks) != 0 {
				t.Errorf("delete() error = %v, wantErr %v", len(tasks), 0)
			}
		})
	}
}

func Test_service_delete_with_item_in_list(t *testing.T) {
	taskPreviousName := "rest for a while"
	taskName := "drink water"
	type fields struct {
		repo Repository
	}
	type args struct {
		taskId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			"delete_with_item_in_list",
			fields{repo: MockRepository()},
			args{taskId: "1"},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			_ = s.newTask(taskPreviousName)
			_ = s.newTask(taskName)
			tasks, _ := s.findAll()
			if len(tasks) != 2 {
				t.Errorf("delete() error = %v, wantErr %v", len(tasks), 2)
			}
			got, err := s.delete(tt.args.taskId)
			if (err != nil) != tt.wantErr {
				t.Errorf("delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("delete() got = %v, want %v", got, tt.want)
			}
			tasks, _ = s.findAll()
			if len(tasks) != 1 {
				t.Errorf("delete() error = %v, wantErr %v", len(tasks), 0)
			}
			if tasks[0].Name != taskName {
				t.Errorf("delete() error = %v, wantErr %v", tasks[0].Name, taskName)
			}
		})
	}
}

func Test_service_taskDone(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		taskId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			"task_done",
			fields{MockRepository()},
			args{taskId: "1"},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			_ = s.newTask("new_task_1")
			got, err := s.taskDone(tt.args.taskId)
			if (err != nil) != tt.wantErr {
				t.Errorf("taskDone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("taskDone() got = %v, want %v", got, tt.want)
			}
			task, _ := s.findById(tt.args.taskId)
			if task.Done != true {
				t.Errorf("taskDone() got = %v, want %v", task.Done, true)
			}
		})
	}
}

func Test_service_findById(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		taskId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Task
		wantErr bool
	}{
		{
			"find_by_id",
			fields{MockRepository()},
			args{taskId: "1"},
			Task{
				TaskId: "1",
				Name:   "new_task",
				Done:   false,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			_ = s.newTask("new_task")
			got, err := s.findById(tt.args.taskId)
			if (err != nil) != tt.wantErr {
				t.Errorf("findById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_findAll(t *testing.T) {
	type fields struct {
		repo Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Task
		wantErr bool
	}{
		{
			"find_all",
			fields{MockRepository()},
			[]Task{{
				TaskId: "1",
				Name:   "new_task",
				Done:   false,
			}, {
				TaskId: "2",
				Name:   "new_task_2",
				Done:   false,
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			_ = s.newTask("new_task")
			_ = s.newTask("new_task_2")
			got, err := s.findAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("findAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
