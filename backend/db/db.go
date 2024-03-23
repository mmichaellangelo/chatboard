package db

import (
	"chatboard/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBPool struct {
	*pgxpool.Pool
}

func Connect(ctx context.Context) (*DBPool, error) {
	pool, err := pgxpool.New(ctx, "postgres://mike:password@db:5432/db")
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	return &DBPool{Pool: pool}, nil
}

func (pool *DBPool) Disconnect() {
	pool.Close()
}

func (pool *DBPool) GetAccounts(ctx context.Context) ([]models.Account, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("context canceled while fetching accounts")
	default:
	}
	rows, err := pool.Query(ctx, "SELECT id, username, email, password, date_created FROM account")
	if err != nil {
		return nil, fmt.Errorf("error getting accounts: %v", err)
	}
	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.Username, &account.Email, &account.Password, &account.DateCreated)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		accounts = append(accounts, account)
	}
	return accounts, nil

}
