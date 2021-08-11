package data_test

import (
	"testing"
	"fmt"
	"github.com/bybrisk/strike-api/data"
)



func TestGetRoomDataCRUDOPS(t *testing.T) {

	register := &data.Strike_Meta_Request_Structure{
		bybrisk_session_variables: data.Bybrisk_session_variables_struct{
			UserId: "abc",
			BusinessId:"xyz",
			Handler: "action",
			Location: data.GeoLocation_struct{
				Latitude:23.45343545,
				Longitude:77.454524,
			},
			Username:"viky",
			Address: "Mumbai",
			Phone:"9340212323",
		},
		user_session_variables: data.User_session_variables_struct{},
	}

	res:= data.GetRoomDataCRUDOPS(register) 

	fmt.Println(res)
}

/*func TestGetUserIDCRUDOPS(t *testing.T){
	res:= data.GetUserIDCRUDOPS("9340232345")
	fmt.Println(res)
}*/

/*func TestRegisterUserToBusinessCRUDOPS(t *testing.T) {
	payload := &data.RegisterUserToBusinessStruct{
		UserID: "6083deb86fcd474489784fee",
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
	}

	res := data.RegisterUserToBusinessCRUDOPS(payload)
	fmt.Println(res)
}*/