package data

import "fmt"

func Fetch_Top_NewsCRUD(d *Strike_Meta_Request_Structure) *Response_wrapper_structure{
	var response Response_wrapper_structure
    
	fmt.Println(d)
	QText_value := "Hi " + d.Bybrisk_session_variables.Username + "! These are the top news for you."
	// Do db operations
	//_,_ = GetRoomDetailsFromDatabase(d.Bybrisk_session_variables.UserId)
	
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
						QContext: "news_id",
						QuestionDS: "string",
					},
					Answer: Answer_structure{
						ResponseType: "Date-Input",
						MultipleSelect: false,
						ResponseDS: "No DS Required",
					},
				},
			}, 
		},
	}

	return &response
}		