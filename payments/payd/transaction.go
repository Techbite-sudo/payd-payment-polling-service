package payd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Techbite-sudo/payd-payment-polling-service/payments/models"
)

func GetTransactionRequests(accountID, username, password string) (*models.TransactionResponse, error) {
	url := fmt.Sprintf("https://api.mypayd.app/api/v1/accounts/%s/transaction-requests", accountID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Add("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d\nResponse body: %s", resp.StatusCode, string(body))
	}

	var transactionResponse models.TransactionResponse
	err = json.Unmarshal(body, &transactionResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &transactionResponse, nil
}
