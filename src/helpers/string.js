exports.splitFirstLine = (inputText) => {
  const firstNewlineIndex = inputText.indexOf('\n');

  if (firstNewlineIndex === -1) {
    return [inputText, ''];
  }

  return [
    inputText.slice(0, firstNewlineIndex),
    inputText.slice(firstNewlineIndex + 1),
  ];
};
