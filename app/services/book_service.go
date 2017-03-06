package services

import (
    "github.com/revel/revel"
)

type BookService struct {
    CommonService
}

func (c BookService) GetBooks(session revel.Session) map[string]interface{} {
    return c.CallGetBackend("GET", "public/books", session) 
}

func (c BookService) GetMyBooks(session revel.Session) map[string]interface{} {
    user_id := session["user_id"] 
    return c.CallGetBackend("GET", "my-books/"+user_id, session) 
}

func (c BookService) AddBook(jsonStr []byte, session revel.Session) map[string]interface{} {
    return c.CallBackend("POST", "book", jsonStr, true, session) 
}

func (c BookService) EditBook(book_id string, jsonStr []byte, session revel.Session) map[string]interface{} {
    return c.CallBackend("PATCH", "book/"+book_id, jsonStr, true, session) 
}

func (c BookService) GetABook(book_id string, session revel.Session) map[string]interface{} {
    return c.CallGetBackend("GET", "book/"+book_id, session) 
}

func (c BookService) GetPublicBooks(jsonStr []byte) []map[string]interface{} {
    return c.CallMircoBackend("POST", "public/books", jsonStr) 
}


