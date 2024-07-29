package services

import (
	"errors"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/repository"
    "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/service"
)

type itemService struct {
	itemRepository repository.ItemRepository
}

func NewItemService(itemRepository repository.ItemRepository) service.ItemService {
	return &itemService{
		itemRepository: itemRepository,
	}
}

func (service *itemService) GetAll(filter domain.Item) ([]domain.Item, error) {
	items, err := service.itemRepository.GetAll(filter)

	if err != nil {
		return []domain.Item{}, errors.New("get item from repository has failed")
	}

	return items, nil
}

// Create implements ports.ItemService.
func (service *itemService) Create(item domain.Item) (*domain.Item, error) {
	itemSaved, err := service.itemRepository.Create(item)

	if err != nil {
		return nil, errors.New("create item on repository has failed")
	}

	return itemSaved, err
}

// Update implements ports.ItemService.
func (service *itemService) Update(itemId uint32, item domain.Item) (*domain.Item, error) {

	itemAlreadySaved, err := service.itemRepository.GetOne(domain.Item{
		ID: itemId,
	})

	if err != nil {
		return nil, errors.New("error on obtain item to update in repository")
	}

	if itemAlreadySaved == nil {
		return nil, errors.New("item not found to update")
	}

	itemUpdated, err := service.itemRepository.Update(itemId, item)

	if err != nil {
		return nil, errors.New("updated item on repository has failed")
	}

	return itemUpdated, err
}

// Delete implements ports.ItemService.
func (service *itemService) Delete(itemId uint32) error {
	itemAlreadySaved, err := service.itemRepository.GetOne(domain.Item{
		ID: itemId,
	})

	if err != nil {
		return errors.New("error on obtain item to delete in repository")
	}

	if itemAlreadySaved == nil {
		return errors.New("item not found to delete")
	}

	err = service.itemRepository.Delete(itemId)

	if err != nil {
		return errors.New("error on delete in repository")
	}

	return err
}
