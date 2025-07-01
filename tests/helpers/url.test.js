const { extractUrls } = require('../../src/helpers/url');

describe('extractUrls', () => {
  it('extracts multiple URLs from multiline text', () => {
    const input = {
      text: `
http://www.google.com
http://www.notion.com`,
    };

    const result = extractUrls(input);

    expect(result).toEqual(['http://www.google.com', 'http://www.notion.com']);
  });

  it('returns empty array if no URL is found', () => {
    const input = {
      text: `no links here`,
    };

    const result = extractUrls(input);

    expect(result).toEqual([]);
  });

  it('extracts single URL', () => {
    const input = {
      text: `Visit https://openai.com for info`,
    };

    const result = extractUrls(input);

    expect(result).toEqual(['https://openai.com']);
  });
});
