package data

import (
	"github.com/go-playground/validator/v10"
)

//post request for registering a user
type Strike_Meta_Request_Structure struct{

	// Bybrisk variable from strike bot
	//
	bybrisk_session_variables Bybrisk_session_variables_struct `json: "bybrisk_session_variables"`

	// Our own variable from previous API
	//
	user_session_variables User_session_variables_struct `json: "user_session_variables"`
}

type Bybrisk_session_variables_struct struct{

	// User ID on Bybrisk
	//
	UserId string `json:"userId"`
	
	// Our own business Id in Bybrisk
	//
	BusinessId string `json:"businessId"`

	// Handler Name for the API chain
    //
	Handler string `json:"handler"`

	// Current location of the user
	//
	Location GeoLocation_struct `json:"location"`

	// Username of the user
	//
	Username string `json:"username"`

	// Address of the user
	//
	Address string `json:"address"`

	// Phone number of the user
	//
	Phone string `json:"phone"`
}

type GeoLocation_struct struct{
	// Latitude
	//
	Latitude float64 `json:"latitude"`

	// Longitude
	//
	Longitude float64 `json:"longitude"`
}

type User_session_variables_struct struct{
}

// Response structure
// Boiler plate response strucutre

type Response_wrapper_structure struct{
	
	// Internal status of the API
	//
	Status int64 `json:"status"`

	// Body structure
	//
	Body Body_structure `json:"body"`
}

type Body_structure struct{

	// Handler name of the API chain
	//
	ActionHandler string `json:"actionHandler"`

	// URI of the next API in the chain
	//
	NextActionHandler string `json:"nextActionHandler"`

	// Question Array structure
	//
	QuestionArray []Transaction_structure `json:"questionArray"`
}

type Transaction_structure struct{

	// Question object
	//
	Question Question_structure `json:"question"`

	// Answer Object
	//
	Answer Answer_structure `json:"answer"`
}

type Question_structure struct{

	// UI type of the question
	//
	QuestionType string `json:"questionType"`

	// This would depend on the questionType
	// So we use all the type for meta description
	//
	// This defines Text for the UI
	QText string `json:"qText"`

	//This defines Card object for the UI
	//
	QCard Card_structure `json:"qCard"`

	// Context of the question
	// The value will be binded in this key
	//
	QContext string `json:"qContext"`

	// Discription of the data strucutre of the question object for strike UI engine
	//
	QuestionDS string `json:"questionDS"`
}

type Answer_structure struct{

	// UI type of the answer
	//
	ResponseType string `json:"responseType"`

	// This would depend on the responseType
	// So we use all the type for meta description
	//
	// This defines Text for the UI
	//QText string `json:"qText"`

	//This defines Card object for the UI
	//
	//QCard Card_structure `json:"qCard"`

	// Set this to true if we want multiple values could be selected by the user for this particular question
	//
	MultipleSelect bool `json:"multipleSelect"`

	// Discription of the data strucutre of the question object for strike UI engine
	//
	ResponseDS string `json:"responseDS"`
}

// UI specific structures

type Card_structure struct{
	faang string `json:"john"`
}

func (d *Strike_Meta_Request_Structure) ValidateStrike_Meta_Request_Structure() error {
	validate := validator.New()
	return validate.Struct(d)
}