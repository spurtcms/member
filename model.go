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

// soft delete check
func IsDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = 0")
}

// Member Group List
func MemberGroupList(membergroup []tblmembergroup, limit int, offset int, filter Filter, getactive bool, DB *gorm.DB) (membergroupl []tblmembergroup, TotalMemberGroup int64, err error) {

	query := DB.Model(TblMemberGroup{}).Scopes(IsDeleted).Order("id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(name)) ILIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")

	}

	if getactive {

		query = query.Where("is_active=1")

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&membergroup)

		return membergroup, 0, err

	}

	query.Find(&membergroup).Count(&TotalMemberGroup)

	return membergroup, TotalMemberGroup, err

}
