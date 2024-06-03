package member

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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
func (member *Member) ListMembers(offset int, limit int, filter Filter, flag bool) (memb []Tblmember, totoalmember int64, err error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return []Tblmember{}, 0, AuthErr
	}

	memberlist, _, _ := Membermodel.MembersList(limit, offset, filter, flag, member.DB)

	_, Total_users, _ := Membermodel.MembersList(0, 0, filter, flag, member.DB)

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

	err := Membermodel.MemberCreate(&cmember, member.DB)

	if err != nil {

		return Tblmember{}, err
	}

	return cmember, nil

}

// Update Member
func (member *Member) UpdateMember(Mc MemberCreationUpdation, id int) error {

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

	err := Membermodel.UpdateMember(&umember, member.DB)

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

	err2 := Membermodel.CreateMemberProfile(&memberprof,member.DB)

	if err2 != nil {

		return err2
	}

	return nil
}


// update memberprofile
func (member *Member) UpdateMemberProfile(Mc MemberprofilecreationUpdation) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return AuthErr
	}

	var memberprof TblMemberProfile

	memberprof.MemberId = Mc.ProfileId

	memberprof.Id = Mc.ProfileId

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

	memberprof.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	memberprof.ModifiedBy = Mc.ModifiedBy

	err2 := Membermodel.MemberprofileUpdate(&memberprof, Mc.ProfileId, member.DB)

	if err2 != nil {

		return err2
	}

	return nil
}

// delete member
func (member *Member) DeleteMember(id int, modifiedBy int) error {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return AuthErr
	}

	var dmember Tblmember

	dmember.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	dmember.DeletedBy = modifiedBy

	err := Membermodel.DeleteMember(&dmember, id, member.DB)

	if err != nil {

		return err
	}

	return nil
}

// member token
func (member *Member) GenerateMemberToken(memberid int, secretKey string) (token string, err error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return token, AuthErr
	}

	var MemberDetails TblMember

	if err := Membermodel.GetMemberDetailsByMemberId(&MemberDetails, memberid, member.DB); err != nil {

		return "", err
	}

	token, tokenerr := CreateMemberToken(MemberDetails.Id, MemberDetails.MemberGroupId, secretKey)

	if tokenerr != nil {

		return "", err
	}

	return token, nil
}

/*Create meber token*/
func CreateMemberToken(userid, roleid int, secretkey string) (string, error) {

	atClaims := jwt.MapClaims{}

	atClaims["member_id"] = userid

	atClaims["group_id"] = roleid

	atClaims["expiry_time"] = time.Now().Add(2 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return token.SignedString([]byte(secretkey))
}

// Get member data
func (member *Member) GetMemberDetails(id int) (members Tblmember, err error) {

	var memberdata Tblmember

	err = Membermodel.MemberDetails(&memberdata, id, member.DB)

	if err != nil {

		return Tblmember{}, err
	}

	return memberdata, nil

}

// Get memberprofile data
func (member *Member) GetMemberProfileByMemberId(memberid int) (memberprofs TblMemberProfile, err error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return TblMemberProfile{}, AuthErr
	}

	var memberprof TblMemberProfile

	err1 := Membermodel.GetMemberProfileByMemberId(&memberprof, memberid, member.DB)

	if err1 != nil {

		return TblMemberProfile{}, err1
	}

	return memberprof, nil

}

// Check Number is already exits or not
func (member *Member) CheckProfileSlugInMember(id int, number string) (bool, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return false, AuthErr
	}

	var memberprof TblMemberProfile

	err := Membermodel.CheckProfileSlugInMember(&memberprof, number, id, member.DB)

	if err != nil {
		return false, err
	}

	return true, nil
}

// member is_active
func (member *Member) MemberStatus(memberid int, status int, modifiedby int) (bool, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return false, AuthErr
	}

	var memberstatus TblMember
	memberstatus.ModifiedBy = modifiedby
	memberstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membermodel.MemberStatus(memberstatus, memberid, status, member.DB)

	if err != nil {
		return false, err
	}

	return true, nil

}

// multiselecte member delete
func (member *Member) MultiSelectedMemberDelete(Memberid []int, modifiedby int) (bool, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {
		return false, AuthErr
	}

	var members TblMember
	members.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	members.DeletedBy = modifiedby
	members.IsDeleted = 1

	err := Membermodel.MultiSelectedMemberDelete(&members, Memberid, member.DB)

	if err != nil {

		return false, err
	}

	return true, nil
}

// multiselecte member status change
func (member *Member) MultiSelectMembersStatus(memberid []int, status int, modifiedby int) (bool, error) {

	if AuthErr := AuthandPermission(member); AuthErr != nil {

		return false, AuthErr
	}

	var memberstatus TblMember

	memberstatus.ModifiedBy = modifiedby

	memberstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Membermodel.MultiMemberIsActive(&memberstatus, memberid, status, member.DB)

	if err != nil {

		return false, err
	}

	return true, nil

}
