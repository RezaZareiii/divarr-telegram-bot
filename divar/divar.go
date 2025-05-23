package divar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const divarApiUrl = "https://api.divar.ir/v8/postlist/w/search"

func Search(page int) ([]PostRowData, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	requestBody := map[string]interface{}{
		"city_ids": []string{
			CityMap[tehran],
		},
		"pagination_data": map[string]interface{}{
			"@type":      "type.googleapis.com/post_list.PaginationData",
			"page":       page,
			"layer_page": page,
		},
		"source_view":            "FILTER",
		"disable_recommendation": false,
		"map_state": map[string]interface{}{
			"camera_info": map[string]interface{}{
				"bbox": map[string]interface{}{},
			},
		},
		"search_data": map[string]interface{}{
			"form_data": map[string]interface{}{
				"data": map[string]interface{}{
					"building-age": map[string]interface{}{
						"number_range": map[string]interface{}{
							"maximum": 15,
						},
					},
					"parking": map[string]interface{}{
						"boolean": map[string]interface{}{
							"value": true,
						},
					},
					"rooms": map[string]interface{}{
						"repeated_string": map[string]interface{}{
							"value": []string{"دو"},
						},
					},
					"credit": map[string]interface{}{
						"number_range": map[string]interface{}{
							"maximum": 600000000,
						},
					},
					"rent": map[string]interface{}{
						"number_range": map[string]interface{}{
							"maximum": 31000000,
						},
					},
					"size": map[string]interface{}{
						"number_range": map[string]interface{}{
							"minimum": 60,
						},
					},
					"districts": map[string]interface{}{
						"repeated_string": map[string]interface{}{
							"value": []string{
								districtsMap[saadatAbad],
								districtsMap[ponak],
								districtsMap[almahdi],
								districtsMap[janatAbadShomali],
								districtsMap[parvaz],
								districtsMap[farahzad],
								districtsMap[shahrakNafte],
								districtsMap[shahrakKohsar],
								districtsMap[shahrakGharb],
								districtsMap[shahin],
								districtsMap[ivanak],
								districtsMap[daria],
								districtsMap[sepahr],
								districtsMap[Asman],
								districtsMap[vanak],
								districtsMap[gandi],
								districtsMap[jordan],
								districtsMap[valenjak],
							},
						},
					},
					"category": map[string]interface{}{
						"str": map[string]interface{}{
							"value": "apartment-rent",
						},
					},
				},
			},
			"server_payload": map[string]interface{}{
				"@type": "type.googleapis.com/widgets.SearchData.ServerPayload",
				"additional_form_data": map[string]interface{}{
					"data": map[string]interface{}{
						"sort": map[string]interface{}{
							"str": map[string]interface{}{
								"value": "sort_date",
							},
						},
					},
				},
			},
		},
	}

	// Convert requestBody to io.Reader
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := client.Post(divarApiUrl, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", req.Status)
	}

	var response SearchResponse
	err = json.NewDecoder(req.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	posts := make([]PostRowData, 0, len(response.Posts))
	for _, post := range response.Posts {

		if post.WidgetType != "POST_ROW" {
			continue
		}

		posts = append(posts, post.PostRowData)
	}

	return posts, nil
}
