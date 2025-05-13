package tests

import (
	"github.com/stretchr/testify/assert"
	"lab8/api"
	"lab8/model"
	"testing"
)

func TestValidAddProduct(t *testing.T) {
	client := api.NewShopApi(baseUrl)

	for testcaseName, product := range validCases {
		t.Run(testcaseName, func(t *testing.T) {
			t.Parallel()
			runValidTestCase(t, client, product)
		})
	}
}

func TestInvalidAddProduct(t *testing.T) {
	client := api.NewShopApi(baseUrl)

	for testcaseName, product := range invalidCases {
		t.Run(testcaseName, func(t *testing.T) {
			t.Parallel()
			runInvalidTestCase(t, client, product)
		})
	}
}

func runValidTestCase(t *testing.T, client *api.ShopApi, product model.Product) {
	id, err := client.AddProduct(product)
	assert.NoError(t, err, "Ошибка при добавлении продукта")

	createdProduct, err := client.GetByID(id)
	assert.NoError(t, err, "Не получилось извлечь продукт после добавления")

	defer func() {
		err := client.DeleteProduct(createdProduct.ID)
		assert.NoError(t, err, "Ошибка при удалении тестового продукта")
	}()

	assertEquals(t, product, createdProduct)
	assert.NotEqual(t, "0", createdProduct.ID, "ID продукта не должен быть '0'")
	assert.NotEmpty(t, createdProduct.Alias, "Alias продукта должен быть сгенерирован автоматически")
}

func runInvalidTestCase(t *testing.T, client *api.ShopApi, product model.Product) {
	_, err := client.AddProduct(product)
	assert.ErrorIs(t, err, api.ErrBadRequest, "Ожидалась ошибка BadRequest для невалидных данных")
}
