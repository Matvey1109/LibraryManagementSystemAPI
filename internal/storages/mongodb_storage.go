package storages

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ! Implements interface Storage
type MongoDBStorage struct {
	client               *mongo.Client
	membersCollection    *mongo.Collection
	booksCollection      *mongo.Collection
	borrowingsCollection *mongo.Collection
}

var _ Storage = (*MongoDBStorage)(nil) // Checker

// * Member
func (ms *MongoDBStorage) GetAllMembersStorage() ([]models.Member, error) {
	cursor, err := ms.membersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var members []models.Member
	for cursor.Next(context.Background()) {
		var (
			mongoMemberMap map[string]interface{}
			member         models.Member
		)

		err := cursor.Decode(&mongoMemberMap)
		if err != nil {
			return nil, err
		}
		for key, value := range mongoMemberMap {
			if key == "_id" {
				member.ID = value.(primitive.ObjectID).Hex()
			} else if key == "name" {
				member.Name = value.(string)
			} else if key == "address" {
				member.Address = value.(string)
			} else if key == "email" {
				member.Email = value.(string)
			} else if key == "createdAt" {
				curTime := value.(primitive.DateTime).Time()
				member.CreatedAt, _ = time.Parse(time.DateTime, curTime.Format(time.DateTime))
			}
		}
		members = append(members, member)
	}
	return members, nil
}

func (ms *MongoDBStorage) GetMemberStorage(id string) (models.Member, error) {
	var member models.Member
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return member, err
	}
	filter := bson.M{"_id": objectID}
	err = ms.membersCollection.FindOne(context.Background(), filter).Decode(&member)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return member, errors.New("member not found")
		}
		return member, err
	}
	member.ID = id
	return member, nil
}

func (ms *MongoDBStorage) AddMemberStorage(member models.Member) error {
	newMember := bson.D{
		{Key: "name", Value: member.Name},
		{Key: "address", Value: member.Address},
		{Key: "email", Value: member.Email},
		{Key: "createdAt", Value: member.CreatedAt},
	}
	fmt.Println(member.CreatedAt)
	_, err := ms.membersCollection.InsertOne(context.Background(), newMember)
	return err
}

func (ms *MongoDBStorage) UpdateMemberStorage(id string, member models.Member) error {
	return nil
}

func (ms *MongoDBStorage) DeleteMemberStorage(id string) error {
	return nil
}

// * Book
func (ms *MongoDBStorage) GetAllBooksStorage() ([]models.Book, error) {
	return nil, nil
}

func (ms *MongoDBStorage) GetBookStorage(id string) (models.Book, error) {
	return models.Book{}, nil
}

func (ms *MongoDBStorage) AddBookStorage(book models.Book) error {
	return nil
}

func (ms *MongoDBStorage) UpdateBookStorage(id string, book models.Book) error {
	return nil
}

func (ms *MongoDBStorage) DeleteBookStorage(id string) error {
	return nil
}

// * Borrowing
func (ms *MongoDBStorage) GetAllBorrowingsStorage() ([]models.Borrowing, error) {
	return nil, nil
}

func (ms *MongoDBStorage) GetBorrowingStorage(id string) (models.Borrowing, error) {
	return models.Borrowing{}, nil
}

func (ms *MongoDBStorage) AddBorrowingStorage(borrowing models.Borrowing) error {
	return nil
}

func (ms *MongoDBStorage) UpdateBorrowingStorage(id string, borrowing models.Borrowing) error {
	return nil
}

func (ms *MongoDBStorage) DeleteBorrowingStorage(id string) error {
	return nil
}
