package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"

	"UserServer/postgresql"
)

const (
	USER_ID_PREFIX = "user-"
)

type User struct {
	ID       int     `json:"-" db:"id"`
	Account  *string `json:"account,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (u *User) Upsert(ctx context.Context, db postgresql.DB) error {
	var (
		index  = 1
		values []string
		args   []interface{}
	)

	cols := []string{
		"account",
		"password",
	}
	setMap := map[string]interface{}{}
	jsonData, err := json.Marshal(u)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jsonData, &setMap); err != nil {
		return err
	}
	for _, col := range cols {
		values = append(values, fmt.Sprintf("$%d", index))
		args = append(args, setMap[col])
		index++
	}

	query := fmt.Sprintf(`
		INSERT INTO users (%s)
		VALUES (%s)
	`,
		strings.Join(cols, ","),
		strings.Join(values, ","),
	)
	_, err = db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Get(ctx context.Context, db postgresql.DB) (*User, error) {
	var (
		index      = 1
		conditions = []string{}
		args       = []interface{}{}
	)

	cols := []string{
		"account",
		"password",
	}

	if u.Account != nil {
		conditions = append(conditions, fmt.Sprintf("account = $%d", index))
		args = append(args, u.Account)
		index++ /*
			conditions = append(conditions, fmt.Sprintf("password = $%d", index))
			args = append(args, u.Password)
			index++*/
	}
	query := fmt.Sprintf(`
		SELECT %s
		FROM users
		WHERE %s`,
		strings.Join(cols, ","),
		strings.Join(conditions, " AND "),
	)
	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	var us []*User
	err = pgxscan.ScanAll(&us, rows)
	if err != nil {
		return nil, err
	}

	if len(us) == 1 {
		if *us[0].Password == *u.Password {
			return us[0], nil
		} else {
			return nil, ErrorPassword
		}
	} else if len(us) > 1 {
		return nil, ErrorMultiplRows
	} else {
		return nil, pgx.ErrNoRows
	}
}
