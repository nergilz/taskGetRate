package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github/nergilz/taskGetRate/internal/domain"
	"io"
	"net/http"
)

type AppTransport struct {
	baseUrl string
}

func New(url string) *AppTransport {
	return &AppTransport{
		baseUrl: url,
	}
}

func (t *AppTransport) GetDataFromGrinexApi(ctx context.Context, market string) (domain.RateResponse, error) {
	params := fmt.Sprintf("depth?market=%srub", market)
	url := fmt.Sprintf("%s%s", t.baseUrl, params)

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return domain.RateResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return domain.RateResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.RateResponse{}, err
	}

	var data domain.GrinexResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return domain.RateResponse{}, err
	}

	return domain.RateResponse{
		Market:    market,
		TimeStamp: data.TimeStamp,
		Ask:       data.Asks[0],
		Bid:       data.Bids[0],
	}, nil
}
