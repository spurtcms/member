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

type MemberCreationUpdation struct {
	FirstName        string
	LastName         string
	Email            string
	MobileNo         string
	CreatedBy        int
	ModifiedBy       int
	IsActive         int
	ProfileImage     string
	ProfileImagePath string
	Username         string
	Password         string
	GroupId          int
}

type memberprofilecreationUpdation struct {
	MemberId        int
	ProfileId       int
	CompanyName     string
	CompanyLocation string
	CompanyLogo     string
	ProfileName     string
	ProfilePage     string
	About           string
	LinkedIn        string
	Website         string
	Twitter         string
	ClaimStatus     int
	ProfileSlug     string
	CreatedBy       int
	ModifiedBy      int
}

type MemberGroupCreationUpdation struct {
	Name        string
	Description string
	CreatedBy   int
	ModifiedBy  int
	IsActive    int
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

// Member list
func (membermodel MemberModel) MembersList(limit int, offset int, filter Filter, flag bool, DB *gorm.DB) (member []tblmember, Total_Member int64, err error) {

	query := DB.Model(TblMember{}).Select("tbl_members.id,tbl_members.uuid,tbl_members.member_group_id,tbl_members.first_name,tbl_members.last_name,tbl_members.email,tbl_members.mobile_no,tbl_members.profile_image,tbl_members.profile_image_path,tbl_members.created_on,tbl_members.created_by,tbl_members.modified_on,tbl_members.modified_by,tbl_members.is_active,tbl_members.is_deleted,tbl_members.deleted_on,tbl_members.deleted_by,tbl_member_groups.name as group_name").
		Joins("inner join tbl_member_groups on tbl_members.member_group_id = tbl_member_groups.id").Joins("inner join tbl_member_profiles on tbl_members.id = tbl_member_profiles.member_id").Where("tbl_members.is_deleted=?", 0).Order("id desc")

	if filter.Keyword != "" {

		query = query.Where("(LOWER(TRIM(tbl_members.first_name)) ILIKE LOWER(TRIM(?))"+" OR LOWER(TRIM(tbl_members.last_name)) ILIKE LOWER(TRIM(?))"+"OR LOWER(TRIM(tbl_member_profiles.company_name)) ILIKE LOWER(TRIM(?))"+" OR LOWER(TRIM(tbl_member_groups.name)) ILIKE LOWER(TRIM(?)))"+" AND tbl_members.is_deleted=0"+" AND tbl_member_groups.is_deleted=0", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

	}

	if filter.FirstName != "" {

		query = query.Debug().Where("LOWER(TRIM(tbl_members.first_name)) ILIKE LOWER(TRIM(?))"+" OR LOWER(TRIM(tbl_members.last_name)) ILIKE LOWER(TRIM(?))", "%"+filter.FirstName+"%", "%"+filter.FirstName+"%")

	}

	if flag {

		query.Find(&member)

		return member, 0, err

	}

	if limit != 0 && !flag {

		query.Offset(offset).Limit(limit).Order("id desc").Find(&member)

		return member, 0, err

	}
	query.Find(&member).Count(&Total_Member)

	return member, Total_Member, nil

}

// Member Insert
func (membermodel MemberModel) MemberCreate(member *tblmember, DB *gorm.DB) error {

	if err := DB.Model(TblMember{}).Create(&member).Error; err != nil {

		return err
	}

	return nil
}

func (membermodel MemberModel) UpdateMemberProfile(memberprof *TblMemberProfile, DB *gorm.DB) error {

	if err := DB.Model(TblMemberProfile{}).Create(&memberprof).Error; err != nil {

		return err
	}

	return nil
}

// Update Member
func (membermodel MemberModel) UpdateMember(member *tblmember, DB *gorm.DB) error {

	query := DB.Model(TblMember{}).Where("id=?", member.Id)

	if member.Password == "" && member.ProfileImage == "" && member.ProfileImagePath == "" {

		query.Omit("password , profile_image , profile_image_path").UpdateColumns(map[string]interface{}{"first_name": member.FirstName, "last_name": member.LastName, "member_group_id": member.MemberGroupId, "email": member.Email, "username": member.Username, "mobile_no": member.MobileNo, "is_active": member.IsActive, "modified_on": member.ModifiedOn, "modified_by": member.ModifiedBy})

	} else {

		query.UpdateColumns(map[string]interface{}{"first_name": member.FirstName, "last_name": member.LastName, "member_group_id": member.MemberGroupId, "email": member.Email, "username": member.Username, "mobile_no": member.MobileNo, "is_active": member.IsActive, "modified_on": member.ModifiedOn, "modified_by": member.ModifiedBy, "profile_image": member.ProfileImage, "profile_image_path": member.ProfileImagePath, "password": member.Password})
	}
	return nil
}

// Get Member group data
func (membermodel MemberModel) GetMemberProfileByMemberId(memberprof *TblMemberProfile, id int, DB *gorm.DB) (err error) {

	if err := DB.Model(TblMemberProfile{}).Where("member_id=?", id).First(&memberprof).Error; err != nil {

		return err
	}

	return nil
}

// update membercompanyprofile
func (membermodel MemberModel) MemberprofileUpdate(memberprof *TblMemberProfile, id int, DB *gorm.DB) error {

	query := DB.Model(TblMemberProfile{}).Where("id=?", id)

	if memberprof.CompanyLogo == "" {

		query.Omit("company_logo").UpdateColumns(map[string]interface{}{"member_id": memberprof.MemberId, "profile_name": memberprof.ProfileName, "profile_slug": memberprof.ProfileSlug, "company_name": memberprof.CompanyName, "company_location": memberprof.CompanyLocation, "about": memberprof.About, "seo_title": memberprof.SeoTitle, "seo_description": memberprof.SeoDescription, "seo_keyword": memberprof.SeoKeyword, "profile_page": memberprof.ProfilePage, "twitter": memberprof.Twitter, "linkedin": memberprof.Linkedin, "website": memberprof.Website, "claim_status": memberprof.ClaimStatus, "modified_by": memberprof.ModifiedBy, "modified_on": memberprof.ModifiedOn})

	} else {

		query.UpdateColumns(map[string]interface{}{"member_id": memberprof.MemberId, "profile_name": memberprof.ProfileName, "profile_slug": memberprof.ProfileSlug, "company_name": memberprof.CompanyName, "company_logo": memberprof.CompanyLogo, "company_location": memberprof.CompanyLocation, "about": memberprof.About, "seo_title": memberprof.SeoTitle, "seo_description": memberprof.SeoDescription, "seo_keyword": memberprof.SeoKeyword, "profile_page": memberprof.ProfilePage, "twitter": memberprof.Twitter, "linkedin": memberprof.Linkedin, "website": memberprof.Website, "claim_status": memberprof.ClaimStatus, "modified_by": memberprof.ModifiedBy, "modified_on": memberprof.ModifiedOn})

	}

	return nil
}

// Delete Member
func (membermodel MemberModel) DeleteMember(member *tblmember, id int, DB *gorm.DB) error {

	if err := DB.Model(TblMember{}).Where("id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": member.DeletedOn, "deleted_by": member.DeletedBy}).Error; err != nil {

		return err

	}
	return nil
}

// Check Email is already exists
func (membermodel MemberModel) CheckEmailInMember(member *TblMember, email string, userid int, DB *gorm.DB) error {

	if userid == 0 {
		if err := DB.Model(TblMember{}).Where("LOWER(TRIM(email))=LOWER(TRIM(?)) and is_deleted=0", email).First(&member).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Model(TblMember{}).Where("LOWER(TRIM(email))=LOWER(TRIM(?)) and id not in (?) and is_deleted = 0 ", email, userid).First(&member).Error; err != nil {

			return err
		}
	}

	return nil
}

func (membermodel MemberModel) CheckNumberInMember(member *TblMember, number string, userid int, DB *gorm.DB) error {

	if userid == 0 {

		if err := DB.Model(TblMember{}).Where("mobile_no = ? and is_deleted = 0", number).First(&member).Error; err != nil {

			return err
		}
	} else {

		if err := DB.Model(TblMember{}).Where("mobile_no = ? and id not in (?) and is_deleted=0", number, userid).First(&member).Error; err != nil {

			return err
		}
	}

	return nil
}

// Name already exists
func (membermodel MemberModel) CheckNameInMember(userid int, name string, DB *gorm.DB) (member tblmember, err error) {

	if userid == 0 {

		if err := DB.Model(TblMember{}).Where("LOWER(TRIM(username))=LOWER(TRIM(?)) and is_deleted=0", name).First(&member).Error; err != nil {

			return tblmember{}, err
		}
	} else {

		if err := DB.Model(TblMember{}).Where("LOWER(TRIM(username))=LOWER(TRIM(?)) and id not in (?) and is_deleted=0", name, userid).First(&member).Error; err != nil {

			return tblmember{}, err
		}
	}

	return member, nil
}

// Member Group Update
func (membermodel MemberModel) MemberGroupUpdate(membergroup *tblmembergroup, id int, DB *gorm.DB) error {

	if err := DB.Model(TblMemberGroup{}).Where("id=?", id).Updates(TblMemberGroup{Name: membergroup.Name, Slug: membergroup.Slug, Description: membergroup.Description, Id: membergroup.Id, ModifiedOn: membergroup.ModifiedOn, ModifiedBy: membergroup.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}

// Member Group Delete
func (membermodel MemberModel) DeleteMemberGroup(membergroup *tblmembergroup, id int, DB *gorm.DB) error {

	if err := DB.Debug().Model(TblMemberGroup{}).Where("id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": 1, "modified_by": membergroup.ModifiedBy}).Error; err != nil {

		return err

	}
	return nil
}
