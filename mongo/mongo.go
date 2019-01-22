package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/chelium/simple-website/todo"
)

type todoRepository struct {
	db      string
	session *mgo.Session
}

func (r *todoRepository) Create(todo *todo.Todo) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")

	_, err := c.Upsert(bson.M{"id": todo.ID}, bson.M{"$set": todo})
	return err
}

func (r *todoRepository) Read(id string) (*todo.Todo, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")

	var result todo.Todo
	if err := c.Find(bson.M{"id": id}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, todo.ErrNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (r *todoRepository) ReadAll() ([]*todo.Todo, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")

	var result []*todo.Todo
	if err := c.Find(bson.M{}).All(&result); err != nil {
		return []*todo.Todo{}, nil
	}
	return result, nil
}

func (r *todoRepository) Update(id string, todo *todo.Todo) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")

	if err := c.Update(bson.M{"id": id}, todo); err != nil {
		if err == mgo.ErrNotFound {
			return todo.ErrNotFound
		}
		return err
	}

	return nil
}

func (r *todoRepository) Delete(id string) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")

	if err := c.Remove(bson.M{"id": id}); err != nil {
		if err == mgo.ErrNotFound {
			return todo.ErrNotFound
		}
		return err
	}

	return nil
}

func NewTodoRepository(db string, session *mgo.Session) (todo.TodoRepository, error) {
	r := &todoRepository{
		db:      db,
		session: session,
	}

	return r, nil
}
