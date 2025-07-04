package scrape

import (
	"time"
	"time"
	"time"
	"time"
	"time"
	"context"
	"time"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/rotisserie/eris"
)

type ChromeDPWebScraperService struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func newChromeDPWebScraperService() WebScraperService {
	// Configure Chrome options for better compatibility
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	return &ChromeDPWebScraperService{
		ctx:    ctx,
		cancel: cancel,
	}
}

// Close properly cleans up ChromeDP resources
func (ws *ChromeDPWebScraperService) Close() error {
	if ws.cancel != nil {
		ws.cancel()
	}
	return nil
}

func (ws *ChromeDPWebScraperService) GetHTML(url string) (string, error) {
	ctx, cancel := context.WithTimeout(ws.ctx, 30*time.Second)
	defer cancel()

	var html string

	// Create a context with timeout derived from the service context
	ctx, cancel := context.WithTimeout(ws.ctx, 10*time.Second)
	defer cancel()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ws.ctx, 30*time.Second)
	defer cancel()

	ctx, cancel := context.WithTimeout(ws.ctx, 30*time.Second)
	defer cancel()

	// Create context with timeout to prevent hanging operations
	ctx, cancel := context.WithTimeout(ws.ctx, 30*time.Second)
	defer cancel()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ws.ctx, 30*time.Second)
	defer cancel()

	// Create context with timeout derived from service context
	ctx, cancel := context.WithTimeout(ws.ctx, 30*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.OuterHTML("body", &html, chromedp.ByQuery),
	)

	if err != nil {
		return "", eris.Wrap(err, "error retrieving html content")
	}

	return html, nil

// Close releases resources used by the ChromeDP web scraper
func (ws *ChromeDPWebScraperService) Close() error {
	if ws.cancel != nil {
		ws.cancel()
	}
	return nil
}

// Close releases resources used by the ChromeDP web scraper
func (ws *ChromeDPWebScraperService) Close() error {
	if ws.cancel != nil {
		ws.cancel()
	}
	return nil
}
}

// Close properly cleans up Chrome browser resources
func (ws *ChromeDPWebScraperService) Close() error {
	if ws.cancel != nil {
		ws.cancel()
	}
	return nil
}

func (ws *ChromeDPWebScraperService) Close() error {
	if ws.cancel != nil {
		ws.cancel()
	}
	return nil
}

// Close properly cleans up the Chrome browser process and associated resources
func (ws *ChromeDPWebScraperService) Close() error {
	if ws.cancel != nil {
		ws.cancel()
	}
	return nil
}

// Close properly cleans up the ChromeDP resources
func (ws *ChromeDPWebScraperService) Close() error {
	if ws.cancel != nil {
		ws.cancel()
	}
	return nil
}
