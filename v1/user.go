package exmo

type UserService struct {
	c *Client
}

type UserInfoStruct struct {
	UID        int64 `json:"uid,int"`
	Reserved   map[string]string
	Balances   map[string]string
	ServerDate int64 `json:"server_date,int"`
}

//Info UserService users info
func (a *UserService) Info() (UserInfoStruct, error) {
	return UserInfoStruct{}, nil
}
