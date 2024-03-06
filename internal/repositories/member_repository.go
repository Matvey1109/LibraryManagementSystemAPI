package repositories

import (
	"app/internal/models"
	"log"
	"time"
)

type MemberRepository struct{}

func (mr *MemberRepository) GetAllMembers() []models.Member {
	members, err := storage.GetAllMembersStorage()
	if err != nil {
		log.Fatal(err)
	}
	return members
}

func (mr *MemberRepository) GetMember(id string) models.Member {
	member, err := storage.GetMemberStorage(id)
	if err != nil {
		log.Fatal(err)
	}
	return member
}

func (mr *MemberRepository) AddMember(name string, address string, email string) {
	newID := GenerateID()
	newMember := models.Member{
		ID:        newID,
		Name:      name,
		Address:   address,
		Email:     email,
		CreatedAt: time.Now(),
	}
	err := storage.AddMemberStorage(newMember)
	if err != nil {
		log.Fatal(err)
	}
}

func (mr *MemberRepository) UpdateMember(id string, name string, address string, email string) {
	member, err := storage.GetMemberStorage(id)
	if err != nil {
		log.Fatal(err)
	}

	if name != "" {
		member.Name = name
	}
	if address != "" {
		member.Address = address
	}
	if email != "" {
		member.Email = email
	}

	err = storage.UpdateMemberStorage(id, member)
	if err != nil {
		log.Fatal(err)
	}
}

func (mr *MemberRepository) DeleteMember(id string) {
	err := storage.DeleteMemberStorage(id)
	if err != nil {
		log.Fatal(err)
	}
}
