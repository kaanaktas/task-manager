package main

import (
	"github.com/labstack/gommon/log"
	db "github/kaanaktas/task-manager/internal/store"
	"os"
	"testing"
)

const datasourceUrl = "./task.sqlite"

func init() {
	_ = os.Setenv("MIGRATE_VERSION", "2")
	_ = os.Setenv("MIGRATE_SCRIPT_URL", "file://../../scripts/sqlite")
	_ = os.Setenv("MIGRATE_DATABASE_URL", "sqlite3://"+datasourceUrl)
	_ = os.Setenv("DRIVER_NAME", "sqlite3")
	_ = os.Setenv("DATASOURCE_URL", datasourceUrl)
}

func Test_doMigrate(t *testing.T) {
	defer func() {
		err := os.Remove(datasourceUrl)
		if err != nil {
			log.Error(err)
		}
		log.Print("database removed")
	}()

	tests := []struct {
		name string
	}{
		{"migrate_db_success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doMigrate()
			dbx := db.LoadDBConnection()

			expected := new(int32)
			*expected = 1

			var got *int32
			_ = dbx.Get(&got, "Select count(task_name) from task_table")
			if *got != *expected {
				t.Errorf("doMigrate() = %v, want %v", *got, *expected)
			}

			_ = dbx.Close()
		})
	}
}
