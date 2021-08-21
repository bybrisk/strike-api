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
				Transaction_structure{
					Question: Question_structure{
						QuestionType: "qCard",
						QCard: []Card_Row_Object{
							Card_Row_Object{
								Type: "header",
								Descriptor: Descriptor_Structure{
									ContextObject: "h1",
									CardType: "FULL",
								},	
							},
							Card_Row_Object{
								Type: "pic_row",
								Descriptor: Descriptor_Structure{
									Value: []string{"https://www.nri.com/-/media/Corporate/jp/Images/common/h1/news/H1_news_top.jpg",},
								},
							},
							Card_Row_Object{
								Type: "h1",
								Descriptor: Descriptor_Structure{
									Value: []string{"This is the test text",},
									Color: "Black",
									Bold: true,
								},
							},
						},
						QContext: "news_id",
						QuestionDS: "string",
					},
					Answer: Answer_structure{
						ResponseType: "Text-Input",
						MultipleSelect: false,
						ResponseDS: "No DS Required",
					},
				},
			}, 
		},
	}

	return &response
}		