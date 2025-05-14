package tests

import (
	"lab9/pages"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"lab9/config"
)

func TestSuccessfulAuth(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		cfg := config.GetExistingCredentials()

		authPage := pages.AuthorizationPage{}
		authPage.Init(driver)
		err := authPage.OpenPage(config.LoginPageURL)
		assert.NoError(t, err, "Не удалось открыть страницу авторизации")

		err = authPage.Login(cfg.Login, cfg.Password)
		assert.NoError(t, err, "Ошибка при вводе существующих учетных данных")

		isLoginSuccessful, err := authPage.IsLoginSuccessful()
		assert.NoError(t, err, "Ошибка при проверке успешной авторизации")
		assert.True(t, isLoginSuccessful, "Ожидалась успешная авторизация")
	}

	runTest(t, testFunc)
}

func TestFailedAuth(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		cfg := config.GetNonExistingCredentials()

		authPage := pages.AuthorizationPage{}
		authPage.Init(driver)
		err := authPage.OpenPage(config.LoginPageURL)
		assert.NoError(t, err, "Не удалось открыть страницу авторизации")

		err = authPage.Login(cfg.Login, cfg.Password)
		assert.NoError(t, err, "Ошибка при вводе невалидных учетных данных")

		isLoginFailed, err := authPage.IsLoginError()
		assert.NoError(t, err, "Ошибка при проверке сообщения об ошибке")
		assert.True(t, isLoginFailed, "Ожидалась ошибка авторизации")
	}

	runTest(t, testFunc)
}
