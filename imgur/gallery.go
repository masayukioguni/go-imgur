package imgur

import ()

// AccountService Digital Ocean API docs: https://developers.digitalocean.com/#account
type GalleryService struct {
	client *Client
}

type Image struct {
	ID   string `json:"id,omitempty"`
	Link string `json:"link,omitempty"`
}

type Gallery struct {
	Images []Image `json:"images,omitempty"`
}

type GalleryAlbumResponse struct {
	Success bool    `json:"success,omitempty"`
	Status  int     `json:"status,omitempty"`
	Data    []Image `json:"data,omitempty"`
}

func (s *GalleryService) GetAlbum() (*[]Image, error) {
	u := "/3/gallery"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	gar := new(GalleryAlbumResponse)
	_, err = s.client.Do(req, gar)
	if err != nil {
		return nil, err
	}

	return &gar.Data, err
}
