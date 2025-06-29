const { v4: uuidv4 } = require('uuid');
const { writeFileSync } = require('fs');
const Tiktok = require("@tobyg74/tiktok-api-dl");

exports.download = async ({ url }) => {
  try {
    const uuid = uuidv4();
    const outputFile = `${uuid}.mp4`

    const response = await Tiktok.Downloader(url, {
      version: 'v3'
    });

    // const buffer = response.result.videoSD;

    // writeFileSync(outputFile, buffer);

    // return outputFile;
    return response.result.videoSD;
  } catch (err) {
    console.error('Failed to download:', err);
  }
}
