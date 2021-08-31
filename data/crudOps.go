package data

import (
	"fmt"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"net/http"
	"encoding/json"
)

func Fetch_Top_NewsCRUD(d *Strike_Meta_Request_Structure) *Response_wrapper_structure{
	var response Response_wrapper_structure
    
	fmt.Println(d)
	QText_value := "Hi " + d.Bybrisk_session_variables.Username + "! Here are the top news"
	// Do db operations
	//_,_ = GetRoomDetailsFromDatabase(d.Bybrisk_session_variables.UserId)
	resp, err := http.Get("https://newsdata.io/api/1/news?apikey=pub_59501c46f0eedf81631013a91f5fc9f83de&country=in&category=top")
	if err != nil {
    	log.Error("Get request error")
		log.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	type NewsResponse struct {
		Status       string `json:"status"`
		TotalResults int    `json:"totalResults"`
		Results      []struct {
			Title       string      `json:"title"`
			Link        string      `json:"link"`
			Keywords    string `json:"keywords"`
			Creator     []string    `json:"creator"`
			VideoURL    string `json:"video_url"`
			Description string      `json:"description"`
			Content     string `json:"content"`
			PubDate     string      `json:"pubDate"`
			ImageURL    string `json:"image_url"`
			SourceID    string      `json:"source_id"`
		} `json:"results"`
		NextPage int `json:"nextPage"`
	}

	var newsRespVar NewsResponse
	err = json.Unmarshal(body, &newsRespVar)

	var array_card_row_array [][]Card_Row_Object
	
	for _,v := range newsRespVar.Results{

		if v.ImageURL == "" {
			v.ImageURL = "https://s3.amazonaws.com/images.seroundtable.com/google-news-bot-gone-1314363466.jpg"
		}		

		card_row_array := []Card_Row_Object{
			Card_Row_Object{
				 Type: "header",
				 Descriptor: Descriptor_Structure{
					ContextObject: "",
					CardType: "FULL",
				 },	
			},
			Card_Row_Object{
				Type: "pic_row",
				Descriptor: Descriptor_Structure{
					Value: []string{v.ImageURL,},
				},
			},
			Card_Row_Object{
				Type: "h3",
				Descriptor: Descriptor_Structure{
					Value: []string{v.Title,},
					Color: "Black",
					Bold: true,
				},
			},
			Card_Row_Object{
				Type: "h5",
				Descriptor: Descriptor_Structure{
					Value: []string{v.SourceID,},
					Color: "#999999",
					Bold: false,
				},
			},
			Card_Row_Object{
				Type: "h3",
				Descriptor: Descriptor_Structure{
					Value: []string{v.Description,},
					Color: "Black",
					Bold: false,
				},
			},
			Card_Row_Object{
				Type: "h4",
				Descriptor: Descriptor_Structure{
					Value: []string{"",},
					Color: "Black",
					Bold: false,
				},
			},
			Card_Row_Object{
				Type: "h5",
				Descriptor: Descriptor_Structure{
					Value: []string{v.Link,},
					Color: "#3884ff",
					Bold: false,
				},	
			},
		}

		array_card_row_array = append(array_card_row_array,card_row_array)
	}

	card_row_array_ads := []Card_Row_Object{
		Card_Row_Object{
			 Type: "header",
			 Descriptor: Descriptor_Structure{
				ContextObject: "pic_row",
				CardType: "FULL",
			 },	
		},
		Card_Row_Object{
			Type: "pic_row",
			Descriptor: Descriptor_Structure{
				Value: []string{"https://resource.chemlinked.com.cn/food/articles/fQrPhd4PtyDwAE8yWuZq.jpg","https://sourceessay.com/essay/wp-content/uploads/2019/05/content-3948823_640-e1557552210608.jpg",},
			},
		},
	}
	array_card_row_array = append(array_card_row_array,card_row_array_ads)

	
	// Prepare response
	response = Response_wrapper_structure{
		Status: 200,
		Body: Body_structure{
			ActionHandler: "postera_news",
			NextActionHandler: "",
			QuestionArray: []Transaction_structure{
				Transaction_structure{
					Question: Question_structure{
						QuestionType: "Text",
						QText: QText_value,
						QContext: "",
						QuestionDS: "string",
					},
					Answer: Answer_structure{
						ResponseType: "Card",
						QCard: array_card_row_array,
						CardOrientation: "VERTICAL",
						MultipleSelect: false,
						ResponseDS: "No DS Required",
					},
				},
			}, 
		},
	}

	return &response
}		