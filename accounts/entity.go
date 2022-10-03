package accounts

import (
	"time"
)

type Person struct {
	ID        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
}
