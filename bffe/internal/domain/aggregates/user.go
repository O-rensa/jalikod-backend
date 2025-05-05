package domainaggregates

import (
	"github.com/google/uuid"
	domainentities "github.com/o-rensa/jalikod-backend/bffe/internal/domain/entities"
)

type User struct {
	ID            int32
	FirstName     string
	MiddleInitial *string
	Surname       string
	NameExtension *string
	Username      string
	Password      *string
	SecurityStamp uuid.NullUUID
	domainentities.FullAuditedEntity
}
