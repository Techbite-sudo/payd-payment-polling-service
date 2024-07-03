package models

type PaymentRequest struct {
	Username    string  `json:"username"`
	NetworkCode string  `json:"network_code"`
	Amount      float64 `json:"amount"`
	PhoneNumber string  `json:"phone_number"`
	Narration   string  `json:"narration"`
	Currency    string  `json:"currency"`
	CallbackURL string  `json:"callback_url"`
}

type PaymentResponse struct {
	Data                 map[string]interface{} `json:"data"`
	MerchantRequestID    string                 `json:"merchantRequestID"`
	CheckoutRequestID    string                 `json:"checkoutRequestID"`
	TransactionReference string                 `json:"transactionReference"`
	Message              string                 `json:"message"`
	CustomerMessage      string                 `json:"customerMessage"`
	PaymentGateway       string                 `json:"paymentGateway"`
}

type WebhookResponse struct {
	TransactionReference string  `json:"transaction_reference"`
	PhoneNumber          string  `json:"phone_number"`
	ResultCode           string  `json:"result_code"`
	Remarks              string  `json:"remarks"`
	ThirdPartyTransID    string  `json:"third_party_trans_id"`
	Amount               float64 `json:"amount"`
	TransactionDate      string  `json:"transaction_date"`
	ForwardURL           string  `json:"forward_url"`
	OrderID              string  `json:"order_id"`
}

type TransactionRequest struct {
	ID                  string  `json:"id"`
	Currency            string  `json:"currency"`
	CountryCode         string  `json:"country_code"`
	Phone               string  `json:"phone"`
	Fullname            string  `json:"fullname"`
	Email               string  `json:"email"`
	AccountID           string  `json:"account_id"`
	UserID              string  `json:"user_id"`
	Reference           string  `json:"reference"`
	Channel             string  `json:"channel"`
	TransactionType     string  `json:"transaction_type"`
	Location            string  `json:"location"`
	TotalCost           float64 `json:"total_cost"`
	Amount              float64 `json:"amount"`
	Charges             float64 `json:"charges"`
	Status              string  `json:"status"`
	ReceiverName        string  `json:"receiver_name"`
	ReceiverAccount     string  `json:"receiver_account"`
	Description         string  `json:"description"`
	CallbackURL         string  `json:"callback_url"`
	Source              string  `json:"source"`
	WPOrderID           string  `json:"wp_order_id"`
	TransactionCategory string  `json:"transaction_category"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
}

type TransactionResponse struct {
	TransactionRequests []TransactionRequest `json:"transaction_requests"`
	Pagination          Pagination           `json:"pagination"`
}

type Pagination struct {
	Count    int     `json:"count"`
	NextPage *string `json:"next_page"`
	NumPages int     `json:"num_pages"`
	Page     int     `json:"page"`
	Per      int     `json:"per"`
	PrevPage *string `json:"prev_page"`
}
