package domainentities

import "time"

type CreationAudited struct {
	CreationTime  time.Time
	CreatorUserid int32
}

type ModificationAudited struct {
	LastModificationTime time.Time
	LastModifierUserid   int32
}

type SoftDelete struct {
	IsDeleted     bool
	DeletionTime  time.Time
	DeleterUserId int32
}

type FullAuditedEntity struct {
	CreationAudited
	ModificationAudited
	SoftDelete
	IsActive bool
}
