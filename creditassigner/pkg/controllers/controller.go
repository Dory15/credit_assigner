package controllers

import (
	"creditassigner/pkg/logger"
	"creditassigner/pkg/models"
	"creditassigner/pkg/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var fileName string = "controller"
var controllerLogger logger.ILogger

var service services.Service

func init() {
	controllerLogger = logger.NewLoggerInstace(fileName)
	service = services.NewService()
}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/credit-assignment", assignController)
	router.Post("/statistics", getStatisticsController)
	return router
}

func assignController(res http.ResponseWriter, req *http.Request) {

	var investment models.Investment

	res.Header().Set("Content-Type", "application/json")
	jsonResponseWriter := json.NewEncoder(res)

	err := json.NewDecoder(req.Body).Decode(&investment)

	if err != nil {
		controllerLogger.Error(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	credits, err := service.CreditAssignerService(investment)

	if err != nil {
		controllerLogger.Error(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonResponseWriter.Encode(credits)

}

func getStatisticsController(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	jsonResponseWriter := json.NewEncoder(res)

	statistics, err := service.GetStatisticsService()

	if err != nil {
		controllerLogger.Error(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponseWriter.Encode(statistics)

}
