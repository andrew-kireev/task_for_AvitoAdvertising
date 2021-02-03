package handler

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http/httptest"
	"tast_for_AvitoAdvertising/internal/model"
	"tast_for_AvitoAdvertising/store"
	"testing"
)

func TestHandler_HandleAdvertCreation(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	rep := NewMockAdvertRepositoryInterface(ctrl)

	store := &store.Store{
		AdvertRep: rep,
	}

	handler := &Handler{
		store:  store,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	advert := &model.Advert{Name: "some name",
		Description: "desc", PhotoLinks: "link1_link2",
		Price: 500, CreationDate: "",
	}

	rep.EXPECT().CreateAdvert(advert).AnyTimes().Return(advert, nil)
	req := httptest.NewRequest("GET", "/", nil)
	req.Form = make(map[string][]string)
	req.Form.Add("name", "some name")
	req.Form.Add("description", "desc")
	req.Form.Add("links", "link1_link2")
	req.Form.Add("price", "500")

	w := httptest.NewRecorder()
	handler.HandleAdvertCreation(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "{\"result\":\"success\",\"id\":0}" {
		t.Errorf("not correct answer")
	}
}

func TestHandler_HandlerGetAdvert(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	rep := NewMockAdvertRepositoryInterface(ctrl)

	store := &store.Store{
		AdvertRep: rep,
	}

	handler := &Handler{
		router: mux.NewRouter(),
	}
	handler.ConfigHandler(store, logrus.New())

	//advert := &model.Advert{Name: "some name",
	//	Description: "desc", PhotoLinks: "link1_link2",
	//	Price: 500, CreationDate: "",
	//}

	rep.EXPECT().GetAdvertById(0, "").AnyTimes()
	req := httptest.NewRequest("GET", "/advert/get/3", nil)

	w := httptest.NewRecorder()
	handler.HandlerGetAdvert(w, req)
}

func TestCreatFailedResp(t *testing.T) {
	resp := string(CreatFailedResp())
	correcResp := "{\"result\":\"failed\",\"id\":-1}"
	if correcResp != resp {
		t.Errorf("results not match, want %v, have %v", correcResp, resp)
	}
}
