package compression

import (
	"compress/gzip"
	"io"
)

// Gzip reads from a source and writes the gzipped result to the target
func Gzip(header gzip.Header, target io.Writer, source io.Reader) error {
	compressor := gzip.NewWriter(target)
	compressor.Header = header
	defer compressor.Close()

	_, err := io.Copy(compressor, source)
	return err
}

// UnGzip reads from a source and writes the un-gzipped result to the target
// returns any error or the header
func UnGzip(target io.Writer, source io.Reader) (header gzip.Header, err error) {
	decompressor, err := gzip.NewReader(source)
	if err != nil {
		return gzip.Header{}, err
	}
	header = decompressor.Header
	defer decompressor.Close()

	_, err = io.Copy(target, decompressor)
	return header, err
}
