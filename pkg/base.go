package pkg


type User struct {
	//带有required的是必须传的参数，其余会自动用0填充，对于没有的参数会自动忽略
	Username string `form:"username" json:"username" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required"`
	Age int `form:"age" json:"age"`
}
