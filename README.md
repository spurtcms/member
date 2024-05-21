# Member Package

The 'Member' package grants website admin the authority to shape their online community through seamless member creation. Administrators hold the power to create members, transforming them into content consumers with exclusive access to tailored content. This streamlined process enables effortless audience management and curation, ensuring a personalized experience for each member within the website. 

## Features

- Members: Retrieve, create, update, and delete members.
- Record Validation: Check for existing email addresses, phone numbers, and names in member records.
- Interaction Features: Provide functionality for member delete popups and active member checks.
- Member Details: Retrieve details of specific members based on ID.



# Installation

``` bash
go get github.com/spurtcms/Member
```


# Usage Example

``` bash
import (
	"github.com/spurtcms/auth"
	"github.com/spurtcms/member"
)

func main() {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "SecretKey@123",
		DB: &gorm.DB{},
		RoleId: 1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Members Group", auth.CRUD)

	members := member.MemberSetup(member.Config{
		DB:               &gorm.DB{},
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	//membergroup
	if permisison {

		//list membergroup
		membergroup, count, err := members.ListMemberGroup(member.MemberGroupListReq{Limit: 10, Offset: 0})
		fmt.Println(membergroup, count, err)

		//create membergroup
		cerr := members.CreateMemberGroup(member.MemberGroupCreation{Name: "Default Group", Description: "default group", CreatedBy: 1})

		if cerr != nil {

			fmt.Println(cerr)
		}

		//update membergroup
		uerr := members.UpdateMemberGroup(member.MemberGroupCreationUpdation{
			Name:        "Default",
			Description: "default group2",
			ModifiedBy:  1,
			IsActive:    1,
		}, 2)

		if uerr != nil {

			fmt.Println(uerr)
		}

		// delete membergroup
		derr := members.DeleteMemberGroup(2, 2)

		if derr != nil {

			fmt.Println(derr)
		}
	}

	cpermisison, _ := Auth.IsGranted("Members", auth.CRUD)

	if cpermisison {

		//members list
		memberlist, count, err := members.ListMembers(10, 0, member.Filter{}, false)
		fmt.Println(memberlist, count, err)

		//create member
		memberdata, cerr := members.CreateMember(member.MemberCreationUpdation{
			FirstName: "tester",
			Username:  "Tester",
			Email:     "tester@gmail.com",
			Password:  "Tester@123",
		})

		fmt.Println(memberdata, cerr)

		//update member
		uerr := members.UpdateMember(member.MemberCreationUpdation{
			FirstName: "tester",
			Username:  "Tester",
			Email:     "tester@gmail.com",
			Password:  "Tester@123",
		}, 1)

		if uerr != nil {

			fmt.Println(uerr)
		}

		//delete member
		derr := members.DeleteMember(2, 1)

		if derr != nil {

			fmt.Println(derr)
		}

	}
}
```

# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/member/issues]. 
