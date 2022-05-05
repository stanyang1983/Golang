package pojo

//定義使用者格式
type User struct {
	Id       int    `json:Id`
	Name     string `json:Name`
	Password string `json:Password`
	Email    string `json:Email`
}
