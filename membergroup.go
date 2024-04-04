package member

import (
	"errors"
	"time"
)

type tblmember struct {
	Id               int `gorm:"primaryKey;auto_increment;"`
	Uuid             string
	FirstName        string
	LastName         string
	Email            string
	MobileNo         string
	IsActive         int
	ProfileImage     string
	ProfileImagePath string
	LastLogin        int
	IsDeleted        int
	DeletedOn        time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	CreatedOn        time.Time `gorm:"DEFAULT:NULL"`
	CreatedDate      string    `gorm:"-"`
	CreatedBy        int
	ModifiedOn       time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	MemberGroupId    int
	GroupName        string `gorm:"-:migration;<-:false"`
	Password         string
	DateString       string    `gorm:"-"`
	Username         string    `gorm:"DEFAULT:NULL"`
	Otp              int       `gorm:"DEFAULT:NULL"`
	OtpExpiry        time.Time `gorm:"DEFAULT:NULL"`
	ModifiedDate     string    `gorm:"-"`
	NameString       string    `gorm:"-"`
	LoginTime        time.Time `gorm:"DEFAULT:NULL"`
}

type tblmembergroup struct {
	Id          int `gorm:"primaryKey;auto_increment;"`
	Name        string
	Slug        string
	Description string
	IsActive    int
	IsDeleted   int
	CreatedOn   time.Time `gorm:"DEFAULT:NULL"`
	CreatedBy   int
	ModifiedOn  time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	DateString  string    `gorm:"-"`
}

// Function ListMemberGroup pass the arguments of limit,offset and filter (eg. keywords)
// It will return the all membergroup lists
func (member *Member) ListMemberGroup(offset, limit int, filter Filter) (membergroup []tblmembergroup, MemberGroupCount int64, err error) {

	if member.Auth {

		if member.Permission {

			member.PermissionFlg = true

			return []tblmembergroup{}, 0, errors.New("unauthorized")
		}

		member.AuthFlg = true

		return []tblmembergroup{}, 0, errors.New("invalid Token")

	} else if member.Permission {

		member.PermissionFlg = true

		return []tblmembergroup{}, 0, errors.New("unauthorized")
	}

	var membergrplist []tblmembergroup

	MemberGroupList(membergrplist, limit, offset, filter, false, member.DBString)

	_, membercounts, _ := MemberGroupList(membergrplist, limit, offset, filter, false, member.DBString)

	membergrouplist, _, _ := MemberGroupList(membergrplist, limit, offset, filter, false, member.DBString)

	var membergrouplists []tblmembergroup

	for _, val := range membergrouplist {

		if !val.ModifiedOn.IsZero() {

			val.DateString = val.ModifiedOn.Format("02 Jan 2006 03:04 PM")

		} else {
			val.DateString = val.CreatedOn.Format("02 Jan 2006 03:04 PM")

		}

		membergrouplists = append(membergrouplists, val)

	}

	return membergrouplists, membercounts, nil

}
