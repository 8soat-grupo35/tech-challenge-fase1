package repository

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

type ClientRepository interface {
	GetAll() ([]domain.Client, error)
	GetOne(domain.Client) (*domain.Client, error)
	Create(client domain.Client) (*domain.Client, error)
	Update(clientId uint32, client domain.Client) (*domain.Client, error)
	Delete(clientId uint32) error
}
