package model

type User struct {
	Username string `json:"username" query:"username" vd:"$!nil"`
	Password string `json:"password" query:"password" vd:"$!nil && mblen($)>=6"`
}
