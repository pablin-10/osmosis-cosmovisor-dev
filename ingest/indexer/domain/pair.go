package domain

import (
	"time"
)

// Pair represents a pair of tokens in a pool and message to be published to PubSub
type Pair struct {
	PoolID     uint64    `json:"pool_id"`
	MultiAsset bool      `json:"multi_asset"`
	Denom0     string    `json:"denom_0"`
	IdxDenom0  uint8     `json:"idx_denom_0"`
	Denom1     string    `json:"denom_1"`
	IdxDenom1  uint8     `json:"idx_denom_1"`
	FeeBps     uint64    `json:"fee_bps"`
	IngestedAt time.Time `json:"ingested_at"`
}
