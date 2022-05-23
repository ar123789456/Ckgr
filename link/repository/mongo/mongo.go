package mongo

import (
	"cgr/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		db: db,
	}
}

type Link struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Logo  string             `bson:"logo"`
	Title string             `bson:"title"`
	Url   string             `bson:"url"`
}

func (r *Repository) Create(con context.Context, link models.Link) error {
	mongoLink := toLink(&link)
	_, err := r.db.InsertOne(con, mongoLink)
	return err
}
func (r *Repository) Delete(con context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.db.DeleteOne(con, bson.M{"_id": objID})
	return err
}
func (r *Repository) GetAll(con context.Context) ([]*models.Link, error) {
	cur, err := r.db.Find(con, bson.M{})
	if err != nil {
		return []*models.Link{}, err
	}
	out := []*models.Link{}
	for cur.Next(con) {
		news := new(Link)
		err := cur.Decode(news)
		if err != nil {
			return []*models.Link{}, err
		}
		out = append(out, toLinkMongo(news))
	}
	return out, nil
}

func toLinkMongo(l *Link) *models.Link {
	return &models.Link{
		ID:    l.ID.Hex(),
		Logo:  l.Logo,
		Title: l.Title,
		Url:   l.Url,
	}
}

func toLink(l *models.Link) *Link {
	return &Link{
		Logo:  l.Logo,
		Title: l.Title,
		Url:   l.Url,
	}
}
