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

func (r *todoRepository) Create(userID string, todo *todo.Todo) (string, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")
	todo.Owner = userID
	todo.CreatedBy = userID
	todo.AssignedTo = userID

	_, err := c.Upsert(bson.M{"id": todo.ID}, bson.M{"$set": todo})
	return todo.ID, err
}

func (r *todoRepository) Read(userID, todoID string) (*todo.Todo, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")

	var result todo.Todo
	if err := c.Find(bson.M{"id": todoID}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, todo.ErrNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (r *todoRepository) ReadAll(userID string) ([]*todo.Todo, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")

	var result []*todo.Todo
	if err := c.Find(bson.M{}).All(&result); err != nil {
		return []*todo.Todo{}, nil
	}
	return result, nil
}

func (r *todoRepository) Update(userID, todoID string, t *todo.Todo) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")

	if err := c.Update(bson.M{"id": todoID}, t); err != nil {
		if err == mgo.ErrNotFound {
			return todo.ErrNotFound
		}
		return err
	}

	return nil
}

func (r *todoRepository) Delete(userID, todoID string) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("todo")

	if err := c.Remove(bson.M{"id": todoID}); err != nil {
		if err == mgo.ErrNotFound {
			return todo.ErrNotFound
		}
		return err
	}

	return nil
}

// NewTodoRepository returns a new instance of a MongoDB todo repository.
func NewTodoRepository(db string, session *mgo.Session) (todo.TodoRepository, error) {
	r := &todoRepository{
		db:      db,
		session: session,
	}

	return r, nil
}
