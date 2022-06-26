package bolt

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
var errNewsNotFound = errors.New("news not found")

func (r *Repository) Post(con context.Context, news models.Project) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("project"))
		if b == nil {
			return errBicketNotFound
		}
		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		news.ID = int(id)
		buff, err := encode(news)
		if err != nil {
			return err
		}
		return b.Put(itob(int(id)), buff.Bytes())
	})
}
func (r *Repository) Get(con context.Context, id int) (models.Project, error) {
	var news models.Project
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("project"))
		if b == nil {
			return errBicketNotFound
		}
		temp := b.Get(itob(id))
		buff := bytes.NewBuffer(temp)
		var err error
		news, err = decode(*buff)
		return err
	})
	return news, err
}
func (r *Repository) Delete(con context.Context, id int) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("project"))
		return b.Delete(itob(id))
	})
}
func (r *Repository) Update(con context.Context, project models.Project) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("project"))
		if b == nil {
			return errBicketNotFound
		}
		n := b.Get(itob(project.ID))
		if n == nil {
			return errNewsNotFound
		}
		err := b.Delete(itob(project.ID))
		if err != nil {
			return err
		}
		buff, err := encode(project)
		return b.Put(itob(int(project.ID)), buff.Bytes())
	})
}
func (r *Repository) GetAllforClient(con context.Context) ([]models.Project, error) {
	var news []models.Project
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("project"))
		if b == nil {
			return errBicketNotFound
		}
		return b.ForEach(func(k, v []byte) error {
			buff := bytes.NewBuffer(v)
			projectTemp, err := decode(*buff)
			if err != nil {
				return err
			}
			if !projectTemp.Enable {
				return nil
			}
			news = append(news, projectTemp)
			return nil
		})
	})
	return news, err
}
func (r *Repository) GetAllForAdmin(con context.Context) ([]models.Project, error) {
	var news []models.Project
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("project"))
		if b == nil {
			return errBicketNotFound
		}
		return b.ForEach(func(k, v []byte) error {
			buff := bytes.NewBuffer(v)
			projectTemp, err := decode(*buff)
			if err != nil {
				return err
			}
			news = append(news, projectTemp)
			return nil
		})
	})
	return news, err
}

func decode(buff bytes.Buffer) (models.Project, error) {
	out := models.Project{}
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&out)
	return out, err
}

func encode(p models.Project) (bytes.Buffer, error) {
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
