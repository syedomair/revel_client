package models

import (
)

type User struct {
    UserId            int     `json:"id"`
    Email             string  `json:"email"`
    FirstName         string  `json:"first_name"`
    LastName          string  `json:"last_name"`
    Password          string  `json:"password"`
}
