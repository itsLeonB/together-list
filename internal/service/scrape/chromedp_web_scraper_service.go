package scrape

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/itsLeonB/together-list/internal/logging"
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
	ctx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(logging.Infof))

	return &ChromeDPWebScraperService{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (ws *ChromeDPWebScraperService) GetHTML(url string) (string, error) {
	var html string

	err := chromedp.Run(ws.ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.OuterHTML("body", &html, chromedp.ByQuery),
	)

	if err != nil {
		return "", eris.Wrap(err, "error retrieving html content")
	}

	return html, nil
}
