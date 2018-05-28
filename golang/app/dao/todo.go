package dao

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/suzujun/yatteiki-cloud/golang/app/model"
)

type (
	// TodoDao ...
	TodoDao interface {
		FindByID(id int) (*model.Todo, error)
		FindAll(limit uint64, marker *int) ([]model.Todo, *int, error)
		Insert(note string) error
		Update(id int, note string) error
		Delete(id int) error
	}
	todoDaoImpl struct {
		baseDao
	}
)

// NewTodoDao ...
func NewTodoDao() TodoDao {
	return &todoDaoImpl{
		baseDao: NewDao(model.Todo{}),
	}
}

// FindByID ...
func (dao todoDaoImpl) FindByID(id int) (*model.Todo, error) {
	builder := dao.newSelectBuilder().
		Where(sq.Eq{"id": id})
	var m *model.Todo
	return m, dao.findOneByBuilder(&builder, m)
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
		ms = ms[:limit-1]
		lastID = &ms[limit-1].ID
	}
	return ms, lastID, nil
}

// Insert ...
func (dao todoDaoImpl) Insert(note string) error {
	m := model.Todo{Note: note}
	return dao.insert(&m)
}

// Update ...
func (dao todoDaoImpl) Update(id int, note string) error {
	m := model.Todo{ID: id, Note: note}
	return dao.update(&m)
}

// Delete ...
func (dao todoDaoImpl) Delete(id int) error {
	m := model.Todo{ID: id}
	return dao.delete(&m)
}
