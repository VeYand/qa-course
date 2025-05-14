package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"lab9/config"
	"lab9/pages"
	"testing"
)

func TestMadeOrderLoggedSuccessful(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		cfg := config.GetExistingCredentials()

		authPage := pages.AuthorizationPage{}
		authPage.Init(driver)
		err := authPage.OpenPage(config.LoginPageURL)
		assert.NoError(t, err, "Не удалось открыть страницу авторизации")

		err = authPage.Login(cfg.Login, cfg.Password)
		assert.NoError(t, err, "Ошибка при вводе логина/пароля")

		isLoginSuccessful, err := authPage.IsLoginSuccessful()
		assert.NoError(t, err, "Ошибка при проверке авторизации")
		assert.True(t, isLoginSuccessful, "Авторизация не прошла успешно")

		orderPage := pages.OrderPage{}
		orderPage.Init(driver)
		err = orderPage.OpenPage(config.CitizenProduct.URL)
		assert.NoError(t, err, "Не удалось открыть страницу товара")

		err = orderPage.AddToCart()
		assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

		err = orderPage.ClickOrderButton()
		assert.NoError(t, err, "Не удалось перейти к оформлению заказа")

		err = orderPage.FillOrderForm(config.ExistingUserToOrderData.Note)
		assert.NoError(t, err, "Ошибка при заполнении заметки к заказу")

		isSuccess, err := orderPage.IsOrderMadeSuccessful()
		assert.NoError(t, err, "Ошибка при проверке успешности заказа")
		assert.True(t, isSuccess, "Заказ не был оформлен успешно")
	}

	runTest(t, testFunc)
}

func TestMadeOrderSuccessful(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		orderPage := pages.OrderPage{}
		orderPage.Init(driver)
		err := orderPage.OpenPage(config.CitizenProduct.URL)
		assert.NoError(t, err, "Не удалось открыть страницу товара")

		err = orderPage.AddToCart()
		assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

		err = orderPage.ClickOrderButton()
		assert.NoError(t, err, "Не удалось перейти к оформлению заказа")

		err = orderPage.FillFullOrderForm(config.NonExistingUserToOrderData)
		assert.NoError(t, err, "Ошибка при заполнении формы заказа")

		isSuccess, err := orderPage.IsOrderMadeSuccessful()
		assert.NoError(t, err, "Ошибка при проверке успешности заказа")
		assert.True(t, isSuccess, "Заказ не был оформлен успешно")
	}

	runTest(t, testFunc)
}

func TestMadeOrderFailed(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		orderPage := pages.OrderPage{}
		orderPage.Init(driver)
		err := orderPage.OpenPage(config.CitizenProduct.URL)
		assert.NoError(t, err, "Не удалось открыть страницу товара")

		err = orderPage.AddToCart()
		assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

		err = orderPage.ClickOrderButton()
		assert.NoError(t, err, "Не удалось перейти к оформлению заказа")

		err = orderPage.FillFullOrderForm(config.ExistingUserToOrderData)
		assert.NoError(t, err, "Ошибка при заполнении формы заказа")

		isFailed, err := orderPage.IsOrderMadeFailed()
		assert.NoError(t, err, "Ошибка при проверке неудачного заказа")
		assert.True(t, isFailed, "Ожидалась ошибка оформления")
	}

	runTest(t, testFunc)
}
