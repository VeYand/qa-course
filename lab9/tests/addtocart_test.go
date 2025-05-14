package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"lab9/config"
	"lab9/pages"
	"testing"
)

func TestAddToCartInMainPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		product := config.CitizenProduct
		productPage := pages.ProductPage{}
		productPage.Init(driver)
		err := productPage.OpenPage("")
		assert.NoError(t, err, "Не удалось открыть главную страницу")

		err = productPage.AddToCart(product.ID)
		assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

		err = productPage.IsProductInCart(product.Name, product.Price, config.QuantityProductsOne)
		assert.NoError(t, err, "Товар отсутствует в корзине или неверные параметры")
	}

	runTest(t, testFunc)
}

func TestAddOneToCartInProductPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		product := config.CitizenProduct
		productPage := pages.ProductPage{}
		productPage.Init(driver)
		err := productPage.OpenPage(product.URL)
		assert.NoError(t, err, "Не удалось открыть страницу товара")

		err = productPage.AddToCart(product.ID)
		assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

		err = productPage.IsProductInCart(product.Name, product.Price, config.QuantityProductsOne)
		assert.NoError(t, err, "Товар отсутствует в корзине или неверные параметры")
	}

	runTest(t, testFunc)
}

func TestAddSeveralToCartInProductPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		product := config.CitizenProduct
		productPage := pages.ProductPage{}
		productPage.Init(driver)
		err := productPage.OpenPage(product.URL)
		assert.NoError(t, err, "Не удалось открыть страницу товара")

		err = productPage.SetProductQuantity(config.QuantityProductsTen)
		assert.NoError(t, err, "Ошибка при установке количества товара")

		err = productPage.AddToCart(product.ID)
		assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

		err = productPage.IsProductInCart(product.Name, product.Price, config.QuantityProductsTen)
		assert.NoError(t, err, "Товар отсутствует в корзине или неверное количество")
	}

	runTest(t, testFunc)
}
