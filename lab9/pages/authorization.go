package pages

import "github.com/tebeka/selenium"

const (
	loginField     = "login"
	passwordField  = "password"
	submitButton   = ".btn-default"
	successMessage = ".alert-success"
	errorMessage   = ".alert-danger"
)

type AuthorizationPage struct {
	Page
}

func (a *AuthorizationPage) Login(login, password string) error {
	loginElem, err := a.FindElement(selenium.ByName, loginField)
	if err != nil {
		return err
	}
	if err := loginElem.SendKeys(login); err != nil {
		return err
	}

	passElem, err := a.FindElement(selenium.ByName, passwordField)
	if err != nil {
		return err
	}
	if err := passElem.SendKeys(password); err != nil {
		return err
	}

	submitElem, err := a.FindElement(selenium.ByCSSSelector, submitButton)
	if err != nil {
		return err
	}

	return submitElem.Click()
}

func (a *AuthorizationPage) IsLoginSuccessful() (bool, error) {
	elem, err := a.FindElement(selenium.ByCSSSelector, successMessage)
	if err != nil {
		return false, err
	}

	return elem.IsDisplayed()
}

func (a *AuthorizationPage) IsLoginError() (bool, error) {
	elem, err := a.FindElement(selenium.ByCSSSelector, errorMessage)
	if err != nil {
		return false, err
	}

	return elem.IsDisplayed()
}
