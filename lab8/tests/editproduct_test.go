package tests

import (
	"lab8/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"lab8/api"
)

func deleteTestProduct(t *testing.T, client *api.ShopApi, id string) {
	t.Helper()
	err := client.DeleteProduct(id)
	assert.NoError(t, err, "Ошибка при удалении тестового продукта")
}

func runValidEditTestCase(t *testing.T, client *api.ShopApi, testcaseName string, updatedProduct model.Product) {

	t.Run(testcaseName, func(t *testing.T) {
		t.Parallel()
		id := createTestProduct(t, client).ID
		defer deleteTestProduct(t, client, id)

		updatedProduct.ID = id
		err := client.EditProduct(updatedProduct)
		assert.NoError(t, err, "Ошибка при редактировании продукта")

		resultProduct, err := client.GetByID(updatedProduct.ID)
		assert.NoError(t, err, "Не получилось извлечь продукт после редактирования")

		assertEquals(t, updatedProduct, resultProduct)
	})
}

func runInvalidEditTestCase(t *testing.T, client *api.ShopApi, testcaseName string, updatedProduct model.Product) {

	t.Run(testcaseName, func(t *testing.T) {
		t.Parallel()
		id := createTestProduct(t, client).ID
		defer deleteTestProduct(t, client, id)

		updatedProduct.ID = id
		err := client.EditProduct(updatedProduct)
		assert.ErrorIs(t, err, api.ErrBadRequest, "Ожидалась ошибка BadRequest")

		resultProduct, err := client.GetByID(updatedProduct.ID)
		assert.NoError(t, err, "Не получилось извлечь продукт после редактирования")

		assertEquals(t, validProductMinCategoryId, resultProduct)
	})
}

func TestValidEditProduct(t *testing.T) {
	client := api.NewShopApi(baseUrl)
	for testcaseName, updatedProduct := range validCases {
		runValidEditTestCase(t, client, testcaseName, updatedProduct)
	}
}

func TestInvalidEditProduct(t *testing.T) {
	client := api.NewShopApi(baseUrl)
	for testcaseName, updatedProduct := range invalidCases {
		runInvalidEditTestCase(t, client, testcaseName, updatedProduct)
	}
}
