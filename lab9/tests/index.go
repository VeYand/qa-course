package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"testing"
)

func runTest(t *testing.T, testFunc func(*testing.T, selenium.WebDriver)) {
	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func runTestForBrowser(t *testing.T, browserName string, testFunc func(*testing.T, selenium.WebDriver)) {
	t.Helper()
	t.Run(browserName, func(t *testing.T) {
		caps := selenium.Capabilities{"browserName": browserName}
		driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
		if !assert.NoError(t, err, "Failed to start "+browserName+" session") {
			return
		}
		defer driver.Quit()

		testFunc(t, driver)
	})
}
