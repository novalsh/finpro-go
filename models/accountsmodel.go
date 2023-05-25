package models

import (
	"database/sql"
	"fmt"

	"github.com/novalsh/finpro-go/config"
	"github.com/novalsh/finpro-go/entities"
)

type AccountsModel struct {
	conn *sql.DB
}

func NewAccountsModel() *AccountsModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &AccountsModel{
		conn: conn,
	}
}

func (p *AccountsModel) FindAll() ([]entities.Accounts, error) {
	rows, err := p.conn.Query("SELECT * FROM accounts")
	if err != nil {
		return []entities.Accounts{}, err
	}

	defer rows.Close()

	var accounts []entities.Accounts

	for rows.Next() {
		var account entities.Accounts
		rows.Scan(&account.Id, &account.Username, &account.Password, &account.Email, &account.Phone, &account.Role, &account.Token, &account.Created_at, &account.Updated_at)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (p *AccountsModel) Create(accounts entities.Accounts) bool {
	result, err := p.conn.Exec("INSERT INTO accounts (username, password, email, phone, role, token) VALUES (?, ?, ?, ?, ?, ?)", accounts.Username, accounts.Password, accounts.Email, accounts.Phone, accounts.Role, accounts.Token)
	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *AccountsModel) Update(accounts entities.Accounts) error {
	_, err := p.conn.Exec("UPDATE accounts SET username = ?, password = ?, email = ?, phone = ?, role = ?, token = ? WHERE id = ?", accounts.Username, accounts.Password, accounts.Email, accounts.Phone, accounts.Role, accounts.Token, accounts.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *AccountsModel) Delete(id int64) error {
	_, err := p.conn.Exec("DELETE FROM accounts WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
