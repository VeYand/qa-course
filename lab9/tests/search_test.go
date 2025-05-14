package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"lab9/config"
	"lab9/pages"
	"testing"
)

func TestSearchInMainPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := pages.CatalogPage{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage("")
		assert.NoError(t, err, "Не удалось открыть главную страницу")

		err = catalogPage.SearchProduct(config.ProductNameCasio)
		assert.NoError(t, err, "Ошибка при поиске товара Casio")

		err = catalogPage.FindProduct(config.ProductNameCasio)
		assert.NoError(t, err, "Товар Casio не найден на странице")
	}

	runTest(t, testFunc)
}

func TestSearchInProductPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := pages.CatalogPage{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage(config.ProductPageURL)
		assert.NoError(t, err, "Не удалось открыть страницу продукта")

		err = catalogPage.SearchProduct(config.ProductNameRoyal)
		assert.NoError(t, err, "Ошибка при поиске товара Royal")

		err = catalogPage.FindProduct(config.ProductNameRoyal)
		assert.NoError(t, err, "Товар Royal не найден на странице продукта")
	}

	runTest(t, testFunc)
}

func TestSearchInCategoryPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := pages.CatalogPage{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage(config.CategoryPageURL)
		assert.NoError(t, err, "Не удалось открыть страницу категории")

		err = catalogPage.SearchProduct(config.ProductNameRoyal)
		assert.NoError(t, err, "Ошибка при поиске товара Royal в категории")

		err = catalogPage.FindProduct(config.ProductNameRoyal)
		assert.NoError(t, err, "Товар Royal не найден в категории")
	}

	runTest(t, testFunc)
}

func TestSearchInSearchPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := pages.CatalogPage{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage(config.SearchPageURL)
		assert.NoError(t, err, "Не удалось открыть страницу поиска")

		err = catalogPage.SearchProduct(config.ProductNameCitizen)
		assert.NoError(t, err, "Ошибка при поиске товара Citizen")

		err = catalogPage.FindProduct(config.ProductNameCitizen)
		assert.NoError(t, err, "Товар Citizen не найден на странице поиска")
	}

	runTest(t, testFunc)
}
