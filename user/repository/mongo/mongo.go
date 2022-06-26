package mongo

import (
	"bytes"
	"cgr/models"
	"context"
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/boltdb/bolt"
)

type Repository struct {
	db *bolt.DB
}

func NewRepository(db *bolt.DB) *Repository {
	return &Repository{
		db: db,
	}
}

var errBicketNotFound = errors.New("Bucket not found")
var errUserNotFound = errors.New("user not found")

func (r *Repository) Create(con context.Context, user models.User) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		if b == nil {
			return errBicketNotFound
		}
		buff, err := encode(user)
		if err != nil {
			return err
		}
		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		return b.Put(itob(int(id)), buff.Bytes())
	})
}
func (r *Repository) Delete(con context.Context, id int) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		b1 := tx.Bucket([]byte("token"))
		if b == nil {
			return errBicketNotFound
		}
		if b1 == nil {
			return errBicketNotFound
		}
		if b1.Get(itob(id)) != nil {
			return b1.Delete(itob(id))
		}
		return b.Delete(itob(id))
	})
}
func (r *Repository) Update(con context.Context, user models.User) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		if b == nil {
			return errBicketNotFound
		}
		n := b.Get(itob(user.ID))
		if n == nil {
			return errUserNotFound
		}
		err := b.Delete(itob(user.ID))
		if err != nil {
			return err
		}
		buff, err := encode(user)
		return b.Put(itob(user.ID), buff.Bytes())
	})
}

func (r *Repository) UpdateSesion(con context.Context, id int, session string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("token"))
		if b == nil {
			return errBicketNotFound
		}
		n := b.Get(itob(id))
		if n == nil {
			b.Put(itob(id), []byte(session))
		}
		err := b.Delete(itob(id))
		if err != nil {
			return err
		}
		return b.Put(itob(id), []byte(session))
	})
}

func (r *Repository) GetAll(con context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		if b == nil {
			return errBicketNotFound
		}
		return b.ForEach(func(k, v []byte) error {
			buff := bytes.NewBuffer(v)
			user, err := decode(*buff)
			if err != nil {
				return err
			}
			users = append(users, user)
			return nil
		})
	})
	return users, err
}
func (r *Repository) Get(con context.Context, name string) (models.User, error) {
	var user models.User
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		b1 := tx.Bucket([]byte("user_name"))

		id := b1.Get([]byte(name))
		if id == nil {
			return errUserNotFound
		}
		temp := b.Get(id)
		buff := bytes.NewBuffer(temp)
		var err error
		user, err = decode(*buff)
		return err
	})
	return user, err
}

func (r *Repository) GetByToken(c context.Context, token string) (bool, error) {
	access := false
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("token"))
		b.ForEach(func(k, v []byte) error {
			if string(v) == token {
				access = true
			}
			return nil
		})
		return nil
	})
	return access, err
}

func decode(buff bytes.Buffer) (models.User, error) {
	out := models.User{}
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&out)
	return out, err
}

func encode(p models.User) (bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	return buf, err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
