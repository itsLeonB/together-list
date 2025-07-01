exports.extractUrls = ({ text }) => {
  return text.match(/(https?:\/\/[^\s]+)/gi) ?? [];
};
