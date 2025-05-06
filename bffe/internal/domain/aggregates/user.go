package domainaggregates

import (
	"strings"

	domainentities "github.com/o-rensa/jalikod-backend/bffe/internal/domain/entities"
)

type User struct {
	id            int32
	firstName     string
	middleInitial *string
	surname       string
	nameExtension *string
	username      string
	domainentities.FullAuditedEntity
}

func NewUser(id int32, username string) *User {
	u := &User{}
	u.id = id
	u.username = username
	u.IsDeleted = false
	u.IsActive = true
	return u
}

// setters
func (u *User) SetFullName(firstname string, middle string, surname string, extension string) {
	u.firstName = firstname
	u.surname = surname
	if strings.TrimSpace(middle) == "" {
		u.middleInitial = nil
	} else {
		m := middle
		u.middleInitial = &m
	}

	if strings.TrimSpace(extension) == "" {
		u.nameExtension = nil
	} else {
		e := extension
		u.nameExtension = &e
	}
}

// getters
func (u *User) GetId() int32 {
	return u.id
}

func (u *User) GetFullname() (firstname string, middleInitial *string, surname string, nameExtension *string) {
	firstname = u.firstName
	middleInitial = u.middleInitial
	surname = u.surname
	nameExtension = u.nameExtension
	return
}

func (u *User) GetUsername() string {
	return u.username
}
