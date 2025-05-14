package pages

import (
	"fmt"
	"github.com/tebeka/selenium"
	"lab9/config"
	"time"
)

const (
	AddToCartButtonID = "productAdd"
	OrderButton       = "a[href='cart/view']"
	LoginField        = "login"
	PasswordField     = "password"
	NameField         = "name"
	EmailField        = "email"
	AddressField      = "address"
	NoteField         = "//textarea[@name='note']"

	SubmitButton = "//button[contains(text(), 'Оформить')]"

	ErrorMessage = ".alert-danger"
	ErrorTitle   = "//h1[text()='Произошла ошибка']"
)

type OrderPage struct {
	Page
}

func (o *OrderPage) AddToCart() error {
	elem, err := o.FindElement(selenium.ByID, AddToCartButtonID)
	if err != nil {
		return err
	}

	return elem.Click()
}

func (o *OrderPage) ClickOrderButton() error {
	time.Sleep(time.Millisecond * 2000)

	orderButton, err := o.FindElement(selenium.ByCSSSelector, OrderButton)
	if err != nil {
		return err
	}

	return orderButton.Click()
}

func (o *OrderPage) FillOrderForm(note string) error {
	if err := o.fillField(selenium.ByXPATH, NoteField, note); err != nil {
		return err
	}

	return o.submit()
}

func (o *OrderPage) FillFullOrderForm(formData config.Order) error {
	if err := o.fillField(selenium.ByName, LoginField, formData.Login); err != nil {
		return err
	}
	if err := o.fillField(selenium.ByName, PasswordField, formData.Password); err != nil {
		return err
	}
	if err := o.fillField(selenium.ByName, NameField, formData.Name); err != nil {
		return err
	}
	if err := o.fillField(selenium.ByName, EmailField, formData.Email); err != nil {
		return err
	}
	if err := o.fillField(selenium.ByName, AddressField, formData.Address); err != nil {
		return err
	}
	if err := o.fillField(selenium.ByXPATH, NoteField, formData.Note); err != nil {
		return err
	}

	return o.submit()
}

func (o *OrderPage) IsOrderMadeSuccessful() (bool, error) {
	unsuccessfull, err := o.IsOrderMadeUnsuccessful()
	if err != nil {
		return false, err
	}

	return !unsuccessfull, nil
}

func (o *OrderPage) IsOrderMadeUnsuccessful() (bool, error) {
	elem, err := o.FindElement(selenium.ByXPATH, ErrorTitle)
	if err != nil {
		return false, err
	}

	return elem.IsDisplayed()
}

func (o *OrderPage) IsOrderMadeFailed() (bool, error) {
	elem, err := o.FindElement(selenium.ByCSSSelector, ErrorMessage)
	if err != nil {
		return false, err
	}

	return elem.IsDisplayed()
}

func (o *OrderPage) fillField(by, selector, text string) error {
	input, err := o.FindElement(by, selector)
	if err != nil {
		return fmt.Errorf("failed to find search input: %v", err)
	}

	if err := input.Clear(); err != nil {
		return fmt.Errorf("failed to clear input: %v", err)
	}

	return input.SendKeys(text)
}

func (o *OrderPage) submit() error {
	input, err := o.FindElement(selenium.ByXPATH, SubmitButton)
	if err != nil {
		return fmt.Errorf("failed to find search input: %v", err)
	}

	if err := input.SendKeys(selenium.EnterKey); err != nil {
		return fmt.Errorf("failed to press Enter: %v", err)
	}

	return nil
}
