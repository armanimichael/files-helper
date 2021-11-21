package data

func CreateSamplePlaintextFiles(maxParagraphs int) {
	createSampleFiles(maxParagraphs, plainFileNameFormat, GetSamplePlaintext)
}

func ReadSamplePlaintext(paragraphs int) (content []byte, err error) {
	return readSampleContent(paragraphs, plainFileNameFormat)
}
