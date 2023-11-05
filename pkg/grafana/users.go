package grafana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

type User struct {
	OrgID      int64    `json:"orgId,omitempty"`
	UserID     int64    `json:"userId,omitempty"`
	Email      string   `json:"email,omitempty"`
	Name       string   `json:"name,omitempty"`
	Login      string   `json:"login,omitempty"`
	IsAdmin    *bool    `json:"isadmin,omitempty"`
	Role       string   `json:"role,omitempty"`
	IsDisabled *bool    `json:"isDisabled,omitempty"`
	AuthLabels []string `json:"authLabels,omitempty"`
	TeamId     int64    `json:"teamId,omitempty"`
	// LastSeenAt    time.Time `json:"lastSeenAt,omitempty"`
	// LastSeenAtAge string    `json:"lastSeenAtAge,omitempty"`

}

func (c *Client) GetAllUsers() (err error) {
	users := make([]User, 0)
	body := bytes.NewReader([]byte{})
	err = c.request(MethodGet, "/api/org/users", body, &users)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error", err)
	}
	res, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println(string(res))
	return
}

func (c *Client) AddUser(loginOrEmail, role string) error {
	u := map[string]string{
		"loginOrEmail": loginOrEmail,
		"role":         role,
	}

	data, err := json.Marshal(u)

	err = c.request(MethodPost, "/api/org/users", bytes.NewReader(data), &u)
	if err != nil {
		fmt.Println("machalim")
	}
	res, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(res))
	return nil

}

func (c *Client) GetUserID(userID int64, err error) {}
