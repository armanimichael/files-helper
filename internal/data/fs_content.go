package data

func CreateSamplePlaintextFiles(maxParagraphs int) {
	createSampleFiles(maxParagraphs, "plain-%02d.txt", GetSamplePlaintext)
}

func CreateSampleHtmlFiles(maxParagraphs int) {
	createSampleFiles(maxParagraphs, "html-%02d.txt", GetSampleHtmltext)
}
