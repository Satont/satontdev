package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/satont/satontdev/internal/config"
	"github.com/satont/satontdev/views/layout"
	"github.com/satont/satontdev/views/pages/index"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg, err := config.New(filepath.Join(pwd, "config"))
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc(
		"GET /", func(w http.ResponseWriter, r *http.Request) {
			layout.Layout(index.Index()).Render(r.Context(), w)
		},
	)

	slog.Info(fmt.Sprintf("Listening on %v port", cfg.Application.Port))

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", cfg.Application.Port), mux); err != nil {
		panic(err)
	}
}
