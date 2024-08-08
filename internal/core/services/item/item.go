package services

import (
	"errors"
	"log"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	custom_errors "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/errors"
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

func (service *itemService) GetAll(category string) ([]domain.Item, error) {
	filter := domain.Item{}

	if category != "" {
		filter.Category = category
		err := filter.ValidateCategory()
	
		if err != nil {
			return []domain.Item{}, &custom_errors.BadRequestError{
				Message: err.Error(),
			}
		}
	}


	items, err := service.itemRepository.GetAll(filter)

	if err != nil {
		return []domain.Item{}, &custom_errors.DatabaseError{
			Message: "get item from repository has failed",
		}
	}

	return items, nil
}

// Create implements ports.ItemService.
func (service *itemService) Create(item dto.ItemDto) (*domain.Item, error) {

	newItem, err := domain.NewItem(item)

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	itemSaved, err := service.itemRepository.Create(*newItem)

	if err != nil {
		return nil, errors.New("create item on repository has failed")
	}

	return itemSaved, err
}

// Update implements ports.ItemService.
func (service *itemService) Update(itemId uint32, item dto.ItemDto) (*domain.Item, error) {

	itemToUpdate, err := domain.NewItem(item)

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	itemAlreadySaved, err := service.itemRepository.GetOne(domain.Item{
		ID: itemId,
	})

	if err != nil {
		log.Println(err.Error())
		return nil, &custom_errors.DatabaseError{
			Message: "error on obtain item to update in repository",
		}
	}

	if itemAlreadySaved == nil {
		return nil, &custom_errors.NotFoundError{
			Message: "item not found to update",
		}
	}

	itemUpdated, err := service.itemRepository.Update(itemId, *itemToUpdate)

	if err != nil {
		log.Println(err.Error())
		return nil, &custom_errors.DatabaseError{
			Message: "updated item on repository has failed",
		}
	}

	return itemUpdated, err
}

// Delete implements ports.ItemService.
func (service *itemService) Delete(itemId uint32) error {
	itemAlreadySaved, err := service.itemRepository.GetOne(domain.Item{
		ID: itemId,
	})

	if err != nil {
		log.Println(err.Error())
		return &custom_errors.DatabaseError{
			Message: "error on obtain item to delete in repository",
		}
	}

	if itemAlreadySaved == nil {
		return &custom_errors.NotFoundError{
			Message: "item not found to delete",
		}
	}

	err = service.itemRepository.Delete(itemId)

	if err != nil {
		return &custom_errors.DatabaseError{
			Message: "error on delete in repository",
		}
	}

	return err
}
