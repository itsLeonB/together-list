const { Client, LocalAuth } = require('whatsapp-web.js');
const qrcode = require('qrcode-terminal');
const notion = require('./notion');

const client = new Client({
	authStrategy: new LocalAuth(),
});

client.on('ready', () => {
	console.log('Client is ready!');
});

client.on('qr', qr => {
	qrcode.generate(qr, { small: true });
});

client.on('message_create', async (message) => {
	const inputMessage = message.body;
	const inputs = inputMessage.split('\n');
	if (inputs[0] === 'add') {
		const responseMsg = await notion.addToDatabase({
			type: inputs[1],
			message: inputs[2],
		});

		message.reply(responseMsg);
	}
});

module.exports = client;
