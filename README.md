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
		membergroup, count, err := member.ListMemberGroup(MemberGroupListReq{Limit: 10, Offset: 0}, 1)

		if err != nil {

			panic(err)
		}

		//create membergroup
		err := member.CreateMemberGroup(MemberGroupCreation{"sports", "indian team", 1}, 1)

		if err != nil {

			panic(err)
		}

		//update membergroup
		err := member.UpdateMemberGroup(MemberGroupCreationUpdation{Name:"Default",Description: "default group2",ModifiedBy: 1,IsActive: 1},2,1)

		if err != nil {

			panic(err)
		}

		// delete membergroup
			err := member.DeleteMemberGroup(2,1,1)

		if err != nil {

			panic(err)
		}
	}

	cpermisison, _ := Auth.IsGranted("Members", auth.CRUD)

	if cpermisison {

		//members list
		memberlist, count, err := member.ListMembers(0, 10, Filter{}, false,1)
		fmt.Println(memberlist, count, err)

		//create member
			memberdata, err := member.CreateMember(MemberCreationUpdation{
			  FirstName: "tester",
			  Username: "Tester",
			  Email: "tester@gmail.com",
			  MobileNo: "9080706050",
			  Password: "Tester@123",
			  TenantId: 1
			})

		fmt.Println(memberdata, cerr)

		//update member
				err := member.UpdateMember(MemberCreationUpdation{
				FirstName: "testers",
				Username: "Testers",
				Email: "testers@gmail.com",
				MobileNo: "9080706050",
				Password: "Testers@123"},1, 1)


		if err != nil {

			fmt.Println(err)
		}


		// create member profile

		err := member.CreateMemberProfile(MemberprofilecreationUpdation{
		         MemberId: 1,
			     ProfileId: 5,
				 CompanyName: "CMN",
				 TenantId: 1
				})

		if err != nil {

			panic(err)
		}

         // update member profile

		 err := member.UpdateMemberProfile(MemberprofilecreationUpdation{
		         ProfileId: 1,
				 CompanyName: "CMNs"
			   }, 1)

		if err != nil {

			panic(err)
		}

		//delete member
		err := member.DeleteMember(1, 1, 1)

		if err != nil {

			fmt.Println(err)
		}

	}
}
```

# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/member/issues].
