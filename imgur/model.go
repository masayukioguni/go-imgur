package imgur

type Image struct {
	ID   string `json:"id,omitempty"`
	Link string `json:"link,omitempty"`
}

type Album struct {
	ID     string  `json:"id,omitempty"`
	Title  string  `json:"title,omitempty"`
	Link   string  `json:"link,omitempty"`
	Images []Image `json:"images,omitempty"`
}

type Gallery struct {
	Images []Image `json:"images,omitempty"`
}

type Tag struct {
	Name       string  `json:"name,omitempty"`
	Followers  int     `json:"followers,omitempty"`
	TotalItems int     `json:"total_items,omitempty"`
	Images     []Image `json:"items,omitempty"`
}
