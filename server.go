// Package classification COVID_REPORT_APP
//
// This application provides 3 main APIs,
//
// 1) Get statewise covid report data for all the states of India
//
// 2) Get covid report data for a given state
//
// 3) Given latitude and longitude of a user, provide covid related information for the
//    state the user belongs to.
//
//     Schemes: https
//     Host: covid-tracker-rest-api.herokuapp.com
//     BasePath: /api
//     Version: 1.0.0
//     Contact: Divya Kumari<divya27@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/divya2703/covid-tracker-rest-api/cache"
	"github.com/divya2703/covid-tracker-rest-api/controllers"
	"github.com/divya2703/covid-tracker-rest-api/repository"
	"github.com/divya2703/covid-tracker-rest-api/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	//config     db.Configuration       = db.GetConfiguration()
	repo       repository.Repository  = repository.NewMongoRepository()
	redisCache cache.ICache           = cache.NewRedisCache("redis-12426.c264.ap-south-1-1.ec2.cloud.redislabs.com:12426")
	serv       service.Service        = service.NewService(repo)
	controller controllers.Controller = controllers.NewController(serv, redisCache)
)

func main() {
	r := mux.NewRouter()

	//Test snippet for ttl
	// fmt.Println(redisCache.Get("Sahib"))
	// var stateReport *entity.StateReport
	// stateReport = new(entity.StateReport)
	// stateReport.Active = 1220
	// stateReport.State = "Sahib"
	// stateReport.Last_Updated_Time = "23-09-2021"
	// redisCache.Set("Sahib", stateReport)
	// for i := 0; i < 1000; i++ {
	// 	fmt.Println(redisCache.Get("Sahib"))
	// }

	r.HandleFunc("/api/states", controller.GetStateReports).Methods("GET")
	r.HandleFunc("/api/geocode", controller.GetStateReportByCoordinates).Methods("GET")
	r.HandleFunc("/api/states/{state}", controller.GetStateReportByStateName).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	PORT := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(PORT, handler))
}
