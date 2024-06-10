package mysql

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TblMemberGroup struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Name        string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	IsActive    int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedBy   int       `gorm:"type:int"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblMember struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	Uuid             string    `gorm:"type:varchar(255)"`
	FirstName        string    `gorm:"type:varchar(255)"`
	LastName         string    `gorm:"type:varchar(255)"`
	Email            string    `gorm:"type:varchar(255)"`
	MobileNo         string    `gorm:"type:varchar(255)"`
	IsActive         int       `gorm:"type:int"`
	ProfileImage     string    `gorm:"type:varchar(255)"`
	ProfileImagePath string    `gorm:"type:varchar(255)"`
	LastLogin        int       `gorm:"type:int"`
	MemberGroupId    int       `gorm:"type:int"`
	Password         string    `gorm:"type:varchar(255)"`
	Username         string    `gorm:"type:varchar(255)"`
	Otp              int       `gorm:"DEFAULT:NULL;type:int"`
	OtpExpiry        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	LoginTime        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:int"`
	DeletedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL;type:int"`
	CreatedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:int"`
	ModifiedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblMemberNotesHighlights struct {
	Id                      int               `gorm:"primaryKey;auto_increment"`
	MemberId                int               `gorm:"type:int"`
	PageId                  int               `gorm:"type:int"`
	NotesHighlightsContent  string            `gorm:"type:varchar(255)"`
	NotesHighlightsType     string            `gorm:"type:varchar(255)"`
	HighlightsConfiguration datatypes.JSONMap `gorm:"type:jsonb"`
	CreatedBy               int               `gorm:"type:int"`
	CreatedOn               time.Time         `gorm:"type:datetime"`
	ModifiedBy              int               `gorm:"type:int"`
	ModifiedOn              time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy               int               `gorm:"type:int"`
	DeletedOn               time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted               int               `gorm:"type:int"`
}

type TblMemberProfile struct {
	Id              int               `gorm:"primaryKey;auto_increment"`
	MemberId        int               `gorm:"type:int"`
	ProfilePage     string            `gorm:"type:varchar(255)"`
	ProfileName     string            `gorm:"type:varchar(255)"`
	ProfileSlug     string            `gorm:"type:varchar(255)"`
	CompanyLogo     string            `gorm:"type:varchar(255)"`
	CompanyName     string            `gorm:"type:varchar(255)"`
	CompanyLocation string            `gorm:"type:varchar(255)"`
	About           string            `gorm:"type:varchar(255)"`
	Linkedin        string            `gorm:"type:varchar(255)"`
	Website         string            `gorm:"type:varchar(255)"`
	Twitter         string            `gorm:"type:varchar(255)"`
	SeoTitle        string            `gorm:"type:varchar(255)"`
	SeoDescription  string            `gorm:"type:varchar(255)"`
	SeoKeyword      string            `gorm:"type:varchar(255)"`
	MemberDetails   datatypes.JSONMap `json:"memberDetails" gorm:"column:member_details;type:jsonb"`
	ClaimStatus     int               `gorm:"DEFAULT:0;type:integer"`
	CreatedBy       int               `gorm:"type:int"`
	CreatedOn       time.Time         `gorm:"type:datetime"`
	ModifiedBy      int               `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn      time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted       int               `gorm:"DEFAULT:0"`
	DeletedBy       int               `gorm:"DEFAULT:NULL;type:int"`
	DeletedOn       time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
}


type TblMemberSetting struct {
	Id                int       `gorm:"primaryKey;auto_increment;"`
	AllowRegistration int       `gorm:"type:int"`
	MemberLogin       string    `gorm:"type:varchar(255)"`
	ModifiedBy        int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	NotificationUsers string    `gorm:"type:varchar(255)"`
}

// MigrateTable creates this package related tables in your database
func MigrateTables(db *gorm.DB) {

	db.AutoMigrate(&TblMemberGroup{}, &TblMember{}, &TblMemberNotesHighlights{}, &TblMemberProfile{},&TblMemberSetting{})

}
