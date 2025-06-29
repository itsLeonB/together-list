const { Client } = require('@notionhq/client');
const { NOTION } = require('./secrets');
const tiktok = require('./tiktok');

const notion = new Client({ auth: NOTION.API_KEY });

exports.addToDatabase = async ({ type, message }) => {
  try {
    const videoFile = await tiktok.download({ url: message });
    if (!videoFile) {
      throw new Error('Failed to download video');
    }

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
        type: {
          select: {
            name: type,
          },
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
        videoFile: {
          rich_text: [
            {
              type: 'text',
              text: {
                content: videoFile,
              },
            },
          ],
        },
      },
    });

    return `Message saved to: ${response.url}`;
  } catch (error) {
    console.error(error);
    return `Failed to save message: ${error.message}`;
  }
};
