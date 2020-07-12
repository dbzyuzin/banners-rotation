package domain

type Banner struct {
	Id          Id
	Description string
}

type BannerRepository interface {
	Insert(banner Banner) (Id, error)
	Update(banner Banner) error
	Delete(id Id) error
}
