package exmo

import "net/url"

type PairSettingStruct struct {
	MinQuantity     string `json:"min_quantity"`
	MaxQuantity     string `json:"max_quantity"`
	MinPrice        string `json:"min_price"`
	MaxPrice        string `json:"max_price"`
	MaxAmount       string `json:"max_amount"`
	MinAmount       string `json:"min_amount"`
	PricePrecision  int    `json:"price_precision"`
	MakerCommission string `json:"commission_maker_percent"`
	TakerCommission string `json:"commission_taker_percent"`
}

type PairSettings map[string]PairSettingStruct

type PairSettingService struct {
	c *Client
}

func (a *PairSettingService) Get() (PairSettings, error) {
	params := url.Values{}
	req, err := a.c.newRequest("GET", "pair_settings", params)

	if err != nil {
		return PairSettings{}, err
	}

	var v PairSettings
	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return PairSettings{}, err
	}

	return v, nil
}
