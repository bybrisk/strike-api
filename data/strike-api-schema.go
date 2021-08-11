package data

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//post request for registering a user
type RegisterUserStructure struct{
	// UserID of the user
	//
	// required: true
	// max length: 1000
	UserID string `json: "userID" validate:"required"`

	// The full Name of the user
	//
	// required: true
	// max length: 1000
	UserName string `json: "userName" validate:"required"`

	//Phone number of the customer
	//
	// required: true
	// max length: 1000
	PhoneNumber string `json: "phoneNumer" validate:"required"`

	// Complete address of the user
	//
	// required: true
	// max length: 1000
	Address string `json: "address" validate:"required"`

	// Latitude of the customer
	//
	// required: true
	// max length: 1000
	Latitude float64 `json: "latitude" validate:"required"`

	// Longitude of the customer
	//
	// required: true
	// max length: 1000
	Longitude float64 `json: "longitude" validate:"required"`
}

//post request for registering a user
type RegisterUserToBusinessStruct struct{
	// UserID of the user 
	//
	// required: true
	// max length: 1000
	UserID string `json: "userID" validate:"required"`

	// BusinessID of the business user is subscribing to
	//
	// required: true
	// max length: 1000
	BusinessID string `json: "businessID" validate:"required"`
}

//post response
type RegisterPostSuccess struct {
	//userID of the user
	//
	UserID string `json:"userID"`
	//Message response
	//
	Message string `json:"message"`

	//status code
	//
	Status int64 `json:"status"`

	//Data of the user
	//
	Data IdOfDoc `json:"data"`
}

//post response
type RegisterToBusinessPostSuccess struct {
	//businessID of the business
	//
	BusinessID string `json:"businessID"`

	//Message response
	//
	Message string `json:"message"`

	//status code
	//
	Status int64 `json:"status"`

	//Business detail object
	//
	Detail SubscriptionStruct `json:"detail"`

}

type IdOfDoc struct{
	ID primitive.ObjectID `json:"-" bson:"_id"` 

	//Registeredlatitude of the user
	//
	Latitude float64 `json:"latitude"`
	
	//Registered longitude of the user
	//
	Longitude float64 `json:"longitude"`
	
	//registered phone number of the user
	//
	PhoneNumber string `json:"phonenumber"`
	
	//Registered name of the user
	//
	UserName string `json:"username"`
	
	//Registered address if the user
	//
	Address string `json:"address"`
	
	//Details of the subcription
	//
	Subscription []SubscriptionStruct `json:"subscription"`
}

type SubscriptionStruct struct {
	//BusinessID of the business
	//
	BusinessID string `json:"businessid"`
	
	//Name of the business
	//
	BusinessName string `json:"businessname"`
	
	//Category of the business
	//
	BusinessCategory string `json:"businesscategory"`

	//Url of the display profile of the business
	//
	Picurl string `json:"picurl"`

	//Address of the business
	//
	Address string `json:"address"`

	//Email of the business
	//
	Email string `json:"email"`

	//Latitude of the business
	//
	Latitude float64 `json:"latitude"`

	//Longitude of the business
	//
	Longitude float64 `json:"longitude"`

}

func (d *RegisterUserStructure) ValidateRegisterUserStructure() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *RegisterUserToBusinessStruct) ValidateRegisterUserToBusinessStruct() error {
	validate := validator.New()
	return validate.Struct(d)
}