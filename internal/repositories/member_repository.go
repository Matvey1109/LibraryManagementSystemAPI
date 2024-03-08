package repositories

import (
	"app/internal/models"
	"time"
)

type MemberRepository struct{}

func (mr *MemberRepository) GetAllMembers() ([]models.Member, error) {
	members, err := storage.GetAllMembersStorage()
	if err != nil {
		return members, err
	}
	return members, nil
}

func (mr *MemberRepository) GetMember(id string) (models.Member, error) {
	member, err := storage.GetMemberStorage(id)
	if err != nil {
		return member, err
	}
	return member, nil
}

func (mr *MemberRepository) AddMember(name string, address string, email string) error {
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
		return err
	}
	return nil
}

func (mr *MemberRepository) UpdateMember(id string, name string, address string, email string) error {
	member, err := storage.GetMemberStorage(id)
	if err != nil {
		return err
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
		return err
	}
	return nil
}

func (mr *MemberRepository) DeleteMember(id string) error {
	err := storage.DeleteMemberStorage(id)
	if err != nil {
		return err
	}
	return nil
}
