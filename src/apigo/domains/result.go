package domains

import (
	"../utils"
)

type Result struct {
	User    *User
	Site    *Site
	Country *Country
	Error   *utils.Apierror
}
