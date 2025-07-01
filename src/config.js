require('dotenv').config();

// Set Puppeteer's Chromium executable path to handle environments missing Chromium
process.env.PUPPETEER_EXECUTABLE_PATH = require('puppeteer').executablePath();

const requiredEnvVars = [
  'NOTION_API_KEY',
  'NOTION_DATABASE_ID',
  'MESSAGE_KEYWORD',
];

const missingVars = requiredEnvVars.filter((varName) => !process.env[varName]);
if (missingVars.length > 0) {
  throw new Error(
    `Missing required environment variables: ${missingVars.join(', ')}`
  );
}

module.exports = Object.freeze({
  NOTION: {
    API_KEY: process.env.NOTION_API_KEY,
    DATABASE_ID: process.env.NOTION_DATABASE_ID,
  },
  MESSAGE: {
    KEYWORDS: new Set(
      (process.env.MESSAGE_KEYWORD || '').split(',').filter((k) => k.trim())
    ),
    RESPONSE: {
      NO_URL: 'No URL found in the message',
      MULTIPLE_URLS: 'Multiple URLs found, saving to multiple entries...',
      ERROR: 'There was an unexpected error. Please contact developer.',
    },
  },
});
