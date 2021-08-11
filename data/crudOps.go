package data

func GetRoomDataCRUDOPS(d *Strike_Meta_Request_Structure) *Response_wrapper_structure{
	var response Response_wrapper_structure
    
	// Do db operations
	//_,_ = GetRoomDetailsFromDatabase(d.Bybrisk_session_variables.UserId)
	
	// Prepare response
	response = Response_wrapper_structure{
		Status: 200,
		Body: Body_structure{
			ActionHandler: "On/Off Devices",
			NextActionHandler: "",
			QuestionArray: []Transaction_structure{
				Transaction_structure{
					Question: Question_structure{
						QuestionType: "Text",
						QText: "This is your question",
						QContext: "app_id",
						QuestionDS: "string",
					},
					Answer: Answer_structure{
						ResponseType: "Date-Ibput",
						MultipleSelect: false,
						ResponseDS: "No DS Required",
					},
				},
			}, 
		},
	}

	return &response
}		