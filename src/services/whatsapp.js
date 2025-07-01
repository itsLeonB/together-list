const { Client, LocalAuth } = require('whatsapp-web.js');
const qrcode = require('qrcode-terminal');

exports.setupAndRun = ({ messageHandlerFunc }) => {
  const client = new Client({
    authStrategy: new LocalAuth(),
  });

  client.on('ready', () => {
    console.log('Client is ready!');
  });

  client.on('qr', (qr) => {
    qrcode.generate(qr, { small: true });
  });

  client.on('message_create', messageHandlerFunc);

  client.initialize();
};
