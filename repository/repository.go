package repository

import (
	"context"
	"gin-practice/models"
)

type Repository interface {
	GetName(ctx context.Context, id int64) (*models.Me, error)
	CreatedMe(ctx context.Context, me *models.Me) error
}

var repo Repository

func SetRepository(r Repository) {
	repo = r
}

func GetName(ctx context.Context, id int64) (*models.Me, error) {
	return repo.GetName(ctx, id)
}

func CreatedMe(ctx context.Context, me *models.Me) error {
	return repo.CreatedMe(ctx, me)
}
