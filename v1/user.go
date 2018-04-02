package exmo

import (
	"net/url"
)

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
	req, err := a.c.newAuthenticatedRequest("user_info", url.Values{})

	if err != nil {
		return UserInfoStruct{}, nil
	}

	var v UserInfoStruct

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return UserInfoStruct{}, err
	}

	return v, nil
}
