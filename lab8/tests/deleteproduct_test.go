package tests

import (
	"github.com/stretchr/testify/assert"
	"lab8/api"
	"lab8/model"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
	client := api.NewShopApi(baseUrl)

	t.Run("Delete existing product", func(t *testing.T) {
		t.Parallel()

		createdProduct := createTestProduct(t, client)

		err := client.DeleteProduct(createdProduct.ID)
		assert.NoError(t, err, "Ошибка при удалении существующего продукта")

		_, err = client.GetByID(createdProduct.ID)
		assert.ErrorIs(t, err, api.NotFound, "Ожидалось, что продукт не существует после удаления")
	})

	t.Run("Delete non-existing product", func(t *testing.T) {
		t.Parallel()

		nonExistingID := "-999999"
		err := client.DeleteProduct(nonExistingID)
		assert.Error(t, err, "Ошибка при попытке удалить несуществующий продукт")
	})
}

func createTestProduct(t *testing.T, client *api.ShopApi) model.Product {
	t.Helper()
	id, err := client.AddProduct(validProductMinCategoryId)
	assert.NoError(t, err, "Не удалось создать продукт для теста")

	product, err := client.GetByID(id)
	assert.NoError(t, err, "Не получилось извлечь продукт после добавления")
	return product
}
