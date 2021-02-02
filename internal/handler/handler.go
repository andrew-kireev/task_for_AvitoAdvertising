package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"tast_for_AvitoAdvertising/internal/model"
	"tast_for_AvitoAdvertising/store"
)

const pageSize = 5

type Handler struct {
	router *mux.Router
	store  *store.Store
	logger *logrus.Logger
}

type CreationResponse struct {
	ResultCode string `json:"result"`
	AdvertId   int    `json:"id"`
}


func (handler *Handler) ConfigHandler(store *store.Store, logger *logrus.Logger) {
	handler.router = mux.NewRouter()
	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	handler.router.HandleFunc("/advert/create", handler.HandleAdvertCreation)
	handler.router.HandleFunc("/advert/get/{id:[0-9]+}", handler.HandlerGetAdvert)
	handler.router.HandleFunc("/advert/list/{page:[0-9]+}", handler.HandlerGetAllAdverts)
	handler.store = store
	handler.logger = logger
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func (handler *Handler) HandleAdvertCreation(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("HandleAdvertCreation")
	advert := &model.Advert{}
	advert.Name = r.FormValue("name")
	advert.Description = r.FormValue("description")
	advert.PhotoLinks = r.FormValue("links")
	advert.Price, _ = strconv.Atoi(r.FormValue("price"))

	advert, err := handler.store.Adverts().CreateAdvert(advert)
	handler.logger.Info(advert)

	if err != nil {
		handler.logger.Errorf("error in CreateAdvert: %v", err)
		response := CreatFailedResp()
		w.Write(response)
		return
	}
	response := CreationResponse{}
	response.AdvertId = advert.AdvertId
	response.ResultCode = "success"
	resp, err := json.Marshal(response)
	if err != nil {
		handler.logger.Errorf("error in maishaling json: %v", err)
		response := CreatFailedResp()
		w.Write(response)
		return
	}
	w.Write(resp)
}

func (handler *Handler) HandlerGetAdvert(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("HandlerGetAdvert")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	optionalFields := r.URL.Query().Get("fields")
	advert, err := handler.store.Adverts().GetAdvertById(id, optionalFields)
	if err != nil {
		handler.logger.Errorf("error in GetAdvertById: %v", err)
		w.Write([]byte("{}"))
		return
	}
	response, err := json.Marshal(advert)
	if err != nil {
		handler.logger.Errorf("error in marshaling json: %v", err)
		w.Write([]byte("{}"))
		return
	}

	w.Write(response)
}

func (handler *Handler) HandlerGetAllAdverts(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("HandlerGetAllAdverts")
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	sort := r.URL.Query().Get("sort")

	adverts, err := handler.store.Adverts().GetAllAdverts(sort)
	if err != nil {
		handler.logger.Errorf("error in GetAllAdverts: %v", err)
		w.Write([]byte("[]"))
		return
	}

	var pageEnd = (page + 1) * pageSize
	var pageBegin = page * pageSize
	if pageEnd > len(adverts) {
		pageEnd = len(adverts)
	}
	if pageBegin > len(adverts) {
		pageBegin = len(adverts)
	}

	adverts = adverts[pageBegin:pageEnd]
	response, err := json.Marshal(adverts)
	if err != nil {
		handler.logger.Errorf("error in marshaling json: %v", err)
		w.Write([]byte("[]"))
		return
	}
	w.Write(response)
}

func CreatFailedResp() []byte {
	response := CreationResponse{}
	response.AdvertId = -1
	response.ResultCode = "failed"
	resp, _ := json.Marshal(response)
	return resp
}
