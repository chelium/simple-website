package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/chelium/simple-website/user"
)

type userRepository struct {
	db      string
	session *mgo.Session
}

func (r *userRepository) Create(user *user.User) (string, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("user")

	if _, err := c.Upsert(bson.M{"id": user.ID}, bson.M{"$set": user}); err != nil {
		return "", err
	}

	return user.ID, nil
}

func (r *userRepository) ReadByID(id string) (*user.User, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("user")

	var result user.User
	if err := c.FindId(id).One(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *userRepository) ReadByName(username string) (*user.User, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("user")

	var result user.User
	if err := c.Find(bson.M{"username": username}).One(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *userRepository) Update(id string, t *user.User) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("user")

	if err := c.Update(bson.M{"id": id}, t); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(id string) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("user")

	if err := c.Remove(bson.M{"id": id}); err != nil {
		return err
	}

	return nil
}

// NewUserRepository returns a new instance of a MongoDB user repository.
func NewUserRepository(db string, session *mgo.Session) (user.UserRepository, error) {
	r := &userRepository{
		db:      db,
		session: session,
	}

	return r, nil
}
