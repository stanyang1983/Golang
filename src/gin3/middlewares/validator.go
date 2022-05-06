package middlewares

import (
	"gin3/pojo"
	"regexp"

	"github.com/go-playground/validator/v10"
)

//限制密碼要大寫+其他組合
func UserPassword(field validator.FieldLevel) bool {
	if match, _ := regexp.MatchString(`^[A-Z]\w{4,10}$`, field.Field().String()); match {
		return true
	}
	return false
}

//目前看不懂...先用不到
func UserList(field validator.StructLevel) {
	users := field.Current().Interface().(pojo.Users)
	if users.UserListSize == len(users.UserList) {

	} else {
		field.ReportError(users.UserListSize, "Size of user list", "UserListSize", "UserListSizeMustEuqalsUserList", "")
	}
}
