package data

import (
	"fmt"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"net/http"
	"encoding/json"
	"github.com/strike-official/go-sdk/strike"
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

	strike := strike.Create("postera","")
	question := strike.Question("").QuestionCard(strike).SetHeaderToQuestion(strike,1,"HALF").AddTextRowToQuestion(strike,"h3",QText_value,"#424242",true)

	answer := question.Answer(strike,true).AnswerCardArray(strike,"VERTICAL")

	
	for _,v := range newsRespVar.Results{

		if v.ImageURL == "" {
			v.ImageURL = "https://s3.amazonaws.com/images.seroundtable.com/google-news-bot-gone-1314363466.jpg"
		}		

		answer = answer.AnswerCard(strike).
		SetHeaderToAnswer(strike, 1, "FULL").
		AddGraphicRowToAnswer(strike, "pic_row", v.ImageURL).
		AddTextRowToAnswer(strike, "h3", v.Title, "black",true).
		AddTextRowToAnswer(strike, "h5", v.SourceID, "#0065c9",false).
		AddTextRowToAnswer(strike, "h3", v.Description, "#595959",true).
		AddTextRowToAnswer(strike, "h6", v.Link, "blue",false)

		// v.ImageURL
		// v.SourceID
		// v.Link
	}
    
	response = Response_wrapper_structure{
		Status:200,
		Body: *strike,
	}

	return &response
}	