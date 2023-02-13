package shared

type WalletUpdateInputWallet struct {
	ExpirationAt *string `json:"expiration_at,omitempty"`
	Name         *string `json:"name,omitempty"`
}

type WalletUpdateInput struct {
	Wallet *WalletUpdateInputWallet `json:"wallet,omitempty"`
}
