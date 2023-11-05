package grafana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Team struct {
	ID    int64  `json:"id,emitempty"`
	Name  string `json:"name,emitempty"`
	Email string `json:"email,emitempty"`
	OrgId int64  `json:"orgId,emitempty"`
}

func (c *Client) GetAllTeams() {
	teams := make([]Team, 0)
	c.request(http.MethodGet, "/api/teams", nil, &teams)
}

func (c *Client) GetTeamMembers(teamID string) {
	users := make([]User, 0)
	c.request(http.MethodGet, fmt.Sprintf("/api/teams/%s/members", teamID), nil, &users)

}

func (c *Client) AddTeam(name, email string, orgId int64) (*Team, error) {
	team := Team{
		Name:  name,
		Email: email,
		OrgId: orgId,
	}
	payload, err := json.Marshal(team)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = c.request(http.MethodPost, "/api/teams", bytes.NewBuffer(payload), &team)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &team, nil

}

func (c *Client) AddTeamMember(teamID string, userID int64) error {
	user := User{
		UserID: userID,
	}
	payload, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = c.request(http.MethodPost, fmt.Sprintf("/api/teams/%s/members", teamID), bytes.NewBuffer(payload), &user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *Client) searchTeamByName(name string) {
	//TODO
}

func (c *Client) searchTeamByID(id int64) {
	//TODO
}
