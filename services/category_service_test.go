package services

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryGetNotFound(t *testing.T) {
	// program mock / arrange
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// act
	category, err := categoryService.Get("1")

	// assert
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryGetSuccess(t *testing.T) {
	Category := entity.Category{
		Id:   "2",
		Name: "fuad",
	}

	categoryRepository.Mock.On("FindById", "2").Return(Category)

	result, err := categoryService.Get("2")

	assert.Equal(t, result.Name, Category.Name)
	assert.Equal(t, result.Id, Category.Id)
	assert.Nil(t, err)
}

func BenchmarkCategoryService_GetSucess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Category := entity.Category{
			Id:   "2",
			Name: "fuad",
		}

		categoryRepository.Mock.On("FindById", "2").Return(Category)

		categoryService.Get("2")
	}
}
