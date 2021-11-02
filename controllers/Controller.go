package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/divya2703/covid-tracker-rest-api/cache"
	"github.com/divya2703/covid-tracker-rest-api/entity"
	"github.com/divya2703/covid-tracker-rest-api/errors"
	"github.com/divya2703/covid-tracker-rest-api/service"
	"github.com/gorilla/mux"
)

type controller struct {
}

var (
	serv   service.Service
	rCache cache.ICache
)

type Controller interface {
	GetStateReports(response http.ResponseWriter, request *http.Request)
	GetStateReportByStateName(response http.ResponseWriter, request *http.Request)
	GetStateReportByCoordinates(response http.ResponseWriter, request *http.Request)
}

func NewController(service service.Service, cache cache.ICache) Controller {
	serv = service
	rCache = cache
	return &controller{}
}

func (*controller) GetStateReports(response http.ResponseWriter, request *http.Request) {

	// swagger:route GET /states report
	//
	// Lists state wise covid report
	//
	// This will show active number of covid cases in all states
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Deprecated: false
	//     Responses:
	//       default:
	//       200: []StateReport
	//       400: ServiceError

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")

	stateReports, err := serv.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{ErrorMessage: "Error getting the posts", StatusCode: http.StatusInternalServerError})
		return

	}
	json.NewEncoder(response).Encode(stateReports)

}
func (*controller) GetStateReportByStateName(response http.ResponseWriter, request *http.Request) {

	// swagger:route GET /states/{state} report
	// This will show active number of covid cases in a given state
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//     Parameters:
	//	   		+ name: state
	//            enum: Assam, Sikkim
	//            in: path
	//            type: string
	//            required: true
	//     Deprecated: false
	//     Responses:
	//       default:
	//       200: StateReport
	//       400: ServiceError

	response.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(request)

	stateName, _ := params["state"]
	var stateReport *entity.StateReport = rCache.Get(stateName)
	fmt.Println(stateReport)
	if stateReport == nil {

		stateReportFromDB, err := serv.FindByName(stateName)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiceError{ErrorMessage: "Error getting the reports", StatusCode: http.StatusInternalServerError})
			return

		}

		rCache.Set(stateName, stateReportFromDB)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(stateReportFromDB)
	} else {
		fmt.Println("From else")
		fmt.Println(stateReport)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(stateReport)
	}

}
func (*controller) GetStateReportByCoordinates(response http.ResponseWriter, request *http.Request) {

	// swagger:route GET /geocode report
	// This will show active number of covid cases in a given state identified by the geocode
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//	   Parameters:
	//			+ name: lat
	//			  in: query
	//            type: string
	//			  required: true
	//			+ name: long
	//			  in: query
	//            type: string
	//			  required: true
	//     Deprecated: false
	//     Responses:
	//       default:
	//       200: StateReport
	//       400: ServiceError

	response.Header().Set("Content-Type", "application/json")
	params := request.URL.Query()
	params.Set("access_token", os.Getenv("LOCATION_IQ_ACCESS_TOKEN"))
	params.Set("format", os.Getenv("LOCATION_IQ_RESPONSE_FORMAT"))
	request.URL.RawQuery = params.Encode()

	client := &http.Client{}
	method := "GET"
	req, err := http.NewRequest(method, os.Getenv("LOCATION_IQ_HOST"), nil)
	query := req.URL.Query()
	query.Add("key", os.Getenv("LOCATION_IQ_ACCESS_TOKEN"))
	query.Add("format", os.Getenv("LOCATION_IQ_RESPONSE_FORMAT"))
	query.Add("lat", params.Get("lat"))
	query.Add("lon", params.Get("long"))

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var locationInfo entity.LocationInfo
	err = json.Unmarshal(body, &locationInfo)
	stateName := locationInfo.Address.State
	stateReport, err := serv.FindByName(stateName)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{ErrorMessage: "Error getting the reports", StatusCode: http.StatusInternalServerError})
		return

	}
	json.NewEncoder(response).Encode(stateReport)

}
