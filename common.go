package member

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
	ErrorEmpty      = errors.New("given some values is empty")
)

func AuthandPermission(member *Member) error {

	//check auth enable if enabled, use auth pkg otherwise it will return error
	if member.AuthEnable && !member.Auth.AuthFlg {

		return ErrorAuth
	}
	//check permission enable if enabled, use team-role pkg otherwise it will return error
	if member.PermissionEnable && !member.Auth.PermissionFlg {

		return ErrorPermission

	}

	return nil
}

func hashingPassword(pass string) string {

	passbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {

		panic(err)

	}

	return string(passbyte)
}
