package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/schollz/progressbar"
	"io"
)

type ProgressReporter struct {
	r         io.Reader
	w         io.Writer
	progress  uint64
	totalSize uint64
	label     string
	bar       *progressbar.ProgressBar
}

func (pr *ProgressReporter) initProgressBar() {
	pr.bar = progressbar.NewOptions64(int64(pr.totalSize),
		progressbar.OptionSetTheme(progressbar.Theme{Saucer: "=", SaucerPadding: "-", BarStart: "[", BarEnd: "]"}),
		progressbar.OptionSetWidth(45),
		)
}

func (pr *ProgressReporter) Read(p []byte) (int, error) {
	n, err := pr.r.Read(p)
	pr.progress += uint64(n)
	pr.report()
	return n, err
}

func (pr *ProgressReporter) Write(p []byte) (int, error) {
	n, err := pr.w.Write(p)
	pr.progress += uint64(n)
	pr.report()
	return n, err
}

func (pr *ProgressReporter) report() {

	if pr.totalSize > 0 {
		if pr.bar == nil {
			pr.initProgressBar()
		}
		pr.bar.Set64(int64(pr.progress))
	} else {
		fmt.Printf("\r%s... %s complete\r", pr.label, humanize.Bytes(uint64(pr.progress)))
	}
}
