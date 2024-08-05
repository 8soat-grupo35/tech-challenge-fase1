package client

import (
	"errors"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/repository"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/service"
)

type clientService struct {
	clientRepository repository.ClientRepository
}

func NewClientService(clientRepository repository.ClientRepository) service.ClientService {
	return &clientService{
		clientRepository: clientRepository,
	}
}

func (service *clientService) GetAll() ([]domain.Client, error) {
	clients, err := service.clientRepository.GetAll()

	if err != nil {
		return []domain.Client{}, errors.New("get client from repository has failed")
	}

	return clients, nil
}

// Create implements ports.ClientService.
func (service *clientService) Create(client domain.Client) (*domain.Client, error) {
	clientSaved, err := service.clientRepository.Create(client)

	if err != nil {
		return nil, errors.New("create client on repository has failed")
	}

	return clientSaved, err
}

func (service *clientService) GetByCpf(cpf string) (*domain.Client, error) {

	client, err := service.clientRepository.GetOne(domain.Client{
		CPF: cpf,
	})

	if err != nil {
		return nil, errors.New("error on obtain client by CPF in repository")
	}

	return client, err
}

// Update implements ports.ClientService.
func (service *clientService) Update(clientId uint32, client domain.Client) (*domain.Client, error) {

	clientAlreadySaved, err := service.clientRepository.GetOne(domain.Client{
		Id: clientId,
	})

	if err != nil {
		return nil, errors.New("error on obtain client to update in repository")
	}

	if clientAlreadySaved == nil {
		return nil, errors.New("client not found to update")
	}

	clientUpdated, err := service.clientRepository.Update(clientId, client)

	if err != nil {
		return nil, errors.New("updated client on repository has failed")
	}

	return clientUpdated, err
}

// Delete implements ports.ClientService.
func (service *clientService) Delete(clientId uint32) error {
	clientAlreadySaved, err := service.clientRepository.GetOne(domain.Client{
		Id: clientId,
	})

	if err != nil {
		return errors.New("error on obtain client to delete in repository")
	}

	if clientAlreadySaved == nil {
		return errors.New("client not found to delete")
	}

	err = service.clientRepository.Delete(clientId)

	if err != nil {
		return errors.New("error on delete in repository")
	}

	return err
}
