# Together List

A Go-based application that automatically scrapes web content, processes it with AI, and manages lists through Notion integration. The application supports both WhatsApp bot functionality and scheduled job processing.

## Features

- **Web Scraping**: Extract content from web pages using Colly or ChromeDP
- **AI Processing**: Summarize and process content using Google Gemini or OpenRouter LLMs
- **Notion Integration**: Automatically manage and update Notion databases
- **WhatsApp Bot**: Interactive bot for processing URLs via WhatsApp
- **Scheduled Jobs**: Background processing for automated tasks
- **Multi-Provider Support**: Flexible LLM provider configuration

## Architecture

The project follows a clean architecture pattern with the following structure:

```
├── cmd/                    # Application entry points
│   ├── whatsapp/          # WhatsApp bot application
│   └── job/               # Job scheduler application
├── internal/              # Internal packages
│   ├── config/           # Configuration management
│   ├── delivery/         # Delivery layer (WhatsApp, Jobs)
│   ├── entity/           # Domain entities
│   ├── repository/       # Data access layer
│   ├── service/          # Business logic
│   │   ├── llm/         # LLM service implementations
│   │   └── scrape/      # Web scraping services
│   ├── dto/             # Data transfer objects
│   ├── provider/        # External service providers
│   └── util/            # Utility functions
└── ci/                   # CI/CD configurations
```

## Prerequisites

- Go 1.24.4 or higher
- PostgreSQL database
- Notion API access
- LLM API keys (Google AI or OpenRouter)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/itsLeonB/together-list.git
cd together-list
```

2. Install dependencies:
```bash
go mod download
```

3. Copy and configure environment variables:
```bash
cp .env.example .env
```

4. Edit `.env` with your configuration values.

## Configuration

The application uses environment variables for configuration. Key settings include:

### Database
- `DATABASE_URL`: PostgreSQL connection string

### Notion
- `NOTION_API_KEY`: Your Notion integration API key
- `NOTION_DATABASE_ID`: Target Notion database ID

### LLM Providers
- `LLM_PROVIDERS`: Comma-separated list of providers (google,openrouter)
- `GOOGLE_LLM_API_KEY`: Google AI API key
- `GOOGLE_LLM_MODEL`: Google AI model name
- `OPEN_ROUTER_API_KEY`: OpenRouter API key
- `OPEN_ROUTER_MODEL`: OpenRouter model name

### Web Scraping
- `WEB_SCRAPER`: Scraper type (colly,chromedp)

### Application
- `SERVICE_TYPE`: Application mode (whatsapp,job)
- `MESSAGE_KEYWORD`: WhatsApp trigger keyword
- `JOB_NAME`: Job type to execute
- `ENV`: Environment (debug,production)
- `TIMEZONE`: Application timezone

## Usage

### WhatsApp Bot

Run the WhatsApp bot:
```bash
make whatsapp
# or
go run cmd/whatsapp/main.go
```

The bot will:
1. Display a QR code for WhatsApp Web authentication
2. Listen for messages containing the configured keyword
3. Process URLs sent in messages
4. Scrape content and generate AI summaries
5. Update Notion database with results

### Job Scheduler

Run scheduled jobs:
```bash
make job
# or
go run cmd/job/main.go
```

Available job types:
- `Summarize`: Process and summarize existing content

### Development

For development with hot reload:
```bash
make hotreload
```

Run tests:
```bash
make test
```

Run linting:
```bash
make lint
```

## Supported LLM Providers

### Google AI (Gemini)
- Models: gemini-2.0-flash-lite, gemini-pro, etc.
- Requires Google AI API key

### OpenRouter
- Access to various models including Gemini, GPT, Claude
- Supports free tier models
- Requires OpenRouter API key

## Web Scraping Options

### Colly
- Fast, lightweight scraper
- Good for simple HTML parsing
- Lower resource usage

### ChromeDP
- Full browser automation
- Handles JavaScript-heavy sites
- Higher resource usage but more reliable

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Ellion Blessan** - [itsLeonB](https://github.com/itsLeonB)

## Support

For issues and questions, please use the GitHub issue tracker.
