package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eduardotashiro/cielo-status-page/internal/domain"
)

func Pegaincidentes() (domain.StatusResponse, error) {
	resp, err := http.Get("https://status.cielo.com.br/api/v2/incidents.json")
	if err != nil {
		return domain.StatusResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.StatusResponse{}, fmt.Errorf("unexpected HTTP status: %d: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.StatusResponse{}, err
	}

	var r domain.StatusResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return domain.StatusResponse{}, err
	}
	return r, nil
}
