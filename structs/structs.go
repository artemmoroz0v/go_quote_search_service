package structs

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

var Db *sql.DB
var Data []Record
var Store []Dictionary

type Record struct {
	Title           string    `json:"title"`
	AuxiliaryText   []string  `json:"auxiliary_text"`
	Category        []string  `json:"category"`
	Wiki            string    `json:"wiki"`
	Language        string    `json:"language"`
	CreateTimestamp time.Time `json:"create_timestamp"`
	Timestamp       time.Time `json:"timestamp"`
}

type Dictionary struct {
	Category []string
	Name     string
}

type Amount struct {
	Number      int      `json:"number of quotes in a given category"`
	QuotesArray []string `json:"titles"`
}

type Change struct {
	OldTitle      string   `json:"old_title"`
	Title         string   `json:"new_title"`
	AuxiliaryText []string `json:"auxiliary_text"`
	Category      []string `json:"category"`
}

func AmountOfTitlesByCategory(Store []Dictionary, Name string) int {
	counter := 0
	for c := 0; c < len(Store); c++ {
		for d := 0; d < len(Store[c].Category); d++ {
			if Store[c].Category[d] == Name {
				counter++ //amount of titles with current category
			}
		}
	}
	return counter
}

func ListOfTitlesByCategory(Store []Dictionary, Name string) []string {
	list := []string{}
	for c := 0; c < len(Store); c++ {
		for d := 0; d < len(Store[c].Category); d++ {
			if Store[c].Category[d] == Name {
				list = append(list, Store[c].Name)
			}
		}
	}
	return list
}

func GetQuote(c *gin.Context) {
	name := c.Param("name")
	quote := Record{}
	request := Db.QueryRow("SELECT * FROM quotes WHERE title = $1", name)
	err := request.Scan(&quote.Title, pq.Array(&quote.AuxiliaryText), pq.Array(&quote.Category), &quote.Wiki, &quote.Language, &quote.CreateTimestamp, &quote.Timestamp)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "can not get quote by name"})
		return
	}
	c.IndentedJSON(http.StatusOK, quote)
}

func GetAllQuotesByCategory(c *gin.Context) {
	name := c.Param("name")
	amount := Amount{}
	amount.Number = AmountOfTitlesByCategory(Store, name)
	if amount.Number == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "there are not any quotes by given category"})
		return
	}
	amount.QuotesArray = ListOfTitlesByCategory(Store, name)
	c.IndentedJSON(http.StatusOK, amount)
}

func ChangeQuote(c *gin.Context) {
	change := Change{}
	err := c.BindJSON(&change)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "something went wrong with binding json"})
		return
	}
	_, err = Db.Exec("UPDATE quotes SET auxiliary_text = $1, category = $2, qtimestamp = $3, title = $4 WHERE (title = $5)", pq.Array(change.AuxiliaryText), pq.Array(change.Category), time.Now().Format("2006-01-02 15:04:05"), change.Title, change.OldTitle)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "something went wrong with updating title"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "data has been updated successfully"})
}
