package models

import (
)

type Book struct {
    Id                int     `json:"id"`
    UserId            int64   `json:"user_id"`
    Name              string  `json:"name"`
    Description       string  `json:"description"`
    Publish           bool    `json:"publish"`
}

type BookResponse struct {
    Id              int64  `json:"id"`
    UserId          int64  `json:"user_id"`
    FirstName       string `json:"first_name"`
    LastName        string `json:"last_name"`
    Name            string  `json:"book_name" `
    Description     string  `json:"description" `
    Publish         bool  `json:"publish" `
}
