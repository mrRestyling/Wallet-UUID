package service

type Service struct {
	Storage StorageItn
}

type StorageItn interface {
}

func New(s StorageItn) *Service {
	return &Service{
		Storage: s,
	}
}
