package storages

import (
	"errors"

	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/models"
	"github.com/Matvey1109/LibraryManagementSystemAPI/pkg/loadenv"
)

// ! Abstract Factory
type StorageFactory interface {
	CreateStorage() (Storage, error)
}

func GetStorageFactory() (StorageFactory, error) {
	typeOfStorage, _, _ := loadenv.LoadEnv()

	if typeOfStorage == "local" {
		return &LocalStorageFactory{}, nil
	}

	return nil, errors.New("typeOfStorage not found")
}

// ! Concrete Factories
type LocalStorageFactory struct{} // * Implements interface StorageFactory

func (f *LocalStorageFactory) CreateStorage() (Storage, error) {
	return &LocalStorage{
		members:    []models.Member{},
		books:      []models.Book{},
		borrowings: []models.Borrowing{},
	}, nil
}

// ! Abstract Product
type Storage interface {
	GetAllMembersStorage() ([]models.Member, error)
	GetMemberStorage(id string) (models.Member, error)
	AddMemberStorage(member models.Member) error
	UpdateMemberStorage(id string, member models.Member) error
	DeleteMemberStorage(id string) error

	GetAllBooksStorage() ([]models.Book, error)
	GetBookStorage(id string) (models.Book, error)
	AddBookStorage(book models.Book) error
	UpdateBookStorage(id string, book models.Book) error
	DeleteBookStorage(id string) error

	GetAllBorrowingsStorage() ([]models.Borrowing, error)
	GetBorrowingStorage(id string) (models.Borrowing, error)
	AddBorrowingStorage(borrowing models.Borrowing) error
	UpdateBorrowingStorage(id string, borrowing models.Borrowing) error
	DeleteBorrowingStorage(id string) error
}

var (
	ExportStorageFactory, _ = GetStorageFactory()
	ExportStorage, _        = ExportStorageFactory.CreateStorage()
)
