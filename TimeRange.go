package zksync

type TimeRange struct {
	ValidFrom  uint32 `json:"validFrom"`
	ValidUntil uint32 `json:"validUntil"`
}

func DefaultTimeRange() *TimeRange {
	return &TimeRange{
		ValidFrom:  0,
		ValidUntil: 4294967295,
	}
}
