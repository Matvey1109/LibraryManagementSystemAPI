package repositories

import (
	"app/internal/storages"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var storage = storages.ExportStorage

type Repository struct {
	MemberRepository
	BookRepository
	BorrowingRepository
}

func NewRepository() *Repository{
	return &Repository{}
}

func GenerateID() string {
	objectID := primitive.NewObjectID()
	stringID := objectID.Hex()
	return stringID
}

var (
	ExportRepository = NewRepository()
)
