package member

// MemberSetup used initialize member configruation
func MemberSetup(config Config) *Member {

	MigrateTables(config.DB)

	return &Member{
		AuthEnable:       config.AuthEnable,
		Permissions:      config.Permissions,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
		DB:               config.DB,
	}

}
