package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"github.com/syedomair/revel_client/app/models"
	"github.com/syedomair/revel_client/app/routes"
	"github.com/syedomair/revel_client/app/services"
	"strconv"
)

type BookController struct {
	*revel.Controller
}

func (c BookController) PublishedBooksMicro() revel.Result {

	bookJson, _ := json.Marshal("")
	var resultMapMicro []map[string]interface{}
	resultMapMicro = services.BookService{}.GetPublicBooks(bookJson)

	user_id := c.Session["user_id"]
	user_name := c.Session["user_name"]
	server := c.Session["server"]
	return c.Render(user_id, user_name, server, resultMapMicro)
}

func (c BookController) PublishedBooks() revel.Result {

	resultMap := make(map[string]interface{})
	resultMap = services.BookService{}.GetBooks(c.Session)
	resultMap1 := make(map[string]interface{})
	resultMap1 = resultMap["data"].(map[string]interface{})
	data := resultMap1["list"]

	user_id := c.Session["user_id"]
	user_name := c.Session["user_name"]
	server := c.Session["server"]
	return c.Render(user_id, user_name, server, data)
}

func (c BookController) MyBooks() revel.Result {

	resultMap := make(map[string]interface{})
	resultMap = services.BookService{}.GetMyBooks(c.Session)
	resultMap1 := make(map[string]interface{})
	resultMap1 = resultMap["data"].(map[string]interface{})
	data := resultMap1["list"]

	user_id := c.Session["user_id"]
	user_name := c.Session["user_name"]
	server := c.Session["server"]
	return c.Render(user_id, user_name, server, data)
}

func (c BookController) AddBookForm() revel.Result {

	user_id := c.Session["user_id"]
	user_name := c.Session["user_name"]
	server := c.Session["server"]
	return c.Render(user_id, user_name, server)
}

func (c BookController) EditBookForm(book_id string) revel.Result {

	resultMap := make(map[string]interface{})
	resultMap = services.BookService{}.GetABook(book_id, c.Session)
	data := make(map[string]interface{})
	data = resultMap["data"].(map[string]interface{})

	user_id := c.Session["user_id"]
	user_name := c.Session["user_name"]
	server := c.Session["server"]
	return c.Render(user_id, user_name, server, data)
}

func (c BookController) EditBook(book models.Book) revel.Result {

	bookJson, err := json.Marshal(book)
	if err != nil {
		fmt.Println("error:", err)
	}
	bookId := strconv.Itoa(book.Id)
	_ = services.BookService{}.EditBook(bookId, bookJson, c.Session)

	return c.Redirect(routes.BookController.MyBooks())
}

func (c BookController) AddBook(book models.Book) revel.Result {

	user_id := c.Session["user_id"]
	book.UserId, _ = strconv.ParseInt(user_id, 10, 64)

	bookJson, err := json.Marshal(book)
	if err != nil {
		fmt.Println("error:", err)
	}
	_ = services.BookService{}.AddBook(bookJson, c.Session)

	return c.Redirect(routes.BookController.MyBooks())
}
