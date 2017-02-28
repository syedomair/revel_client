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
