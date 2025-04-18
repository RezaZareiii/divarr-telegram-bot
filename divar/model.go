package divar

const (
	saadatAbad       = "saadatabad"
	almahdi          = "almahdi"
	ponak            = "ponak"
	tehran           = "tehran"
	janatAbadShomali = "janatAbadShomali"
)

var CityMap = map[string]string{
	tehran: "1",
}

var districtsMap = map[string]string{
	saadatAbad:       "75",
	ponak:            "82",
	almahdi:          "143",
	janatAbadShomali: "145",
}

type SearchResponse struct {
	Posts     []PostRow `json:"list_widgets"`
	ActionLog ActionLog `json:"action_log"`
}

type ActionLog struct {
	ServerSideInfo ServerSideInfo `json:"server_side_info"`
}

type ActionLogInfo struct {
	Tokens     []string `json:"tokens"`
	SearchUUID string   `json:"search_uuid"`
}

type ServerSideInfo struct {
	ActionLogInfo ActionLogInfo `json:"info"`
}

type PostRow struct {
	WidgetType  string      `json:"widget_type"`
	PostRowData PostRowData `json:"data"`
}

type PostRowData struct {
	Title    string `json:"title"`
	Rent     string `json:"middle_description_text"`
	Credit   string `json:"top_description_text"`
	Location string `json:"bottom_description_text"`
	Token    string `json:"token"`
}
