require('dotenv').config();

module.exports = Object.freeze({
  NOTION: {
    API_KEY: process.env.NOTION_API_KEY,
    DATABASE_ID: process.env.NOTION_DATABASE_ID,
  },
  MESSAGE: {
    KEYWORDS: new Set((process.env.MESSAGE_KEYWORD || '').split(',')),
    RESPONSE: {
      NO_URL: 'No URL found in the message',
      MULTIPLE_URLS: 'Multiple URLs found, saving to multiple entries...',
    },
  },
});
