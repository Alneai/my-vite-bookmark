package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type link struct {
	Id          int    `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func createTableLink() {
	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	table := `
    CREATE TABLE IF NOT EXISTS link (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		type string VARCHAR(64) NOT NULL,
		title string VARCHAR(64) NOT NULL,
        link VARCHAR(128) NOT NULL,
        description VARCHAR(128) NULL
    );
    `
	_, err = db.Exec(table)
	checkErr(err)

	db.Close()
}

func queryAllLink(linkType string) []link {
	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM link WHERE type = ?", linkType)
	checkErr(err)

	var links []link
	for rows.Next() {
		var result link
		err = rows.Scan(&result.Id, &result.Type, &result.Title, &result.Link, &result.Description)
		checkErr(err)
		links = append(links, result)
	}
	db.Close()

	return links
}

func addLink(value link) {
	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO link(type, title, link, description) VALUES(?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(value.Type, value.Title, value.Link, value.Description)
	checkErr(err)
}

func updateLinkDb(value link) {
	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("UPDATE link SET title=?, link=?, description=? where id=?")
	checkErr(err)

	_, err = stmt.Exec(value.Title, value.Link, value.Description, value.Id)
	checkErr(err)
}

func deleteLinkDb(id int) {
	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("DELETE FROM link WHERE id=?")
	checkErr(err)

	_, err = stmt.Exec(id)
	checkErr(err)
}

func getLink(c *gin.Context) {
	linkType := c.Param("type")
	fmt.Printf("[getLink] type: %s\n", linkType)

	links := queryAllLink(linkType)
	// fmt.Printf("[Result] links: %+v\n", links)
	c.IndentedJSON(http.StatusOK, links)
}

func postLink(c *gin.Context) {
	var value link
	if err := c.BindJSON(&value); err != nil {
		return
	}

	addLink(value)
	fmt.Printf("[postLink] link: %+v\n", value)
	c.IndentedJSON(http.StatusCreated, value)
}

func updateLink(c *gin.Context) {
	var value link
	if err := c.BindJSON(&value); err != nil {
		return
	}

	updateLinkDb(value)
	fmt.Printf("[updateLink] link: %+v\n", value)
	c.IndentedJSON(http.StatusCreated, value)
}

func deleteLink(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("[deleteLink] type: %d\n", id)

	deleteLinkDb(id)
	c.IndentedJSON(http.StatusOK, id)
}

func main() {
	createTableLink()
	router := gin.Default()
	router.GET("/link/:type", getLink)
	router.POST("/link/add_link", postLink)
	router.POST("/link/update_link", updateLink)
	router.GET("/link/delete_link/:id", deleteLink)
	router.Run("localhost:5201")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
