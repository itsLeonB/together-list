const notionService = require('./src/services/notion');
const whatsappService = require('./src/services/whatsapp');
const PostgresStore = require('./src/services/postgres');
const urlHelper = require('./src/helpers/url');
const stringHelper = require('./src/helpers/string');
const { MESSAGE } = require('./src/config');

async function main() {
  const store = new PostgresStore();
  await store.init();

  const messageHandlerFunc = async (message) => {
    try {
      const inputs = stringHelper.splitFirstLine(message.body);

      if (inputs.length === 0 || !MESSAGE.KEYWORDS.has(inputs[0])) {
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

      const responses = await Promise.allSettled(
        urls.map((url) =>
          notionService.addToDatabase({
            message: originalMessage,
            url,
          })
        )
      );

      const successfulResponses = responses
        .filter((result) => result.status === 'fulfilled')
        .map((result) => result.value);

      const failedCount = responses.length - successfulResponses.length;

      let replyMessage = successfulResponses.join('\n\n');
      if (failedCount > 0) {
        replyMessage += `\n\n⚠️ ${failedCount} URL(s) failed to save.`;
      }

      message.reply(replyMessage);
    } catch (err) {
      console.error(err);
      message.reply(MESSAGE.RESPONSE.ERROR);
    }
  };

  whatsappService.setupAndRun({ messageHandlerFunc, store });
}

main().catch(console.error);
