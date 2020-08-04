package main

import (
	"errors"
	"io"
	"os"

	"github.com/mitchellh/ioprogress"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrThisIsDirectory       = errors.New("this is directory")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrNoInFile              = errors.New("не указан исходный файл или он не существует")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	fIn, err := os.Open(fromPath)
	defer func() {
		if fIn.Close() != nil {
			return
		}
	}()
	if err != nil {
		return ErrNoInFile
	}

	fInInfo, err := fIn.Stat()
	if err != nil {
		return err
	}
	fInMode := fInInfo.Mode()
	if fInMode.IsDir() {
		return ErrThisIsDirectory
	}
	if !fInMode.IsRegular() {
		return ErrUnsupportedFile
	}

	fInSize := fInInfo.Size()

	switch {
	case offset >= fInSize:
		return ErrOffsetExceedsFileSize
	case limit == 0:
		limit = fInSize - offset
	case fInSize < limit+offset:
		limit = fInSize - offset
	default:
	}

	fOut, err := os.Create(toPath)
	defer func() {
		if fOut.Close() != nil {
			return
		}
	}()
	if err != nil {
		return err
	}

	pg := &ioprogress.Reader{
		Reader:   fIn,
		Size:     limit,
		DrawFunc: ioprogress.DrawTerminalf(os.Stdout, ioprogress.DrawTextFormatBar(100)),
	}

	if _, err = fIn.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	if _, err = io.CopyN(fOut, pg, limit); err != nil {
		return err
	}
	return nil
}
