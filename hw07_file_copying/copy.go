package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	src, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer src.Close()

	limit, err = prepareLimit(src, offset, limit)
	if err != nil {
		return err
	}
	src.Seek(offset, 0)

	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = copyBuffer(dst, src, limit)
	if err != nil {
		return err
	}

	return nil
}

func prepareLimit(file *os.File, offset, limit int64) (int64, error) {
	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}
	size := stat.Size()
	switch {
	case size == 0:
		return 0, ErrUnsupportedFile
	case size < offset:
		return 0, ErrOffsetExceedsFileSize
	}
	if limit == 0 || limit > size-offset {
		limit = size - offset
	}
	return limit, nil
}

func copyBuffer(dst io.Writer, src io.Reader, limit int64) (written int64, err error) {
	bar := pb.StartNew(int(limit))
	tmpl := `{{ red "Copy:" }} {{ counters . }} {{ bar . "[" "-" ">" " " "]"}} {{speed . }} {{percent . | green}}`
	bar.SetTemplateString(tmpl)
	defer bar.Finish()

	src = io.LimitReader(src, limit)
	buf := make([]byte, 1)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
			bar.Increment()
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
