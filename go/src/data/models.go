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
    TeamName string `json:"teamName"`
    TeamLeader string `json:"teamLeader"`
    TeamMembers []string `json:"teamMembers"`
}