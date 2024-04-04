package member

// MemberSetup used initialize member configruation
func MemberSetup(config Config) *Member {

	return &Member{
		Auth:       config.Auth,
		Permission: config.Permission,
	}

}

