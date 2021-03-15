package api

import (
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"sync"
)

var mute = &sync.Mutex{}

func RunSql(dbx *sqlx.DB, sqlFile string) {
	mute.Lock()

	log.Printf("executing sql statement: %v", sqlFile)

	data, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		log.Print("error in RunSql(). couldn't read the file")
		panic(err)
	}
	_, err = dbx.Exec(string(data))
	if err != nil {
		log.Print("error in RunSql(). couldn't exec the sql")
		panic(err)
	}
	log.Printf("sql statement has been executed: %v", sqlFile)

	mute.Unlock()
}
