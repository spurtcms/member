package member

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
)

// test listmembers function
func TestListMembers(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Members", auth.CRUD)

	member := MemberSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		memberlist,count,err := member.ListMembers(10,0,Filter{},false)

		if err != nil {

			panic(err)
		}

		fmt.Println(memberlist,count)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// test createmember function
func TestCreateMember(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Members", auth.CRUD)

	member := MemberSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		memberdata,err := member.CreateMember(MemberCreationUpdation{FirstName: "tester",Username: "Tester",Email:"tester@gmail.com",MobileNo:"9080706050",Password: "Tester@123"})

		if err != nil {

			panic(err)
		}

		fmt.Println(memberdata)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// test updatemember function
func TestUpdateMember(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Members", auth.CRUD)

	member := MemberSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		err := member.UpdateMember(MemberCreationUpdation{FirstName: "tester",Username: "Tester",Email:"tester@gmail.com",MobileNo:"9080706050",Password: "Tester@123"},1)

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// test updatemember function
func TestCreateMemberProfile(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Members", auth.CRUD)

	member := MemberSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		err := member.CreateMemberProfile(MemberprofilecreationUpdation{MemberId: 1,ProfileId: 5,CompanyName: "CMN",})

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}

}
// test updatememberprofile function
func TestUpdateMemberProfile(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Members", auth.CRUD)

	member := MemberSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		err := member.UpdateMemberProfile(MemberprofilecreationUpdation{ProfileId: 5,CompanyName: "CMNs",})

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// test deletemember function
func TestDeleteMember(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Members", auth.CRUD)

	member := MemberSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		err := member.DeleteMember(18,1)

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}

}