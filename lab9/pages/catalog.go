package pages

import (
	"fmt"
	"github.com/tebeka/selenium"
)

const (
	ProductNameInCatalog     = "//input[@id='typeahead']"
	ProductNameInProductPage = "//h3[contains(text(), '%s')]"
)

type CatalogPage struct {
	Page
}

func (c *CatalogPage) SearchProduct(text string) error {
	if err := c.typeInSearchInputSearchProduct(text); err != nil {
		return err
	}

	return c.submitSearchProductWithEnter()
}

func (c *CatalogPage) FindProduct(productName string) error {
	_, err := c.FindElement(selenium.ByXPATH, fmt.Sprintf(ProductNameInProductPage, productName))
	return err
}

func (c *CatalogPage) typeInSearchInputSearchProduct(text string) error {
	input, err := c.FindElement(selenium.ByXPATH, ProductNameInCatalog)
	if err != nil {
		return fmt.Errorf("failed to find search input: %v", err)
	}

	if err := input.Clear(); err != nil {
		return fmt.Errorf("failed to clear input: %v", err)
	}

	return input.SendKeys(text)
}

func (c *CatalogPage) submitSearchProductWithEnter() error {
	input, err := c.FindElement(selenium.ByXPATH, ProductNameInCatalog)
	if err != nil {
		return fmt.Errorf("failed to find search input: %v", err)
	}

	if err := input.SendKeys(selenium.EnterKey); err != nil {
		return fmt.Errorf("failed to press Enter: %v", err)
	}

	return nil
}
