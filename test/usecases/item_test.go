package usecases

import (
	"errors"
	"testing"

	"github.com/8soat-grupo35/tech-challenge-fase1/src/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/entities"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/usecases"
	mock_repository "github.com/8soat-grupo35/tech-challenge-fase1/test/gateways/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	filterItem := entities.Item{
		Category: "LANCHE",
	}
	itemData := []entities.Item{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetAll(filterItem).Return(itemData, nil).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)

	result, err := itemUseCase.GetAll(filterItem.Category)

	assert.NoError(t, err)
	assert.Equal(t, itemData, result)
}

func TestGetAllError(t *testing.T) {
	ctrl := gomock.NewController(t)
	filterItem := entities.Item{
		Category: "LANCHE",
	}

	mockErroRepo := errors.New("error mock")
	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetAll(filterItem).Return(nil, mockErroRepo).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)

	_, err := itemUseCase.GetAll(filterItem.Category)

	assert.EqualError(t, err, "get item from repository has failed")
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToCreate := dto.ItemDto{
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemDomainToCreate, _ := entities.NewItem(itemToCreate)

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().Create(*itemDomainToCreate).Return(itemDomainToCreate, nil).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)

	createdItem, err := itemUseCase.Create(itemToCreate)

	assert.NoError(t, err)
	assert.Equal(t, itemDomainToCreate, createdItem)
}

func TestCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockErroRepo := errors.New("error mock")
	itemToCreateDto := dto.ItemDto{
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemToCreateDomain := entities.Item{
		Name:     itemToCreateDto.Name,
		Category: itemToCreateDto.Category,
		Price:    itemToCreateDto.Price,
		ImageUrl: itemToCreateDto.ImageUrl,
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().Create(itemToCreateDomain).Return(nil, mockErroRepo).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)

	_, err := itemUseCase.Create(itemToCreateDto)

	assert.EqualError(t, err, "create item on repository has failed")
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdateDto := dto.ItemDto{
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemToUpdateDomain := entities.Item{
		Name:     itemToUpdateDto.Name,
		Category: itemToUpdateDto.Category,
		Price:    itemToUpdateDto.Price,
		ImageUrl: itemToUpdateDto.ImageUrl,
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(entities.Item{
		ID: itemToUpdateDomain.ID,
	}).Return(&itemToUpdateDomain, nil).Times(1)
	itemRepo.EXPECT().Update(itemToUpdateDomain.ID, itemToUpdateDomain).Return(&itemToUpdateDomain, nil).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)
	updatedItem, err := itemUseCase.Update(itemToUpdateDomain.ID, itemToUpdateDto)

	assert.NoError(t, err)
	assert.Equal(t, itemToUpdateDomain, *updatedItem)
}

func TestUpdateNotFoundItemError(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdateDto := dto.ItemDto{
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemToUpdateDomain := entities.Item{
		ID:       1,
		Name:     itemToUpdateDto.Name,
		Category: itemToUpdateDto.Category,
		Price:    itemToUpdateDto.Price,
		ImageUrl: itemToUpdateDto.ImageUrl,
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(entities.Item{
		ID: itemToUpdateDomain.ID,
	}).Return(nil, gorm.ErrRecordNotFound).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)
	_, err := itemUseCase.Update(itemToUpdateDomain.ID, itemToUpdateDto)

	assert.EqualError(t, err, "error on obtain item to update in repository")
}

func TestUpdateError(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdateDto := dto.ItemDto{
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemToUpdateDomain := entities.Item{
		Name:     itemToUpdateDto.Name,
		Category: itemToUpdateDto.Category,
		Price:    itemToUpdateDto.Price,
		ImageUrl: itemToUpdateDto.ImageUrl,
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(entities.Item{
		ID: itemToUpdateDomain.ID,
	}).Return(&itemToUpdateDomain, nil).Times(1)
	itemRepo.EXPECT().Update(itemToUpdateDomain.ID, itemToUpdateDomain).Return(nil, gorm.ErrInvalidValue).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)
	_, err := itemUseCase.Update(itemToUpdateDomain.ID, itemToUpdateDto)

	assert.EqualError(t, err, "updated item on repository has failed")
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdate := entities.Item{
		ID:       1,
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(entities.Item{
		ID: itemToUpdate.ID,
	}).Return(&itemToUpdate, nil).Times(1)
	itemRepo.EXPECT().Delete(itemToUpdate.ID).Return(nil).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)
	err := itemUseCase.Delete(itemToUpdate.ID)

	assert.NoError(t, err)
}

func TestDeleteNotFoundItemError(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdate := entities.Item{
		ID: 1,
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(entities.Item{
		ID: itemToUpdate.ID,
	}).Return(nil, gorm.ErrRecordNotFound).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)
	err := itemUseCase.Delete(itemToUpdate.ID)

	assert.EqualError(t, err, "error on obtain item to delete in repository")
}

func TestDeleteError(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdate := entities.Item{
		ID:       1,
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(entities.Item{
		ID: itemToUpdate.ID,
	}).Return(&itemToUpdate, nil).Times(1)
	itemRepo.EXPECT().Delete(itemToUpdate.ID).Return(gorm.ErrMissingWhereClause).Times(1)

	itemUseCase := usecases.NewItemUseCase(itemRepo)
	err := itemUseCase.Delete(itemToUpdate.ID)

	assert.EqualError(t, err, "error on delete in repository")
}
