const { splitFirstLine } = require('../../src/helpers/string');

describe('splitFirstLine', () => {
  it('splits first line from the rest', () => {
    const input = `beasiswa
http://www.google.com
http://www.notion.com`;

    const result = splitFirstLine(input);

    expect(result).toEqual([
      'beasiswa',
      'http://www.google.com\nhttp://www.notion.com',
    ]);
  });

  it('returns entire string as first item if no newline', () => {
    const input = 'single line input';

    const result = splitFirstLine(input);

    expect(result).toEqual(['single line input', '']);
  });

  it('handles empty string', () => {
    const input = '';

    const result = splitFirstLine(input);

    expect(result).toEqual(['', '']);
  });

  it('handles string that starts with newline', () => {
    const input = `
http://only.com`;

    const result = splitFirstLine(input);

    expect(result).toEqual(['', 'http://only.com']);
  });
});
