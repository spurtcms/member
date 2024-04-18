package member

import (
	"strings"
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
func (member *Member) ListMemberGroup(listreq MemberGroupListReq) (membergroup []tblmembergroup, MemberGroupCount int64, err error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return []tblmembergroup{}, 0, AuthErr
	}

	_, membercounts, _ := Membermodel.MemberGroupList(listreq, member.DB)

	membergrouplist, _, _ := Membermodel.MemberGroupList(listreq, member.DB)

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

/*Create Member Group*/
func (member *Member) CreateMemberGroup(membergrpc MemberGroupCreation) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return AuthErr
	}

	if membergrpc.Name == "" {

		return ErrorEmpty
	}

	var membergroup TblMemberGroup

	membergroup.Name = membergrpc.Name

	membergroup.Slug = strings.ToLower(membergrpc.Name)

	membergroup.Description = membergrpc.Description

	membergroup.CreatedBy = membergrpc.CreatedBy

	membergroup.IsActive = 1

	membergroup.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membermodel.MemberGroupCreate(&membergroup, member.DB)

	if err != nil {

		return err
	}

	return nil
}
