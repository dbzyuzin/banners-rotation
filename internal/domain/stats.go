package domain

type Stats struct {
	Id          Id
	SlotId      Id
	BannerId    Id
	UserGroupId Id
	ShowCount   int64
	ClickCount  int64
}
