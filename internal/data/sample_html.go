package data

func CreateSampleHtmlFiles(maxParagraphs int) {
	createSampleFiles(maxParagraphs, htmlFileNameFormat, GetSampleHtmltext)
}

func ReadSampleHtml(paragraphs int) (content []byte, err error) {
	return readSampleContent(paragraphs, htmlFileNameFormat)
}
