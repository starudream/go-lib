package sqlite

import (
	"context"
	"time"

	"gorm.io/gorm"
	gormSD "gorm.io/plugin/soft_delete"
)

// DeletedAt
// `gorm:"uniqueIndex:udx_account_deleted_at" json:"deletedAt,omitempty"`
type DeletedAt = gormSD.DeletedAt

type Meta struct {
	Id        string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"createdAt"`
	CreatedBy string    `gorm:"type:varchar(36)"     json:"createdBy"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updatedAt,omitempty"`
	UpdatedBy string    `gorm:"type:varchar(36)"     json:"updatedBy,omitempty"`
}

var (
	GenId  func(context.Context) string
	GenUid func(context.Context) string
)

func (t *Meta) BeforeCreate(tx *gorm.DB) error {
	if GenId != nil && t.Id == "" {
		t.Id = GenId(tx.Statement.Context)
	}
	if GenUid != nil && t.CreatedBy == "" {
		t.CreatedBy = GenUid(tx.Statement.Context)
	}
	return nil
}

func (t *Meta) BeforeUpdate(tx *gorm.DB) error {
	if GenUid != nil && t.UpdatedBy == "" {
		t.UpdatedBy = GenUid(tx.Statement.Context)
	}
	return nil
}

func (t *Meta) BeforeDelete(tx *gorm.DB) error {
	if GenUid != nil && t.UpdatedBy == "" {
		t.UpdatedBy = GenUid(tx.Statement.Context)
	}
	return nil
}
