package sqlite

import (
	"cgr/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		db: db,
	}
}

type newsMongo struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Enable  bool               `bson:"enable"`
	Image   string             `bson:"image"`
	Tags    []string           `bson:"tags"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
}

func (r *Repository) Post(con context.Context, news models.News) error {
	mongoNews := toNewsMongo(news)
	_, err := r.db.InsertOne(con, mongoNews)
	return err
}
func (r *Repository) Get(con context.Context, id string) (*models.News, error) {
	news := new(newsMongo)
	idb, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.db.FindOne(con, bson.M{
		"_id": idb,
	}).Decode(news)
	return toNewsModel(news), err
}
func (r *Repository) Delete(con context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.db.DeleteOne(con, bson.M{"_id": objID})
	return err
}
func (r *Repository) Update(con context.Context, news models.News) error {
	objID, _ := primitive.ObjectIDFromHex(news.ID)
	update := bson.M{"$set": bson.M{"enable": news.Enable, "image": news.Image, "tags": news.Tags, "title": news.Title, "content": news.Content}}
	_, err := r.db.UpdateByID(con, objID, update)
	return err
}
func (r *Repository) GetAllForClient(con context.Context) ([]*models.News, error) {
	cur, err := r.db.Find(con, bson.M{"enable": true})
	if err != nil {
		return []*models.News{}, err
	}
	out := []*models.News{}
	for cur.Next(con) {
		news := new(newsMongo)
		err := cur.Decode(news)
		if err != nil {
			return []*models.News{}, err
		}
		out = append(out, toNewsModel(news))
	}
	return out, nil
}
func (r *Repository) GetAllForAdmin(con context.Context) ([]*models.News, error) {
	cur, err := r.db.Find(con, bson.M{})
	if err != nil {
		return []*models.News{}, err
	}
	out := []*models.News{}
	for cur.Next(con) {
		news := new(newsMongo)
		err := cur.Decode(news)
		if err != nil {
			return []*models.News{}, err
		}
		out = append(out, toNewsModel(news))
	}
	return out, nil
}

func toNewsMongo(news models.News) *newsMongo {
	return &newsMongo{
		Enable:  news.Enable,
		Image:   news.Image,
		Tags:    news.Tags,
		Title:   news.Title,
		Content: news.Content,
	}
}

func toNewsModel(news *newsMongo) *models.News {
	return &models.News{
		ID:      news.ID.Hex(),
		Enable:  news.Enable,
		Image:   news.Image,
		Tags:    news.Tags,
		Title:   news.Title,
		Content: news.Content,
	}
}
