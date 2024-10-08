package url_repository

import (
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
)

type Repository interface {
	CreateULR(url *url.URL) (int64, error)
	GetAllByUser(id int) ([]*url.URL, error)
	GetByName(description string) ([]url.URL, error)
	UpdateById(u *url.URL) error
	DeleteById(id int) error
}
