package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	s "task/structs"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	connection_line = "user=postgres password=postgres dbname=Quotes sslmode=disable"
)

func main() {
	//Parse JSON
	bytesRead, _ := ioutil.ReadFile("ruwikiquote-20221212-cirrussearch-general.json")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	i := 0
	for _, line := range lines {
		var r s.Record
		err := json.Unmarshal([]byte(line), &r)
		if reflect.DeepEqual(r, s.Record{}) {
			continue
		}
		if err != nil {
			continue
		}
		temp := s.Dictionary{}
		temp.Category = r.Category
		temp.Name = r.Title
		s.Data = append(s.Data, r)
		s.Store = append(s.Store, temp)
		i++
	}
	//fmt.Print("Total amount ", i) 12334

	//Initialization of Database, adding all records
	var err error
	s.Db, err = sql.Open("postgres", connection_line)
	if err != nil {
		log.Fatal(errors.New("error with connecting"))
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(errors.New("error with closing database"))
		}
	}(s.Db)
	err = s.Db.Ping()
	if err != nil {
		log.Fatal(errors.New("error with pinging database"))
	}
	_, err = s.Db.Exec("CREATE TABLE IF NOT EXISTS quotes (title VARCHAR(255), auxiliary_text text[], category text[], wiki VARCHAR(50) NOT NULL, qlanguage VARCHAR(20) NOT NULL, create_timestamp TIMESTAMP NOT NULL, qtimestamp TIMESTAMP NOT NULL)")
	if err != nil {
		log.Fatal(errors.New("error with creating database"))
	}
	for j := 0; j < i; j++ {
		_, err = s.Db.Exec("INSERT INTO quotes (title, auxiliary_text, category, wiki, qlanguage, create_timestamp, qtimestamp) VALUES ($1, $2, $3, $4, $5, $6, $7)", s.Data[j].Title, pq.Array(s.Data[j].AuxiliaryText), pq.Array(s.Data[j].Category), s.Data[j].Wiki, s.Data[j].Language, s.Data[j].CreateTimestamp, s.Data[j].Timestamp)
		if err != nil {
			fmt.Print(err)
		}
	}
	router := gin.Default()
	router.GET("/wiki/:name", s.GetQuote)                      //получение статьи по названию WORKS
	router.GET("/wiki/amount/:name", s.GetAllQuotesByCategory) //получение количества названий по категории
	router.PUT("/wiki/change", s.ChangeQuote)                  //редактирование статьи

	//running server
	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal(errors.New("error with running router"))
	}
}
