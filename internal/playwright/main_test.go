//go:build integration

package playwright

import (
	"log"
	"testing"

	"github.com/playwright-community/playwright-go"
)

// https://github.com/playwright-community/playwright-go/blob/main/examples/screenshot/main.go
func TestScreenshot(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		t.Fatal(err)
	}
	if _, err = page.Goto("https://news.ycombinator.com/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		t.Fatalf("could not goto: %v", err)
	}
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("screenshot.png"),
	}); err != nil {
		t.Fatalf("could not create screenshot: %v", err)
	}
	if err = browser.Close(); err != nil {
		t.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		t.Fatalf("could not stop Playwright: %v", err)
	}
}
