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

type mongoUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	FullName string             `bson:"name"`
	Special  string             `bson:"special"`
	Session  string             `bson:"session"`
}

func (r *Repository) Create(con context.Context, user *models.User) error {
	mongoUser := toUser(user)
	_, err := r.db.InsertOne(con, mongoUser)
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
func (r *Repository) Update(con context.Context, user *models.User) error {
	userMongo := toUser(user)
	filtr := bson.M{"_id": userMongo.ID}
	_, err := r.db.ReplaceOne(con, filtr, user)
	return err
}

func (r *Repository) UpdateSesion(con context.Context, id string, session string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": bson.M{"session": session}}
	_, err := r.db.UpdateByID(con, objID, update)
	return err
}

func (r *Repository) GetAll(con context.Context) ([]models.User, error) {
	cur, err := r.db.Find(con, bson.M{"enable": true})
	if err != nil {
		return []models.User{}, err
	}
	out := []models.User{}
	for cur.Next(con) {
		news := new(mongoUser)
		err := cur.Decode(news)
		if err != nil {
			return out, err
		}
		out = append(out, toUserModel(news))
	}
	return out, nil
}
func (r *Repository) Get(con context.Context, id string) (models.User, error) {
	news := new(mongoUser)
	idb, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	err = r.db.FindOne(con, bson.M{
		"_id": idb,
	}).Decode(news)
	return toUserModel(news), err
}

func toUser(user *models.User) *mongoUser {
	objID, _ := primitive.ObjectIDFromHex(user.ID)
	return &mongoUser{
		ID:       objID,
		FullName: user.FullName,
		Special:  user.Special,
	}
}

func toUserModel(user *mongoUser) models.User {
	return models.User{
		ID:       user.ID.Hex(),
		FullName: user.FullName,
		Special:  user.Special,
	}
}
