package xml_serializers

import (
	"encoding/xml"
	"time"

	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/schemas"
)

type MemberXML struct {
	XMLName   xml.Name  `xml:"member"`
	ID        string    `xml:"id"`
	Name      string    `xml:"name"`
	Address   string    `xml:"address"`
	Email     string    `xml:"email"`
	CreatedAt time.Time `xml:"createdAt"`
}

func SerializeMemberToXML(member schemas.Member) ([]byte, error) {
	memberXML := MemberXML{
		ID:        member.ID,
		Name:      member.Name,
		Address:   member.Address,
		Email:     member.Email,
		CreatedAt: member.CreatedAt,
	}

	data, err := xml.MarshalIndent(memberXML, "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func DeserializeMemberFromXML(data []byte) (schemas.Member, error) {
	var memberXML MemberXML
	err := xml.Unmarshal(data, &memberXML)
	if err != nil {
		return schemas.Member{}, err
	}
	return schemas.Member{
		ID:        memberXML.ID,
		Name:      memberXML.Name,
		Address:   memberXML.Address,
		Email:     memberXML.Email,
		CreatedAt: memberXML.CreatedAt,
	}, nil
}
