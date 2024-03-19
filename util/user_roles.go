package util

// Supported user roles
const (
	ADMIN = "ADMIN"
	MOD   = "MOD"
	USER  = "USER"
)

// isSupportedRole returns true if the user role is supported
func IsSupportedRole(role string) bool {
	switch role {
	case ADMIN, MOD, USER:
		return true
	}

	return false
}
