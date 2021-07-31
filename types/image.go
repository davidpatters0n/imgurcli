package types

type Image struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageType   string `json:"type"`
	Animated    bool   `json:"animated"`
	Size        int    `json:"size"`
	Views       int    `json:"views"`
	Link        string `json:"link"`
	Favorite    bool   `json:"favorite"`
	InGallery   bool   `json:"in_gallery"`
}

type ImageData struct {
	Image `json:"data"`
}
