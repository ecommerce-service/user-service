package models

import (
	"database/sql"
	"time"
)

type Category struct {
	id        string
	name      string
	slug      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func NewCategoryModel() *Category {
	return &Category{}
}

func (model *Category) Id() string {
	return model.id
}

func (model *Category) SetId(id string) *Category {
	model.id = id

	return model
}

func (model *Category) Name() string {
	return model.name
}

func (model *Category) SetName(name string) *Category {
	model.name = name

	return model
}

func (model *Category) Slug() string {
	return model.slug
}

func (model *Category) SetSlug(slug string) *Category {
	model.slug = slug

	return model
}

func (model *Category) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Category) SetCreatedAt(createdAt time.Time) *Category {
	model.createdAt = createdAt

	return model
}

func (model *Category) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Category) SetUpdatedAt(updatedAt time.Time) *Category {
	model.updatedAt = updatedAt

	return model
}

func (model *Category) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Category) SetDeletedAt(deletedAt sql.NullTime) *Category {
	model.deletedAt = deletedAt

	return model
}

const (
	CategorySelectStatement = `SELECT id,name,slug FROM categories`
)

func (model *Category) ScanRows(rows *sql.Rows) (*Category, error) {
	err := rows.Scan(&model.id, &model.name)
	if err != nil {
		return model, err
	}

	return model, nil
}
