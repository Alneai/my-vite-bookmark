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

type note struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}

type archive struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}

func createTable() {
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

	table = `
	    CREATE TABLE IF NOT EXISTS note (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        data VARCHAR(500) NULL
    );
	`
	_, err = db.Exec(table)
	checkErr(err)

	table = `
		CREATE TABLE IF NOT EXISTS archive (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		data VARCHAR(500) NULL
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

func getLink(c *gin.Context) {
	linkType := c.Param("type")
	fmt.Printf("[getLink] type: %s\n", linkType)

	links := queryAllLink(linkType)
	c.IndentedJSON(http.StatusOK, links)
}

func postLink(c *gin.Context) {
	var value link
	if err := c.BindJSON(&value); err != nil {
		return
	}

	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO link(type, title, link, description) VALUES(?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(value.Type, value.Title, value.Link, value.Description)
	checkErr(err)

	fmt.Printf("[postLink] link: %+v\n", value)
	c.IndentedJSON(http.StatusCreated, value)
}

func updateLink(c *gin.Context) {
	var value link
	if err := c.BindJSON(&value); err != nil {
		return
	}

	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("UPDATE link SET title=?, link=?, description=? WHERE id=?")
	checkErr(err)

	_, err = stmt.Exec(value.Title, value.Link, value.Description, value.Id)
	checkErr(err)

	fmt.Printf("[updateLink] link: %+v\n", value)
	c.IndentedJSON(http.StatusCreated, value)
}

func deleteLink(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("[deleteLink] type: %d\n", id)

	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("DELETE FROM link WHERE id=?")
	checkErr(err)

	_, err = stmt.Exec(id)
	checkErr(err)

	c.IndentedJSON(http.StatusOK, id)
}

func getNote(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM note")
	checkErr(err)

	var notes []note
	for rows.Next() {
		var result note
		err = rows.Scan(&result.Id, &result.Data)
		checkErr(err)
		notes = append(notes, result)
	}
	db.Close()

	c.IndentedJSON(http.StatusOK, notes)
}

func postNote(c *gin.Context) {
	var value note
	if err := c.BindJSON(&value); err != nil {
		return
	}

	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO note(data) VALUES(?)")
	checkErr(err)

	_, err = stmt.Exec(value.Data)
	checkErr(err)

	c.IndentedJSON(http.StatusCreated, value)
}

func updateNote(c *gin.Context) {
	var value note
	if err := c.BindJSON(&value); err != nil {
		return
	}

	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("UPDATE note SET data=? WHERE id=?")
	checkErr(err)

	_, err = stmt.Exec(value.Data, value.Id)
	checkErr(err)

	c.IndentedJSON(http.StatusCreated, value)
}

func deleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("DELETE FROM note WHERE id=?")
	checkErr(err)

	_, err = stmt.Exec(id)
	checkErr(err)
	c.IndentedJSON(http.StatusOK, id)
}

func getArchive(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM archive")
	checkErr(err)

	var archives []archive
	for rows.Next() {
		var result archive
		err = rows.Scan(&result.Id, &result.Data)
		checkErr(err)
		archives = append(archives, result)
	}
	db.Close()

	c.IndentedJSON(http.StatusOK, archives)
}

func postArchive(c *gin.Context) {
	var value archive
	if err := c.BindJSON(&value); err != nil {
		return
	}

	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO archive(data) VALUES(?)")
	checkErr(err)

	_, err = stmt.Exec(value.Data)
	checkErr(err)

	c.IndentedJSON(http.StatusCreated, value)
}

func updateArchive(c *gin.Context) {
	var value archive
	if err := c.BindJSON(&value); err != nil {
		return
	}

	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("UPDATE archive SET data=? WHERE id=?")
	checkErr(err)

	_, err = stmt.Exec(value.Data, value.Id)
	checkErr(err)

	c.IndentedJSON(http.StatusCreated, value)
}

func deleteArchive(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	db, err := sql.Open("sqlite3", "./link.db")
	checkErr(err)

	stmt, err := db.Prepare("DELETE FROM archive WHERE id=?")
	checkErr(err)

	_, err = stmt.Exec(id)
	checkErr(err)
	c.IndentedJSON(http.StatusOK, id)
}

// TODO: 面向对象重写
func main() {
	createTable()
	router := gin.Default()
	router.GET("/link/:type", getLink)
	router.POST("/link", postLink)
	router.POST("/link/update", updateLink)
	router.GET("/link/delete/:id", deleteLink)

	router.GET("/note", getNote)
	router.POST("/note", postNote)
	router.POST("/note/update", updateNote)
	router.GET("/note/delete/:id", deleteNote)

	router.GET("/archive", getArchive)
	router.POST("/archive", postArchive)
	router.POST("/archive/update", updateArchive)
	router.GET("/archive/delete/:id", deleteArchive)
	router.Run("localhost:5201")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
