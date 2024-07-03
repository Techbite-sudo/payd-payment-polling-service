package payd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Techbite-sudo/payd-payment-polling-service/payments/models"
)

func RunPayment(apiUsername, apiPassword string, paymentReq models.PaymentRequest) (*models.PaymentResponse, error) {
	url := "https://api.mypayd.app/api/v2/payments"
	method := "POST"

	payloadBytes, err := json.Marshal(paymentReq)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payment request: %v", err)
	}
	payload := strings.NewReader(string(payloadBytes))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	auth := base64.StdEncoding.EncodeToString([]byte(apiUsername + ":" + apiPassword))
	req.Header.Set("Authorization", "Basic "+auth)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d\nResponse body: %s", res.StatusCode, string(body))
	}

	var paymentResponse models.PaymentResponse
	err = json.Unmarshal(body, &paymentResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &paymentResponse, nil
}
