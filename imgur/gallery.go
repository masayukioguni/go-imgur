package imgur

import (
	"net/http"
)

// GalleryService Imgur API
type GalleryService struct {
	client *Client
}

type GalleryAlbumResponse struct {
	Success bool    `json:"success,omitempty"`
	Status  int     `json:"status,omitempty"`
	Data    []Image `json:"data,omitempty"`
}

func (s *GalleryService) GetAlbum() ([]Image, *http.Response, error) {
	u := "/3/gallery"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gar := new(GalleryAlbumResponse)
	resp, err := s.client.Do(req, gar)
	if err != nil {
		return nil, resp, err
	}

	return gar.Data, resp, err
}
