package usecase

import (
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"
)

type URLUseCase struct {
	URLRepository url_repository.Repository
}

func NewURLService(URLRepository url_repository.Repository) *URLUseCase {
	return &URLUseCase{
		URLRepository: URLRepository,
	}
}

func (uc *URLUseCase) Create() error {
	entitiesURL := &url.URL{
		OriginalURL:    "Origin",
		DestinationURL: "string",
		UserID:         1,
	}

	_, err := uc.URLRepository.CreateULR(entitiesURL)
	if err != nil {
		return err
	}

	return nil
}

func (uc *URLUseCase) GetAll() ([]*url.URL, error) {
	urls, err := uc.URLRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return urls, err
}
