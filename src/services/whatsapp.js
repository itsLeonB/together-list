const fs = require('fs');
const path = require('path');
const qrcode = require('qrcode-terminal');
const { Client, RemoteAuth } = require('whatsapp-web.js');

exports.setupAndRun = ({ messageHandlerFunc, store }) => {
  const client = new Client({
    authStrategy: new RemoteAuth({
      store,
      backupSyncIntervalMs: 300000,
    }),
    puppeteer: { headless: true },
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

  client.on('qr', async (qr) => {
    qrcode.generate(qr, { small: true });
  });

  client.on('message_create', (message) => {
    if (!message.fromMe) {
      messageHandlerFunc(message);
    }
  });

  cleanChromeLock();

  client.initialize();
};

const cleanChromeLock = () => {
  const profilePath = path.join('.wwebjs_auth', 'session');
  const lockFiles = ['SingletonLock', 'SingletonSocket', 'SingletonCookie'];

  for (const file of lockFiles) {
    const filePath = path.join(profilePath, file);
    if (fs.existsSync(filePath)) {
      fs.unlinkSync(filePath);
      console.info(`🧹 Deleted lock file: ${file}`);
    }
  }
};
