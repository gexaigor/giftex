package model

// Role represents the role for this application
type Role string

const (
	// USER ...
	USER Role = "USER"
	// COMPANY ...
	COMPANY = "COMPANY"
	// ADMIN ...
	ADMIN = "ADMIN"
)

// String ...
func (r Role) String() string {
	roles := [...]string{"USER", "COMPANY", "ADMIN"}

	x := string(r)
	for _, v := range roles {
		if v == x {
			return x
		}
	}

	return ""
}
