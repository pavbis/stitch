package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/pavbis/stitch/config"
	"github.com/pavbis/stitch/monitor"
	"log"
	"time"
)

type ProducerStreamRelation struct {
	UUID string
	Name string
}

func main() {
	a := time.Now()
	fmt.Printf("StartTime: %v\n", time.Now())

	mList := retrieveDataFromMariaDB()
	writeDataToPostgresSQL(mList)

	delta := time.Now().Sub(a)

	fmt.Printf("Execution time: %v secs\n", delta.Seconds())
	fmt.Println("Program finished successfully")
	fmt.Printf("StopTime: %v\n", time.Now())

	monitor.System()
}

func retrieveDataFromMariaDB() []ProducerStreamRelation {
	mariaDsn := config.NewMariaDsnFromEnv()
	mariaConnection, err := sql.Open("mysql", mariaDsn)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := mariaConnection.Prepare(`SELECT id,created FROM react_widgets`)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := statement.Query()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	mList := make([]ProducerStreamRelation, 0)
	producerStreamRelation := ProducerStreamRelation{}
	for rows.Next() {
		_ = rows.Scan(&producerStreamRelation.UUID, &producerStreamRelation.Name)
		mList = append(mList, producerStreamRelation)
	}

	_ = mariaConnection.Close()

	return mList
}

func writeDataToPostgresSQL(mList []ProducerStreamRelation) {
	postgresDsn := config.NewPostgresDsnFromEnv()
	psqConnection, _ := sql.Open("postgres", postgresDsn)
	psqConnection.SetMaxIdleConns(10)
	psqConnection.SetMaxOpenConns(10)
	psqConnection.SetConnMaxLifetime(0)

	transaction, err := psqConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, _ := transaction.Prepare(pq.CopyIn("producerStreamRelations", "producerId", "streamName"))
	for _, relation := range mList {
		_, err := stmt.Exec(relation.UUID, relation.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
