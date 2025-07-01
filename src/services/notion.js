const { Client } = require('@notionhq/client');
const { NOTION } = require('../config');

const notion = new Client({ auth: NOTION.API_KEY });

exports.addToDatabase = async ({ message, url }) => {
  try {
    const response = await notion.pages.create({
      parent: {
        database_id: NOTION.DATABASE_ID,
      },
      properties: {
        title: {
          type: 'title',
          title: [
            {
              type: 'text',
              text: {
                content: 'pending',
              },
            },
          ],
        },
        originalMessage: {
          rich_text: [
            {
              type: 'text',
              text: {
                content: message,
              },
            },
          ],
        },
        extractedLink: {
          type: 'url',
          url,
        },
      },
    });

    return `Message saved to: ${response.url}`;
  } catch (error) {
    console.error(error);
    return `Failed to save message: ${error.message}`;
  }
};
