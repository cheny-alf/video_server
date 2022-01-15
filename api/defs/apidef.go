package defs

//requests
type UserCredential struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

//data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}
