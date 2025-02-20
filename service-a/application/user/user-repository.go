package user

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github/herochi/orbi/service-a/application/user/interfaces"
	"github/herochi/orbi/service-a/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db             *mongo.Database
	collection     *mongo.Collection
	collectionRole *mongo.Collection
}

func NewUserRepository(c *mongo.Database) interfaces.UserRepository {
	return &userRepository{db: c, collection: c.Collection("users"), collectionRole: c.Collection("roles")}
}

func (u *userRepository) Save(ctx context.Context, user *domain.User) (string, error) {

	user.CreatedAt = timeNow()
	user.UpdatedAt = timeNow()

	result, err := u.collection.InsertOne(ctx, user)

	if err != nil {
		return "", err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user.ID.Hex(), nil
}

func timeNow() time.Time {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	return now
}

func (u *userRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	idT, _ := primitive.ObjectIDFromHex(id)

	err := u.collection.FindOne(ctx, bson.D{{"_id", idT}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, userID string, dataToUpdate *domain.UpdateUser) (*domain.User, error) {

	var updatedUser domain.User

	id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", u.fillBsonD(dataToUpdate)},
	}
	after := options.After
	options := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := u.collection.FindOneAndUpdate(ctx, filter, update, &options)

	err := result.Decode(&updatedUser)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (u *userRepository) fillBsonD(dataToUpdate *domain.UpdateUser) primitive.D {
	update := bson.D{}

	if dataToUpdate.Name != "" {
		update = append(update, bson.E{Key: "name", Value: dataToUpdate.Name})
	}

	if dataToUpdate.Email != "" {
		update = append(update, bson.E{Key: "email", Value: dataToUpdate.Email})
	}

	update = append(update, bson.E{Key: "updatedAt", Value: timeNow()})
	return update
}
