package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tast_for_AvitoAdvertising/internal/handler"
	"tast_for_AvitoAdvertising/store"
)

type HttpServer struct {
	Conf    *Config
	Handler *handler.Handler
	store   *store.Store
	logger  *logrus.Logger
}

func NewServer(config *Config) (*HttpServer, error) {
	serv := &HttpServer{
		Conf:   config,
		Handler: &handler.Handler{},
		logger: logrus.New(),
	}

	err := serv.ConfigLogger()
	if err != nil {
		return nil, err
	}

	if err = serv.ConfigStore(); err != nil {
		serv.logger.Error("error creating db")
		return nil, err
	}
	serv.Handler.ConfigHandler(serv.store, serv.logger)

	serv.logger.Info("server created")
	return serv, nil
}

func Start(config *Config) error {
	serv, err := NewServer(config)
	if err != nil {
		return err
	}

	return http.ListenAndServe(config.BindAddr, serv.Handler)
}

func (serv *HttpServer) ConfigStore() error {
	serv.logger.Info(serv.Conf.StoreConfig.DataBaseUrl)
	st := store.NewStore(serv.Conf.StoreConfig)

	if err := st.Open(); err != nil {
		return err
	}
	serv.store = st

	return nil
}

func (serv *HttpServer) ConfigLogger() error {
	level, err := logrus.ParseLevel(serv.Conf.LogLevel)
	if err != nil {
		return err
	}

	serv.logger.SetLevel(level)
	return nil
}