package quotes

import (
	"encoding/json"
	"net/http"
)

type QuoteHttpRequest struct {
	PickupPostcode   string `json:"pickup_postcode"`
	DeliveryPostcode string `json:"delivery_postcode"`
}

type QuoteHttpRespons struct {
	QuoteHttpRequest
	Price string `json:"price"`
}

func QuoteHandler(w http.ResponseWriter, r *http.Request) {
	var quoteData QuoteHttpRequest
	json.NewDecoder(r.Body).Decode(&quoteData)

}
