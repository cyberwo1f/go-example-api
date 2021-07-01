package version

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	logger   *zap.Logger
	res *response
}

type response struct {
	Version  string `json:"version"`
	Revision string `json:"revision"`
}

func NewHandler(logger *zap.Logger, ver, rev string) *Handler {
	return &Handler{
		logger: logger,
		res: &response{
			Version:  ver,
			Revision: rev,
		},
	}
}

func (h *Handler) GetVersion() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(h.res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.logger.Error("failed to marshal version", zap.Error(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})
}
