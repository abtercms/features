// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"testing"
	"github.com/stretchr/testify/assert"
)

const baseUrl = "https://abtercms.test/"

func TestHelloWorld(t *testing.T) {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(baseUrl),
		// wait for footer element is visible (ie, page is loaded)
		//chromedp.WaitVisible(`.container`),
		// retrieve the value of the textarea
		chromedp.Text(`.container`, &example),
	)

	assert.Nil(t, err)

	// assert equality
	assert.Equal(t, "Hello, World!", example, "they should be equal")
}
