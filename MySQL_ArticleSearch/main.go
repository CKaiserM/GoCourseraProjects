package main

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
)

type Article struct {
	Id       string
	Title    string
	Url      string
	Keywords string
}

func RenderMainPage(c *fiber.Ctx) error {
	return c.Render("main_page", fiber.Map{})
}

func ProcessAddForm(c *fiber.Ctx) error {

	fmt.Println("ProcessAddForm(): ---------------------------------")

	// Task 5 TODO - get a database connection
	db, err := connectDB()
	if err != nil {
		return c.Render("add_result", fiber.Map{"message": "Failed to connect to database"})
	}

	// fetch the user input from the form
	title := c.FormValue("title")
	url := c.FormValue("url")
	keywords := c.FormValue("keywords")

	// Task 5 TODO - trim white spaces, sort keywords
	transformedKeywords := transformKeywordStirng(keywords)

	keywordsParam := strings.Join(transformedKeywords[:], " ")
	fmt.Printf("ProcessAddForm(): keywordsParam = %v\n", keywordsParam)
	// generate a uuid for article ID
	id := uuid.NewString()

	// Task 5 TODO - execute the insert using the 4 values above
	// excute the select
	_, err = db.Exec("INSERT INTO web_article (id, title, url, keywords) VALUES (?, ?, ?, ?)", id, title, url, keywordsParam)
	if err != nil {
		return c.Render("add_result", fiber.Map{"message": "Failed to add new article " + title})
	}

	// make a Article to pass to the result page
	article := Article{
		Id:       id,
		Title:    title,
		Url:      url,
		Keywords: keywords,
	}

	// Task 5 TODO - uncomment and pass the Article to the result page
	return c.Render("add_result", fiber.Map{"message": "The following article has been added to the database:", "article": article})
}

func ProcessDeleteForm(c *fiber.Ctx) error {

	fmt.Println("ProcessDeleteForm(): ---------------------------------")

	// Task 6 TODO - get a database connection
	db, err := connectDB()
	if err != nil {
		return c.Render("delete_result", fiber.Map{"message": "Failed to connect to database"})
	}

	// fetch the user input from the form
	id := c.FormValue("id")

	// try to select the article to make sure that it exists
	// also get the rest of the article information to display later

	// Task 6 TODO - uncomment the following line and use SQL to find their values
	var title, url, keywords string
	matchingRow := db.QueryRow("SELECT title, url, keywords FROM web_article WHERE id=?", id)
	// Task 6 TODO - use the following line if the select failed

	err = matchingRow.Scan(&title, &url, &keywords)
	if err != nil {
		return c.Render("delete_result", fiber.Map{"message": "No article found with ID: " + id})
	}
	// Task 6 TODO - uncomment the following lines when implementing the feature

	// make a Article to pass to the result page
	article := Article{
		Id:       id,
		Title:    title,
		Url:      url,
		Keywords: keywords,
	}

	// Task 6 TODO - execute the delete
	_, err = db.Exec("DELETE FROM web_article WHERE id = ?", id)
	if err != nil {
		return c.Render("delete_result", fiber.Map{"message": "Failed to delete this article:", "article": article})
	}

	// pass the Article to the result page

	// Task 6 TODO - uncomment the following line
	return c.Render("delete_result", fiber.Map{"message": "The following article has been deleted from the database:", "article": article})
}

func ProcessSearchForm(c *fiber.Ctx) error {

	fmt.Println("ProcessSearchForm(): ---------------------------------")

	// fetch the user input from the form
	keywordStr := c.FormValue("keywords")

	// Task 2 TODO
	// get a database connection

	// Task 2 TODO - uncomment the next line and use it in case of error
	db, err := connectDB()
	if err != nil {
		return c.Render("search_result", fiber.Map{"message": "Failed to connect to database"})
	}

	// Task 3 TODO
	// trim white spaces, sort keywords
	keywords := transformKeywordStirng(keywordStr)
	// build the string to be used with SQL LIKE, a string that is
	// keywords separated by % and preceded by and end with a %
	likeParam := "%"

	// Task 3 TODO - make likeParam looks like this "%keyword1%keyword2%" etc

	for _, keyword := range keywords {
		likeParam += keyword + "%"
	}

	fmt.Printf("ProcessSearchForm(): likeParam = %v\n", likeParam)

	// Task 4 TODO - execute the SELECT SQL using likeParam
	// excute the select
	rows, err := db.Query("SELECT * FROM web_article WHERE keywords LIKE ?", likeParam)
	if err != nil {
		return c.Render("search_result", fiber.Map{"message": "db.Query failed"})
	}

	defer rows.Close()
	defer fmt.Println("ProcessSearchForm(): Rows closed (deffered)")

	// iterate through the rows to parse the matching articles
	var articles []Article

	// Task 4 TODO - put the result of the SELECT SQL into the slice articles
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.Id, &article.Title, &article.Url, &article.Keywords); err != nil {
			continue
		}
		articles = append(articles, article)
	}

	// Task 4 TODO - uncomment the next block
	// no match, simply pass a message, no articles
	if len(articles) == 0 {
		return c.Render("search_result", fiber.Map{"message": "No matches found"})
	}

	// Task 4 TODO - uncomment the next 2 lines and pass the articles to the result page
	// pass the matching articles slices to the result page
	message := fmt.Sprintf("Using keywords: %v Aert", strings.Trim(keywordStr, " "))
	return c.Render("search_result", fiber.Map{"message": message, "articles": articles})

}

func main() {

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	// Serve static files (HTML templates and stylesheets).
	app.Static("/", "./static")

	// Define routes.
	app.Get("/", RenderMainPage)
	app.Post("/search", ProcessSearchForm)
	app.Post("/add", ProcessAddForm)
	app.Post("/delete", ProcessDeleteForm)

	// Start the Fiber app on port 3000.
	app.Listen(":3000")
}

func connectDB() (*sql.DB, error) {
	var db *sql.DB
	cfg := mysql.Config{
		User:   "ituser",
		Passwd: "ituser",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "it_support",
	}
	// connect to the database
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Printf("connectDB(): Failed to connect to database %v\n", cfg.DBName)
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	fmt.Printf("connectDB(): Successfully connected to database %v\n", cfg.DBName)
	return db, nil
}
func transformKeywordStirng(keywordStr string) []string {
	keywordStr = strings.Trim(keywordStr, " ")
	keywordsAll := strings.Split(keywordStr, " ")
	fmt.Printf("transformKeywordStirng(): keywordsAll = %v\n", keywordsAll)

	var keywordsNonEmpty []string
	for i, w := range keywordsAll {
		if len(w) > 0 {
			keywordsNonEmpty = append(keywordsNonEmpty, strings.ToLower(keywordsAll[i]))
			// keywords = append(keywords, strings.ToLower(w))
		}

	}
	fmt.Printf("transformKeywordStirng(): keywordsNonEmpty = %v\n", keywordsNonEmpty)

	sort.Strings(keywordsNonEmpty)
	fmt.Printf("transformKeywordStirng(): keywordsNonEmpty sorted = %v\n", keywordsNonEmpty)

	return keywordsNonEmpty
}
