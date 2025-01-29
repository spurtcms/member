package member

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"github.com/spurtcms/member/migration"
)

// MemberSetup used initialize member configruation
func MemberSetup(config Config) *Member {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Member{
		AuthEnable:       config.AuthEnable,
		Permissions:      config.Permissions,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
		DB:               config.DB,
	}

}

// list member
func (member *Member) ListMembers(offset int, limit int, filter Filter, flag bool, TenantId int) (memb []Tblmember, totoalmember int64, err error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return []Tblmember{}, 0, AuthErr
	}

	Membermodel.Userid = member.UserId
	Membermodel.DataAccess = member.DataAccess

	memberlist, _, _ := Membermodel.MembersList(limit, offset, filter, flag, member.DB, TenantId)

	_, Total_users, _ := Membermodel.MembersList(0, 0, filter, flag, member.DB, TenantId)

	var memberlists []Tblmember

	for _, val := range memberlist {

		var first = val.FirstName
		var last = val.LastName
		var firstn = strings.ToUpper(first[:1])
		var lastn string
		if val.LastName != "" {
			lastn = strings.ToUpper(last[:1])
		}
		var Name = firstn + lastn
		val.NameString = Name
		val.CreatedDate = val.CreatedOn.Format("02 Jan 2006 03:04 PM")
		if !val.ModifiedOn.IsZero() {
			val.ModifiedDate = val.ModifiedOn.Format("02 Jan 2006 03:04 PM")
		} else {
			val.ModifiedDate = val.CreatedOn.Format("02 Jan 2006 03:04 PM")
		}

		memberlists = append(memberlists, val)

	}

	return memberlists, Total_users, nil

}

// Create Member
func (member *Member) CreateMember(Mc MemberCreationUpdation) (Tblmember, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return Tblmember{}, AuthErr
	}
	uvuid := (uuid.New()).String()

	var cmember Tblmember
	cmember.Uuid = uvuid
	cmember.ProfileImage = Mc.ProfileImage
	cmember.ProfileImagePath = Mc.ProfileImagePath
	cmember.MemberGroupId = Mc.GroupId
	cmember.FirstName = Mc.FirstName
	cmember.LastName = Mc.LastName
	cmember.Email = Mc.Email
	cmember.MobileNo = Mc.MobileNo
	cmember.IsActive = Mc.IsActive
	cmember.Username = Mc.Username
	if Mc.Password != "" {
		hash_pass := hashingPassword(Mc.Password)
		cmember.Password = hash_pass
	}
	cmember.CreatedBy = Mc.CreatedBy
	cmember.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	cmember.StorageType = Mc.StorageType
	cmember.TenantId = Mc.TenantId
	err := Membermodel.MemberCreate(&cmember, member.DB)
	if err != nil {

		return Tblmember{}, err
	}

	return cmember, nil

}

// Update Member
func (member *Member) UpdateMember(Mc MemberCreationUpdation, id int, tenantid int) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return AuthErr
	}

	uvuid := (uuid.New()).String()

	var umember Tblmember
	umember.Uuid = uvuid
	umember.Id = id
	umember.MemberGroupId = Mc.GroupId
	umember.FirstName = Mc.FirstName
	umember.LastName = Mc.LastName
	umember.Email = Mc.Email
	umember.MobileNo = Mc.MobileNo
	umember.ProfileImage = Mc.ProfileImage
	umember.ProfileImagePath = Mc.ProfileImagePath
	umember.IsActive = Mc.IsActive
	umember.ModifiedBy = Mc.ModifiedBy
	umember.Username = Mc.Username
	password := Mc.Password
	if password != "" {
		hash_pass := hashingPassword(password)
		umember.Password = hash_pass
	}
	umember.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	umember.StorageType = Mc.StorageType
	err := Membermodel.UpdateMember(&umember, member.DB, tenantid)
	if err != nil {

		return err
	}

	return nil
}

// create member profile
func (member *Member) CreateMemberProfile(Mc MemberprofilecreationUpdation) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return AuthErr
	}

	var memberprof TblMemberProfile
	memberprof.MemberId = Mc.MemberId
	memberprof.CompanyName = Mc.CompanyName
	memberprof.CompanyLocation = Mc.CompanyLocation
	memberprof.CompanyLogo = Mc.CompanyLogo
	memberprof.ProfileName = Mc.ProfileName
	memberprof.ProfileSlug = Mc.ProfileSlug
	memberprof.About = Mc.About
	memberprof.Linkedin = Mc.LinkedIn
	memberprof.Twitter = Mc.Twitter
	memberprof.Website = Mc.Website
	memberprof.ClaimStatus = Mc.ClaimStatus
	memberprof.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	memberprof.CreatedBy = Mc.ModifiedBy
	memberprof.StorageType = Mc.StorageType
	memberprof.TenantId = Mc.TenantId
	err2 := Membermodel.CreateMemberProfile(&memberprof, member.DB)
	if err2 != nil {

		return err2
	}

	return nil
}

// update memberprofile
func (member *Member) UpdateMemberProfile(Mc MemberprofilecreationUpdation, tenantid int) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return AuthErr
	}

	var memberprof TblMemberProfile

	memberprof.MemberId = Mc.MemberId
	memberprof.Id = Mc.ProfileId
	memberprof.CompanyName = Mc.CompanyName
	memberprof.CompanyLocation = Mc.CompanyLocation
	memberprof.CompanyLogo = Mc.CompanyLogo
	memberprof.ProfileName = Mc.ProfileName
	memberprof.ProfileSlug = Mc.ProfileSlug
	memberprof.SeoTitle = Mc.SeoTitle
	memberprof.SeoDescription = Mc.SeoDescription
	memberprof.SeoKeyword = Mc.SeoKeyword
	memberprof.About = Mc.About
	memberprof.Linkedin = Mc.LinkedIn
	memberprof.Twitter = Mc.Twitter
	memberprof.Website = Mc.Website
	memberprof.ClaimStatus = Mc.ClaimStatus
	memberprof.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	memberprof.ModifiedBy = Mc.ModifiedBy
	memberprof.StorageType = Mc.StorageType
	memberprof.TenantId = tenantid
	err2 := Membermodel.MemberprofileUpdate(&memberprof, Mc.ProfileId, member.DB, tenantid)

	if err2 != nil {

		return err2
	}

	return nil
}

// delete member
func (member *Member) DeleteMember(id int, modifiedBy int, tenantid int) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return AuthErr
	}

	var dmember Tblmember

	dmember.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	dmember.DeletedBy = modifiedBy

	err := Membermodel.DeleteMember(&dmember, id, member.DB, tenantid)

	Membermodel.DeleteMemberProfile(id, modifiedBy, dmember.DeletedOn, member.DB, tenantid)

	if err != nil {

		return err
	}

	return nil
}

// Get member data
func (member *Member) GetMemberDetails(id int, tenantid int) (members Tblmember, err error) {

	var memberdata Tblmember

	err = Membermodel.MemberDetails(&memberdata, id, member.DB, tenantid)
	if err != nil {

		return Tblmember{}, err
	}

	return memberdata, nil

}

// Get memberprofile data
func (member *Member) GetMemberProfileByMemberId(memberid int, tenantid int) (memberprofs TblMemberProfile, err error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return TblMemberProfile{}, AuthErr
	}

	var memberprof TblMemberProfile
	err1 := Membermodel.GetMemberProfileByMemberId(&memberprof, memberid, member.DB, tenantid)
	if err1 != nil {

		return TblMemberProfile{}, err1
	}

	return memberprof, nil

}

// Check Number is already exits or not
func (member *Member) CheckProfileSlugInMember(id int, number string, tenantid int) (bool, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return false, AuthErr
	}

	var memberprof TblMemberProfile
	err := Membermodel.CheckProfileSlugInMember(&memberprof, number, id, member.DB, tenantid)
	if err != nil {
		return false, err
	}

	return true, nil
}

// member is_active
func (member *Member) MemberStatus(memberid int, status int, modifiedby int, tenantid int) (bool, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return false, AuthErr
	}

	var memberstatus TblMember
	memberstatus.ModifiedBy = modifiedby
	memberstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membermodel.MemberStatus(memberstatus, memberid, status, member.DB, tenantid)
	if err != nil {
		return false, err
	}

	return true, nil

}

// multiselecte member delete
func (member *Member) MultiSelectedMemberDelete(Memberid []int, modifiedby int, tenantid int) (bool, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return false, AuthErr
	}

	var members TblMember
	members.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	members.DeletedBy = modifiedby
	members.IsDeleted = 1

	err := Membermodel.MultiSelectedMemberDelete(&members, Memberid, member.DB, tenantid)
	if err != nil {

		return false, err
	}

	return true, nil
}

// multiselecte member status change
func (member *Member) MultiSelectMembersStatus(memberid []int, status int, modifiedby int, tenantid int) (bool, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return false, AuthErr
	}

	var memberStatus TblMember
	memberStatus.ModifiedBy = modifiedby
	memberStatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membermodel.MultiMemberIsActive(&memberStatus, memberid, status, member.DB, tenantid)
	if err != nil {

		return false, err
	}

	return true, nil

}

func (member *Member) CheckProfileSlug(profileSlug string, profileID int, tenantid int) (TblMemberProfile, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return TblMemberProfile{}, AuthErr
	}

	profile, err := Membermodel.CheckProfileSlug(profileSlug, member.DB, tenantid)

	if err != nil {
		return TblMemberProfile{}, err
	}

	if profile.Id != 0 && profile.Id == profileID {
		return profile, nil
	} else if profile.Id == 0 {
		return profile, nil
	}

	return TblMemberProfile{}, nil
}

func (member *Member) GetMemberAndProfileData(memberId int, emailid string, profileId int, profileSlug string, tenantid int) (Tblmember, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return Tblmember{}, AuthErr
	}

	profile, err := Membermodel.GetMemberProfile(memberId, emailid, profileId, profileSlug, member.DB, tenantid)
	if err != nil {
		return Tblmember{}, err
	}

	return profile, nil
}

func (member *Member) DashboardMemberCount(tenantid int) (totalcount int, lasttendayscount int, err error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return 0, 0, AuthErr
	}

	allmembercount, err := Membermodel.AllMemberCount(member.DB, tenantid)
	if err != nil {
		return 0, 0, err
	}

	Lmembercount, err := Membermodel.NewmemberCount(member.DB, tenantid)
	if err != nil {
		return 0, 0, err
	}

	return int(allmembercount), int(Lmembercount), nil
}

// Active MemberList Function//
func (member *Member) ActiveMemberList(limit int, tenantid int) (memberdata []Tblmember, err error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return []Tblmember{}, AuthErr
	}

	var members []Tblmember
	activememlist, err := Membermodel.ActiveMemberList(members, limit, member.DB, tenantid)

	var memberlist []Tblmember
	for _, val := range activememlist {
		val.DateString = val.LoginTime.Format("02 Jan 2006 03:04 PM")
		memberlist = append(memberlist, val)
	}

	if err != nil {
		return []Tblmember{}, err
	}

	return memberlist, nil

}

// Member flexible update functionality
func (member *Member) MemberFlexibleUpdate(memberData map[string]interface{}, memberId, modifiedBy int, tenantid int) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return AuthErr
	}

	currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	memberData["modified_on"] = currentTime
	memberData["modified_by"] = modifiedBy
	err := Membermodel.FlexibleMemberUpdate(memberData, memberId, member.DB, tenantid)
	if err != nil {
		return err
	}

	return nil

}

// Memeber profile flexible update
func (member *Member) MemberProfileFlexibleUpdate(memberProfileData map[string]interface{}, memberId, modifiedBy int, tenantid int) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return AuthErr
	}

	currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	memberProfileData["modified_on"] = currentTime
	memberProfileData["modified_by"] = modifiedBy
	err := Membermodel.FlexibleMemberProfileUpdate(memberProfileData, memberId, member.DB, tenantid)
	if err != nil {
		return err
	}

	return nil
}

// Member password update functionality
func (member *Member) MemberPasswordUpdate(newPassword, confirmPassword, oldPassword string, memberId, modifiedBy int, tenantid int) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return AuthErr
	}

	var memberData TblMember
	if err := Membermodel.GetMemberDetailsByMemberId(&memberData, memberId, member.DB, tenantid); err != nil {
		return err
	}
	if oldPassword != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(memberData.Password), []byte(oldPassword)); err != nil {
			return err
		}
	}
	if newPassword != confirmPassword {
		return ErrorPassMissMatch
	}

	hash_pass := hashingPassword(confirmPassword)
	if oldPassword != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(hash_pass), []byte(oldPassword)); err != nil {
			return err
		}
	}
	memberData.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	memberData.ModifiedBy = modifiedBy
	memberData.Password = hash_pass
	err := Membermodel.MemberPasswordUpdate(memberData, memberId, member.DB, tenantid)
	if err != nil {
		return err
	}

	return nil

}

// Get member settings
func (member *Member) GetMemberSettings(tenantid int) (TblMemberSetting, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return TblMemberSetting{}, AuthErr
	}

	membersetttings, err := Membermodel.GetMemberSettings(member.DB, tenantid)

	if err != nil {
		return TblMemberSetting{}, err
	}

	return membersetttings, nil
}

// set member settings
func (member *Member) SetMemberSettings(membersett MemberSettings, tenantid int) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return AuthErr
	}

	var updatedetails = make(map[string]interface{})
	updatedetails["member_login"] = membersett.MemberLogin
	updatedetails["allow_registration"] = membersett.AllowRegistration
	updatedetails["notification_users"] = membersett.NotificationUsers
	updatedetails["modified_by"] = membersett.ModifiedBy
	updatedetails["modified_on"], _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if err := Membermodel.UpdateMemberSetting(updatedetails, member.DB, tenantid); err != nil {
		return err
	}

	return nil
}
