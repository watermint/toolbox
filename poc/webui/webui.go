package webui

import (
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type WebUI struct {
	Port   int
	Logger *zap.Logger
}

func (z *WebUI) FlagConfig(f *flag.FlagSet) {
	f.IntVar(&z.Port, "port", 18080, "Port number")
}

func (z *WebUI) Start() string {
	contentStatic := rice.MustFindBox("static").HTTPBox()
	//	contentWeb := rice.MustFindBox("pages").HTTPBox()

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(contentStatic))
	//	router.Handle("/pages", http.FileServer(contentWeb))

	addr := fmt.Sprintf(":%d", z.Port)
	z.Logger.Info("Starting server", zap.String("addr", addr))
	if err := http.ListenAndServe(addr, router); err != nil {
		z.Logger.Error(
			"Unable to start server",
			zap.String("addr", addr),
			zap.Error(err),
		)
		return ""
	}

	return addr
}
