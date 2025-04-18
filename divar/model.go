package divar

/*
   "value": "922",
   "title": "پرواز",

    "value": "59",
    "title": "فرحزاد",


	"value": "921",
	"title": "شهرک نفت (منطقه ۵)",


	"value": "140",
	"title": "شهرک کوهسار",

	"value": "147",
    "title": "شاهین",


	"value": "927",
	"title": "ایوانک",


	"value": "78",
    "title": "شهرک غرب",



	"value": "929",
	"title": "دریا",


	"value": "926",
    "title": "سپهر",


	"value": "925",
    "title": "آسمان",

	"value": "315",
    "title": "ونک",

	"value": "96",
    "title": "گاندی",


	"value": "86",
    "title": "جردن",



*/

const (
	daria            = "daria"
	sepahr           = "sepahr"
	Asman            = "asman"
	vanak            = "vanak"
	gandi            = "gandi"
	jordan           = "jordan"
	shahrakGharb     = "shahrakGharb"
	saadatAbad       = "saadatabad"
	almahdi          = "almahdi"
	ponak            = "ponak"
	tehran           = "tehran"
	janatAbadShomali = "janatAbadShomali"
	parvaz           = "parvaz"
	farahzad         = "farahzad"
	shahrakNafte     = "shahrakNafte"
	shahrakKohsar    = "shahrakKohsar"
	shahin           = "shahin"
	ivanak           = "ivank"
	valenjak         = "valenjak"
)

var CityMap = map[string]string{
	tehran: "1",
}

var districtsMap = map[string]string{
	saadatAbad:       "75",
	ponak:            "82",
	almahdi:          "143",
	janatAbadShomali: "145",
	parvaz:           "922",
	farahzad:         "59",
	shahrakNafte:     "921",
	shahrakKohsar:    "140",
	shahrakGharb:     "78",
	shahin:           "147",
	ivanak:           "927",
	daria:            "929",
	sepahr:           "926",
	Asman:            "925",
	vanak:            "315",
	gandi:            "96",
	jordan:           "86",
	valenjak:         "55",
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
