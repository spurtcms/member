package postgres


import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TblMemberGroup struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name        string    `gorm:"type:character varying"`
	Slug        string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	IsActive    int       `gorm:"type:integer"`
	IsDeleted   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
}

type TblMember struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Uuid             string    `gorm:"type:character varying"`
	FirstName        string    `gorm:"type:character varying"`
	LastName         string    `gorm:"type:character varying"`
	Email            string    `gorm:"type:character varying"`
	MobileNo         string    `gorm:"type:character varying"`
	IsActive         int       `gorm:"type:integer"`
	ProfileImage     string    `gorm:"type:character varying"`
	ProfileImagePath string    `gorm:"type:character varying"`
	LastLogin        int       `gorm:"type:integer"`
	MemberGroupId    int       `gorm:"type:integer"`
	Password         string    `gorm:"type:character varying"`
	Username         string    `gorm:"DEFAULT:NULL"`
	Otp              int       `gorm:"DEFAULT:NULL"`
	OtpExpiry        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	LoginTime        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
}

type TblMemberProfile struct {
	Id              int               `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId        int               `gorm:"type:integer"`
	ProfilePage     string            `gorm:"type:character varying"`
	ProfileName     string            `gorm:"type:character varying"`
	ProfileSlug     string            `gorm:"type:character varying"`
	CompanyLogo     string            `gorm:"type:character varying"`
	CompanyName     string            `gorm:"type:character varying"`
	CompanyLocation string            `gorm:"type:character varying"`
	About           string            `gorm:"type:character varying"`
	Linkedin        string            `gorm:"type:character varying"`
	Website         string            `gorm:"type:character varying"`
	Twitter         string            `gorm:"type:character varying"`
	SeoTitle        string            `gorm:"type:character varying"`
	SeoDescription  string            `gorm:"type:character varying"`
	SeoKeyword      string            `gorm:"type:character varying"`
	MemberDetails   datatypes.JSONMap `json:"memberDetails" gorm:"column:member_details;type:jsonb"`
	ClaimStatus     int               `gorm:"DEFAULT:0;type:integer"`
	CreatedBy       int               `gorm:"type:integer"`
	CreatedOn       time.Time         `gorm:"type:timestamp without time zone"`
	ModifiedBy      int               `gorm:"DEFAULT:NULL"`
	ModifiedOn      time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted       int               `gorm:"DEFAULT:0"`
	DeletedBy       int               `gorm:"DEFAULT:NULL"`
	DeletedOn       time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblMemberNotesHighlights struct {
	Id                      int               `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId                int               `gorm:"type:integer"`
	PageId                  int               `gorm:"type:integer"`
	NotesHighlightsContent  string            `gorm:"type:character varying"`
	NotesHighlightsType     string            `gorm:"type:character varying"`
	HighlightsConfiguration datatypes.JSONMap `gorm:"type:jsonb"`
	CreatedBy               int               `gorm:"type:integer"`
	CreatedOn               time.Time         `gorm:"type:timestamp without time zone"`
	ModifiedBy              int               `gorm:"type:integer"`
	ModifiedOn              time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy               int               `gorm:"type:integer"`
	DeletedOn               time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted               int               `gorm:"type:integer"`
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
