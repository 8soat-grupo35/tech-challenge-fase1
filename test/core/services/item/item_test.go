package service_item_test

import (
	"errors"
	"testing"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	services "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/item"
	mock_repository "github.com/8soat-grupo35/tech-challenge-fase1/test/core/ports/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	filterItem := domain.Item{}
	itemData := []domain.Item{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetAll(filterItem).Return(itemData, nil).Times(1)

	itemService := services.NewItemService(itemRepo)

	result, err := itemService.GetAll(filterItem)

	assert.NoError(t, err)
	assert.Equal(t, itemData, result)
}

func TestGetAllError(t *testing.T) {
	ctrl := gomock.NewController(t)
	filterItem := domain.Item{}

	mockErroRepo := errors.New("error mock")
	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetAll(filterItem).Return(nil, mockErroRepo).Times(1)

	itemService := services.NewItemService(itemRepo)

	_, err := itemService.GetAll(filterItem)

	assert.EqualError(t, err, "get item from repository has failed")
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToCreate := domain.Item{
		ID:       1,
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().Create(itemToCreate).Return(&itemToCreate, nil).Times(1)

	itemService := services.NewItemService(itemRepo)

	createdItem, err := itemService.Create(itemToCreate)

	assert.NoError(t, err)
	assert.Equal(t, itemToCreate, *createdItem)
}

func TestCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockErroRepo := errors.New("error mock")
	itemToCreate := domain.Item{
		ID:       1,
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().Create(itemToCreate).Return(nil, mockErroRepo).Times(1)

	itemService := services.NewItemService(itemRepo)

	_, err := itemService.Create(itemToCreate)

	assert.EqualError(t, err, "create item on repository has failed")
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdate := domain.Item{
		ID:       1,
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(domain.Item{
		ID: itemToUpdate.ID,
	}).Return(&itemToUpdate, nil).Times(1)
	itemRepo.EXPECT().Update(itemToUpdate.ID, itemToUpdate).Return(&itemToUpdate, nil).Times(1)

	itemService := services.NewItemService(itemRepo)
	updatedItem, err := itemService.Update(itemToUpdate.ID, itemToUpdate)

	assert.NoError(t, err)
	assert.Equal(t, itemToUpdate, *updatedItem)
}

func TestUpdateNotFoundItemError(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdate := domain.Item{
		ID:       1,
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(domain.Item{
		ID: itemToUpdate.ID,
	}).Return(nil, gorm.ErrRecordNotFound).Times(1)

	itemService := services.NewItemService(itemRepo)
	_, err := itemService.Update(itemToUpdate.ID, itemToUpdate)

	assert.EqualError(t, err, "error on obtain item to update in repository")
}

func TestUpdateError(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdate := domain.Item{
		ID:       1,
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(domain.Item{
		ID: itemToUpdate.ID,
	}).Return(&itemToUpdate, nil).Times(1)
	itemRepo.EXPECT().Update(itemToUpdate.ID, itemToUpdate).Return(nil, gorm.ErrInvalidValue).Times(1)

	itemService := services.NewItemService(itemRepo)
	_, err := itemService.Update(itemToUpdate.ID, itemToUpdate)

	assert.EqualError(t, err, "updated item on repository has failed")
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdate := domain.Item{
		ID:       1,
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(domain.Item{
		ID: itemToUpdate.ID,
	}).Return(&itemToUpdate, nil).Times(1)
	itemRepo.EXPECT().Delete(itemToUpdate.ID).Return(nil).Times(1)

	itemService := services.NewItemService(itemRepo)
	err := itemService.Delete(itemToUpdate.ID)

	assert.NoError(t, err)
}

func TestDeleteNotFoundItemError(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdate := domain.Item{
		ID: 1,
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(domain.Item{
		ID: itemToUpdate.ID,
	}).Return(nil, gorm.ErrRecordNotFound).Times(1)

	itemService := services.NewItemService(itemRepo)
	err := itemService.Delete(itemToUpdate.ID)

	assert.EqualError(t, err, "error on obtain item to delete in repository")
}

func TestDeleteError(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemToUpdate := domain.Item{
		ID:       1,
		Name:     "X-BURGUER",
		Category: "LANCHE",
		Price:    20,
		ImageUrl: "www.aaa.com.br",
	}

	itemRepo := mock_repository.NewMockItemRepository(ctrl)
	itemRepo.EXPECT().GetOne(domain.Item{
		ID: itemToUpdate.ID,
	}).Return(&itemToUpdate, nil).Times(1)
	itemRepo.EXPECT().Delete(itemToUpdate.ID).Return(gorm.ErrMissingWhereClause).Times(1)

	itemService := services.NewItemService(itemRepo)
	err := itemService.Delete(itemToUpdate.ID)

	assert.EqualError(t, err, "error on delete in repository")
}
