package model

type TransferLog struct {
	ID        uint64
	UserID1   uint64
	UserID2   uint64
	Amount    int32
	CreatedAt uint64
	Status    byte
}
