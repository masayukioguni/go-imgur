package imgur

import (
	"fmt"
	"net/http"
)

// AlbumService Digital Ocean API docs: https://api.imgur.com/models/album
type AlbumService struct {
	client *Client
}

type AlbumResponse struct {
	Success bool  `json:"success,omitempty"`
	Status  int   `json:"status,omitempty"`
	Data    Album `json:"data,omitempty"`
}

func (s *AlbumService) GetAlbum(id string) (*Album, *http.Response, error) {
	u := fmt.Sprintf("/3/album/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AlbumResponse)
	resp, err := s.client.Do(req, ar)
	if err != nil {
		return nil, resp, err
	}

	return &ar.Data, resp, err
}
