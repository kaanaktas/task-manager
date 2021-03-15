package manager

import (
	"github.com/jmoiron/sqlx"
	"github/kaanaktas/task-manager/internal/store"
	"reflect"
	"testing"
)

func Test_repository_findAll(t *testing.T) {
	task1 := TaskData{
		Id:     1,
		TaskId: "1",
		Name:   "task_1",
		Done:   false,
	}
	type fields struct {
		db *sqlx.DB
	}
	tests := []struct {
		name         string
		fields       fields
		expectedLen  int
		expectedData TaskData
		wantErr      bool
	}{
		{
			"findAll()",
			fields{db: store.LoadDBConnection()},
			5,
			task1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.findAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("findAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.expectedLen {
				t.Errorf("findAll() got = %v, want %v", len(got), tt.expectedLen)
			}

			if !reflect.DeepEqual(got[0], tt.expectedData) {
				t.Errorf("findAll() got = %v, want %v", got[0], tt.expectedData)
			}
		})
	}
}

func Test_repository_delete(t *testing.T) {
	type fields struct {
		db *sqlx.DB
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
			"delete_first_task()",
			fields{db: store.LoadDBConnection()},
			args{taskId: "1"},
			1,
			false,
		},
		{
			"delete_missing_task()",
			fields{db: store.LoadDBConnection()},
			args{taskId: "-1"},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.delete(tt.args.taskId)
			if (err != nil) != tt.wantErr {
				t.Errorf("delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_newTask(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"newTask()",
			fields{db: store.LoadDBConnection()},
			args{name: "new_task_1"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.newTask(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("newTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_taskDone(t *testing.T) {
	taskId := "123456"

	type fields struct {
		db *sqlx.DB
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
			"mark_task_done()",
			fields{db: store.LoadDBConnection()},
			args{taskId: taskId},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}

			got, err := r.taskDone(tt.args.taskId)
			if (err != nil) != tt.wantErr {
				t.Errorf("taskDone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("taskDone() got = %v, want %v", got, tt.want)
			}
			taskData, err := r.findById(taskId)
			if err != nil {
				t.Errorf("taskDone() error = %v", err)
				return
			}
			if taskData.Done != true {
				t.Errorf("taskDone() got = %v, want %v", taskData.Done, true)
			}
		})
	}
}

func Test_repository_findById(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		taskId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"findById()",
			fields{db: store.LoadDBConnection()},
			args{"123456"},
			"task_123456",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.findById(tt.args.taskId)
			if (err != nil) != tt.wantErr {
				t.Errorf("findById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Name != tt.want {
				t.Errorf("findById() got = %v, want %v", got.Name, tt.want)
			}
		})
	}
}
