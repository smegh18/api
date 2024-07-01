package domain

import (
	m "book_seller/model"
	"context"
	"database/sql"
)

func (d *DomainDB) CreateUser(user m.User) error {
	sql := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := d.db.QueryRow(sql, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *DomainDB) GetUsers() (*sql.Rows, error) {
	rows, err := d.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows, nil
}

func (d *DomainDB) GetUser(id string, user m.User) error {
	err := d.db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (d *DomainDB) DeleteUser(id string) error {
	sql := `DELETE FROM users WHERE id=$1`
	_, err := d.db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DomainDB) UpdateUser(user m.User, id string) error {
	sql := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	_, err := d.db.Exec(sql, user.Name, user.Email, id)
	if err != nil {
		return err
	}
	return nil
}

type DomainDB struct {
	db SQLDatabase
}

func NewDomainDB(db SQLDatabase) *DomainDB {
	return &DomainDB{db: db}
}

// SQLDatabase ...
type SQLDatabase interface {
	PingContext(ctx context.Context) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type Service interface {
	CreateUser(user m.User) error
	GetUsers() (*sql.Rows, error)
	GetUser(id string, user m.User) error
	DeleteUser(id string) error
	UpdateUser(user m.User, id string) error
}
