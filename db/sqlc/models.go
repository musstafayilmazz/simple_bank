// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID          int64         `json:"id"`
	Owner       string        `json:"owner"`
	Balance     int64         `json:"balance"`
	Currency    string        `json:"currency"`
	CreatedAt   time.Time     `json:"created_at"`
	CountryCode sql.NullInt32 `json:"country_code"`
}

type Entry struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int64          `json:"id"`
	FromAccountID int64          `json:"from_account_id"`
	ToAccountID   int64          `json:"to_account_id"`
	Code          int32          `json:"code"`
	Name          sql.NullString `json:"name"`
	ContinentName sql.NullString `json:"continent_name"`
	CreatedAt     time.Time      `json:"created_at"`
	Amount        int64          `json:"amount"`
}
