package sqlite

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

var errBucketNotFound = errors.New("Bucket not found")
var errNewsNotFound = errors.New("news not found")

func (r *Repository) Post(con context.Context, news models.News) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("news"))
		if b == nil {
			return errBucketNotFound
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
func (r *Repository) Get(con context.Context, id int) (models.News, error) {
	var news models.News
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("news"))
		if b == nil {
			return errBucketNotFound
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
		b := tx.Bucket([]byte("news"))
		return b.Delete(itob(id))
	})
}
func (r *Repository) Update(con context.Context, news models.News) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("news"))
		if b == nil {
			return errBucketNotFound
		}
		n := b.Get(itob(news.ID))
		if n == nil {
			return errNewsNotFound
		}
		err := b.Delete(itob(news.ID))
		if err != nil {
			return err
		}
		buff, err := encode(news)
		return b.Put(itob(int(news.ID)), buff.Bytes())
	})
}
func (r *Repository) GetAllForClient(con context.Context) ([]models.News, error) {
	var news []models.News
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("news"))
		if b == nil {
			return errBucketNotFound
		}
		return b.ForEach(func(k, v []byte) error {
			buff := bytes.NewBuffer(v)
			newsTemp, err := decode(*buff)
			if err != nil {
				return err
			}
			if !newsTemp.Enable {
				return nil
			}
			news = append(news, newsTemp)
			return nil
		})
	})
	return news, err
}
func (r *Repository) GetAllForAdmin(con context.Context) ([]models.News, error) {
	var news []models.News
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("news"))
		if b == nil {
			return errBucketNotFound
		}
		return b.ForEach(func(k, v []byte) error {
			buff := bytes.NewBuffer(v)
			newsTemp, err := decode(*buff)
			if err != nil {
				return err
			}
			news = append(news, newsTemp)
			return nil
		})
	})
	return news, err
}

func decode(buff bytes.Buffer) (models.News, error) {
	out := models.News{}
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&out)
	return out, err
}

func encode(p models.News) (bytes.Buffer, error) {
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
