package archive

import (
	"archive/tar"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func archiveFatal(err error) {
	if err != nil {
		log.Fatalf("backup error: %s", err)
	}
}

func createTar(target string) (*os.File, error) {
	tarfile, err := os.Create(target)
	if err != nil {
		return nil, err
	}
	return tarfile, nil
}

func createCurrentFileArchive(path string, tarball *tar.Writer) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(tarball, file)
	return err
}

func getRootDir(info os.FileInfo, source string) string {
	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}
	return baseDir
}

func walk(source string, rootDir string, tarball *tar.Writer) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		// Walk error
		archiveFatal(err)

		header, err := tar.FileInfoHeader(info, info.Name())
		archiveFatal(err)

		if rootDir != "" {
			header.Name = filepath.Join(rootDir, strings.TrimPrefix(path, source))
		}
		tarball.WriteHeader(header)
		archiveFatal(err)

		if info.IsDir() {
			return nil
		}

		err = createCurrentFileArchive(path, tarball)
		archiveFatal(err)
		return err
	}
}

func untarEntries(tarReader *tar.Reader, target string) error {
	for {
		// Read Tar entry
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Create corresponding folder
		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		// If file then untar
		err = untarEntry(tarReader, path, info)
		if err != nil {
			return err
		}
	}
	return nil
}

func untarEntry(tarReader *tar.Reader, path string, info fs.FileInfo) error {
	perms := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	file, err := os.OpenFile(path, perms, info.Mode())
	if err != nil {
		return err
	}

	// Populate with untarred file
	_, err = io.Copy(file, tarReader)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func TarFolder(source string) error {
	filename := filepath.Base(source)
	target := fmt.Sprintf("%s.tar", filename)

	tarfile, err := createTar(target)
	archiveFatal(err)
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(source)
	archiveFatal(err)

	rootDir := getRootDir(info, source)
	return filepath.Walk(source, walk(source, rootDir, tarball))
}

func UntarFolder(tarball, target string) error {
	reader, err := os.Open(tarball)
	archiveFatal(err)
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	return untarEntries(tarReader, target)
}
