package data

/*
{
    "username": "",
    "password": "",
    "accessLevel": ""
}
*/
type User struct {
	Username   string	`json:"username"`
	Password   string	`json:"password"`
	AccessLevel string	`json:"accessLevel"`
}
/*
{
    "teamName": "",
    "teamLeader": "",
    "teamMembers": []
}
*/
type Team struct {
    Teamname string `json:"teamname"`
    Teamleader string `json:"teamleader"`
    TeamMembers []User `json:"teamMembers"`
}