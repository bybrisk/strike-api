
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/strike-api/data"
)

// swagger:route POST /strike/postera/top postera get
// Fetch top news.
//
// responses:
//	200: roomPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Strike_Register) Fetch_Top_News (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> strike-api Module postera top_news")
	request := &data.Strike_Meta_Request_Structure{}
	err:=request.FromJSONToStrike_Meta_Request_Structure(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = request.ValidateStrike_Meta_Request_Structure()
	if err!=nil {
		p.l.Println("Validation error in POST request -> strike-api Module postera top_news \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//get data from db
	response := data.Fetch_Top_NewsCRUD(request)

	//writing to the io.Writer
	w.Header().Set("Content-Type", "application/json")
	
	err = response.Response_wrapper_structureToJSON(w)
	fmt.Println(response)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}

