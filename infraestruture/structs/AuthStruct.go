package structs

// User demo
type User struct {
	Name      string
	UserEmail string
	Active    bool
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
