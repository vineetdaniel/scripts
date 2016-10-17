package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type Archiver interface {
	DestFmt() string
	Archive(src, dest string) error
}

type zipper struct{}

/* we are defining a variable called ZIP of type Archiver, so from outside the package
its very clear that we can use that variable wherever archiver is needed. we assign it
with nil cast to the type *zipper. */
var ZIP Archiver = (*zipper)(nil)

func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err

	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer out.Close()
	w := zip.NewWriter(out)
	defer w.Close()
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil //skip
		}
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		io.Copy(f, in)
		return nil
	})
}

func (z *zipper) DestFmt() string {
	return "%d.zip"
}
