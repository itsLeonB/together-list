const notionService = require('./src/services/notion');
const whatsappService = require('./src/services/whatsapp');
const urlHelper = require('./src/helpers/url');
const stringHelper = require('./src/helpers/string');
const { MESSAGE } = require('./src/config');

const messageHandlerFunc = async (message) => {
  const inputs = stringHelper.splitFirstLine(message.body);

  if (!MESSAGE.KEYWORDS.has(inputs[0])) {
    return;
  }

  console.info(`Handling message from: ${message.from}`);
  console.info(`Device: ${message.deviceType}`);
  console.info(`Full text: ${message.body}`);

  const originalMessage = inputs[1];
  const urls = urlHelper.extractUrls({ text: originalMessage });

  if (urls.length === 0) {
    message.reply(MESSAGE.RESPONSE.NO_URL);
    return;
  }

  if (urls.length === 1) {
    const responseMsg = await notionService.addToDatabase({
      message: originalMessage,
      url: urls[0],
    });

    message.reply(responseMsg);
    return;
  }

  message.reply(MESSAGE.RESPONSE.MULTIPLE_URLS);

  const responses = await Promise.all(
    urls.map((url) =>
      notionService.addToDatabase({
        message: originalMessage,
        url,
      })
    )
  );

  message.reply(responses.join('\n\n'));
};

whatsappService.setupAndRun({ messageHandlerFunc });
