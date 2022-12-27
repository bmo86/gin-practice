package database

import (
	"context"
	"database/sql"
	"gin-practice/models"
)

type InstaceProstgres struct {
	db *sql.DB
}

func NewConnectionDatabase(url string) (*InstaceProstgres, error) {
	cn, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &InstaceProstgres{db: cn}, nil
}

func (i *InstaceProstgres) CreatedMe(ctx context.Context, me *models.Me) error {
	_, err := i.db.ExecContext(ctx, "INSERT INTO me(name, lastname, age) VALUES ($1, $2, $3)", me.Name, me.Lastname, me.Age)
	return err
}

func (i *InstaceProstgres) GetName(ctx context.Context, id int64) (*models.Me, error) {
	me, err := i.db.QueryContext(ctx, "SELECT name, lastname, age FROM me WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer me.Close()

	var meResponse = models.Me{}

	for me.Next() {
		if err = me.Scan(&meResponse.Name, &meResponse.Lastname, &meResponse.Age); err == nil {
			return &meResponse, nil
		}
	}

	if err = me.Err(); err != nil {
		return &meResponse, nil
	}

	return &meResponse, nil
}
