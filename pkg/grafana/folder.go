package grafana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Folder struct {
	ID    int64  `json:"id"`
	UID   string `json:"uid"`
	Title string `json:"title"`
	URL   string `json:"URL"`
}

func (c *Client) GetAllFolders() ([]Folder, error) {
	var folders []Folder
	if err := c.request(http.MethodGet, "/api/folders/", nil, &folders); err != nil {
		return nil, err
	}
	return folders, nil
}

func (c *Client) CreateFolder(title string) (*Folder, error) {
	newFolder := &Folder{Title: title}
	newFolderPayload, err := json.Marshal(newFolder)
	if err != nil {
		return nil, err
	}
	err = c.request(http.MethodPost, "/api/folders", bytes.NewBuffer(newFolderPayload), &newFolder)
	if err != nil {
		return nil, err
	}

	fmt.Println(newFolder)
	return newFolder, nil
}

//func (c *Client) DeleteFolder() ([]Folder, error) {}
