package dao

import (
	sq "github.com/Masterminds/squirrel"
	null "gopkg.in/guregu/null.v3"

	"github.com/suzujun/yatteiki-cloud/goapp/model"
)

type (
	// TodoDao ...
	TodoDao interface {
		FindByID(id int) (*model.Todo, error)
		FindAll(limit uint64, marker *int) ([]model.Todo, *int, error)
		Insert(title string) (int, error)
		Update(id int, title null.String, completed null.Bool) error
		Delete(id int) error
	}
	todoDaoImpl struct {
		baseDao
	}
)

// NewTodoDao ...
func NewTodoDao() TodoDao {
	return &todoDaoImpl{
		baseDao: newDao(model.Todo{}),
	}
}

// FindByID ...
func (dao todoDaoImpl) FindByID(id int) (*model.Todo, error) {
	builder := dao.newSelectBuilder().
		Where(sq.Eq{"id": id})
	var m model.Todo
	if err := dao.findOneByBuilder(&builder, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// FindAll ...
func (dao todoDaoImpl) FindAll(limit uint64, marker *int) ([]model.Todo, *int, error) {
	builder := dao.newSelectBuilder().
		OrderBy("id").
		Limit(limit + 1)
	if marker != nil {
		builder = builder.Where(sq.Gt{"id": marker})
	}
	var ms []model.Todo
	if err := dao.findManyByBuilder(&builder, &ms); err != nil {
		return nil, nil, err
	}
	var lastID *int
	if len(ms) > int(limit) {
		ms = ms[:limit]
		lastID = &ms[len(ms)-1].ID
	}
	return ms, lastID, nil
}

// Insert ...
func (dao todoDaoImpl) Insert(title string) (int, error) {
	m := model.Todo{Title: title}
	return m.ID, dao.insert(&m)
}

// Update ...
func (dao todoDaoImpl) Update(id int, title null.String, completed null.Bool) error {
	builder := dao.newUpdateBuilder().
		Where(sq.Eq{"id": id})
	if title.Valid {
		builder = builder.Set("title", title.String)
	}
	if completed.Valid {
		builder = builder.Set("completed", completed.Bool)
	}
	_, err := dao.updateByBuilder(&builder)
	return err
}

// Delete ...
func (dao todoDaoImpl) Delete(id int) error {
	m := model.Todo{ID: id}
	return dao.delete(&m)
}
