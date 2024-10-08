//go:build integration

package playwright

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/gosimple/slug"
	"github.com/playwright-community/playwright-go"
)

type Browser int

func init() {
	err := playwright.Install()
	if err != nil {
		log.Fatalf("Error installing playwright: %v", err)
	}
}

const (
	Chromium Browser = iota
	Firefox
	WebKit
)

// https://github.com/playwright-community/playwright-go/blob/main/examples/screenshot/main.go
func TestScreenshot(t *testing.T) {
	f := func(ssURL string, browserType Browser) {
		t.Helper()

		pw, err := playwright.Run()
		if err != nil {
			t.Fatalf("could not launch playwright: %v", err)
		}

		var browser playwright.Browser
		switch browserType {
		case Chromium:
			browser, err = pw.Chromium.Launch()
		case Firefox:
			browser, err = pw.Firefox.Launch()
		case WebKit:
			browser, err = pw.WebKit.Launch()
		}

		if err != nil {
			t.Fatalf("could not launch browser: %v", err)
		}

		page, err := browser.NewPage()
		if err != nil {
			t.Fatal(err)
		}
		if _, err = page.Goto(ssURL, playwright.PageGotoOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		}); err != nil {
			t.Fatalf("could not goto: %v", err)
		}

		parsedURL, err := url.Parse(ssURL)
		if err != nil {
			t.Fatal(err)
		}

		urlSlug := slug.Make(parsedURL.Host + parsedURL.Path)
		ssFileName := fmt.Sprintf("%s.png", urlSlug)

		if _, err = page.Screenshot(playwright.PageScreenshotOptions{
			Path: playwright.String(ssFileName),
		}); err != nil {
			t.Fatalf("could not create screenshot: %v", err)
		}

		t.Logf("SCREENSHOT %s", ssFileName)

		if err = browser.Close(); err != nil {
			t.Fatalf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			t.Fatalf("could not stop Playwright: %v", err)
		}
	}

	f("https://news.ycombinator.com/", Chromium)
	f("https://news.google.com/", Firefox)
	f("https://github.com/playwright-community/playwright-go", WebKit)
}
