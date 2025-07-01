const { Client, LocalAuth } = require('whatsapp-web.js');
const qrcode = require('qrcode-terminal');

exports.setupAndRun = ({ messageHandlerFunc }) => {
  const client = new Client({
    authStrategy: new LocalAuth(),
  });

  client.on('ready', () => {
    console.info('Client is ready!');
  });

  client.on('auth_failure', (msg) => {
    console.error('Authentication failed:', msg);
    process.exit(1);
  });

  client.on('disconnected', (reason) => {
    console.error('Client was logged out:', reason);
    process.exit(1);
  });

  client.on('qr', (qr) => {
    qrcode.generate(qr, { small: true });
  });

  client.on('message_create', (message) => {
    if (!message.fromMe) {
      messageHandlerFunc(message);
    }
  });

  client.initialize();
};
