package service

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

type ClientService interface {
	GetAll() ([]domain.Client, error)
	Create(domain.Client) (*domain.Client, error)
	GetByCpf(cpf string) (*domain.Client, error)
	Update(clientId uint32, client domain.Client) (*domain.Client, error)
	Delete(clientId uint32) error
}
