package shared

type WalletTransactionObjectStatusEnum string

const (
	WalletTransactionObjectStatusEnumPending WalletTransactionObjectStatusEnum = "pending"
	WalletTransactionObjectStatusEnumSettled WalletTransactionObjectStatusEnum = "settled"
)

type WalletTransactionObjectTransactionTypeEnum string

const (
	WalletTransactionObjectTransactionTypeEnumInbound  WalletTransactionObjectTransactionTypeEnum = "inbound"
	WalletTransactionObjectTransactionTypeEnumOutbound WalletTransactionObjectTransactionTypeEnum = "outbound"
)

type WalletTransactionObject struct {
	Amount          *float64                                    `json:"amount,omitempty"`
	CreatedAt       *string                                     `json:"created_at,omitempty"`
	CreditAmount    *float64                                    `json:"credit_amount,omitempty"`
	LagoID          *string                                     `json:"lago_id,omitempty"`
	LagoWalletID    *string                                     `json:"lago_wallet_id,omitempty"`
	SettledAt       *string                                     `json:"settled_at,omitempty"`
	Status          *WalletTransactionObjectStatusEnum          `json:"status,omitempty"`
	TransactionType *WalletTransactionObjectTransactionTypeEnum `json:"transaction_type,omitempty"`
}
