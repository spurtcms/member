package member

import "gorm.io/gorm"

type Filter struct {
	Keyword   string
	Category  string
	Status    string
	FromDate  string
	ToDate    string
	FirstName string
}

type MemberGroupListReq struct {
	Limit            int
	Offset           int
	Keyword          string
	Category         string
	Status           string
	FromDate         string
	ToDate           string
	FirstName        string
	ActiveGroupsOnly bool
}

type MemberGroupCreation struct {
	Name        string
	Description string
	CreatedBy   int
}

// soft delete check
func IsDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = 0")
}

type MemberModel struct{}

var Membermodel MemberModel

// Member Group List
func (membermodel MemberModel) MemberGroupList(listre MemberGroupListReq, DB *gorm.DB) (membergroup []tblmembergroup, TotalMemberGroup int64, err error) {

	query := DB.Model(TblMemberGroup{}).Scopes(IsDeleted).Order("id desc")

	if listre.Keyword != "" {

		query = query.Where("LOWER(TRIM(name)) ILIKE LOWER(TRIM(?))", "%"+listre.Keyword+"%")

	}

	if listre.ActiveGroupsOnly {

		query = query.Where("is_active=1")

	}

	if listre.Limit != 0 {

		query.Limit(listre.Limit).Offset(listre.Offset).Find(&membergroup)

		return membergroup, 0, err

	}

	query.Find(&membergroup).Count(&TotalMemberGroup)

	return membergroup, TotalMemberGroup, err

}

// Member Group Insert
func (membermodel MemberModel) MemberGroupCreate(membergroup *TblMemberGroup, DB *gorm.DB) error {

	if err := DB.Model(TblMemberGroup{}).Create(&membergroup).Error; err != nil {

		return err
	}

	return nil
}
