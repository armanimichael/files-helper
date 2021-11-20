package data

func CreateSampleHtmlFiles(maxParagraphs int) {
	createSampleFiles(maxParagraphs, htmlFileNameFormat, GetSampleHtmltext)
}
