package controller

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/suzujun/yatteiki-cloud/goapp/dao"
)

type todo struct {
	todoDao dao.TodoDao
}

func NewTodo(todoDao dao.TodoDao) todo {
	return todo{todoDao: todoDao}
}

func (t todo) GetList(c *gin.Context) {
	var limit uint64 = 100
	if v, err := strconv.Atoi(c.Query("limit")); err == nil {
		limit = uint64(v)
	}
	var marker *int
	if v, err := strconv.Atoi(c.Query("marker")); err == nil {
		marker = &v
	}
	todos, marker, err := t.todoDao.FindAll(limit, marker)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data":       todos,
		"nextMarker": marker,
	})
}

func (t todo) Post(c *gin.Context) {
	body := struct {
		Note string `json:"note"`
	}{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	insertedID, err := t.todoDao.Insert(body.Note)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"insertedId": insertedID,
	})
}

func (t todo) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todos id"})
		return
	}
	todo, err := t.todoDao.FindByID(id)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "not found todo"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": todo,
	})
}

func (t todo) Put(c *gin.Context) {
	body := struct {
		Note string `json:"note"`
	}{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todos id"})
		return
	}
	if err := t.todoDao.Update(id, body.Note); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{})
}

func (t todo) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid todos id"})
		return
	}
	if err := t.todoDao.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{})
}
