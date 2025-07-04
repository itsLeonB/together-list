package appconstant

const (
	PromptSummarizePage = `Parse this page. Extract these two pieces of information:
1. Title: A single concise sentence without a period.
2. Short summary: One paragraph capturing the key points of the page.
Respond only with the extracted information in the specified format. Do not include the URL or any commentary. Return JSON string with "title" and "summary": %s`
)
