package response

type XMLIndexData struct {
	Token        string `xml:"token"`
	PositionList string `xml:"positions_list"`
}

type HTMLIndexData struct {
	Token        string `json:"token"`
	PositionList string `json:"positions_list"`
}

type JSONIndexData struct {
	Token        string `json:"token"`
	PositionList string `json:"positions_list"`
}

type DocData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	URL   string `json:"url"`
}

type XMLDocData struct {
	Title string `xml:"Title"`
	Body  string `xml:"Body"`
	URL   string `xml:"URL"`
}
