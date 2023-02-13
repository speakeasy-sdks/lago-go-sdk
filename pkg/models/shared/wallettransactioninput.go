package shared

type WalletTransactionInputWalletTransaction struct {
	GrantedCredits *float64 `json:"granted_credits,omitempty"`
	PaidCredits    *float64 `json:"paid_credits,omitempty"`
	WalletID       *string  `json:"wallet_id,omitempty"`
}

type WalletTransactionInput struct {
	WalletTransaction *WalletTransactionInputWalletTransaction `json:"wallet_transaction,omitempty"`
}
