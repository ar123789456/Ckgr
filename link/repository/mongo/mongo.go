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

var errNotBucket = errors.New("invalid bucket name")

func NewRepository(db *bolt.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(con context.Context, link models.Link) error {
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("link"))
		if b == nil {
			return errNotBucket
		}

		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		link.ID = int(id)
		buff, err := encode(link)
		if err != nil {
			return err
		}
		err = b.Put(itob(int(id)), buff.Bytes())
		return err
	})
	return err
}
func (r *Repository) Delete(con context.Context, id int) error {
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("link"))
		if b == nil {
			return errNotBucket
		}
		return b.Delete(itob(id))
	})
	return err
}
func (r *Repository) GetAll(con context.Context) ([]models.Link, error) {
	var allLink []models.Link
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("link"))
		if b == nil {
			return errNotBucket
		}
		return b.ForEach(func(k, v []byte) error {
			buff := bytes.NewBuffer(v)
			link, err := decode(*buff)
			if err != nil {
				return err
			}
			allLink = append(allLink, link)
			return nil
		})
	})
	return allLink, err
}

func decode(buff bytes.Buffer) (models.Link, error) {
	out := models.Link{}
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&out)
	return out, err
}

func encode(p models.Link) (bytes.Buffer, error) {
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
