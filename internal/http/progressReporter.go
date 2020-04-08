package http

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v2"
	"io"
)

//
// ProgressReporter is used to display download progress
//
type progressReporter struct {
	r         io.Reader
	w         io.Writer
	progress  int64
	totalSize uint64
	label     string
	bar       *progressbar.ProgressBar
}

func (pr *progressReporter) initProgressBar() {
	pr.bar = progressbar.NewOptions64(int64(pr.totalSize),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription(pr.label),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerPadding: "-",
			SaucerHead:    "[green]>[reset]",
			BarStart:      "[", BarEnd: "]",
		}),
		progressbar.OptionSetWidth(45),
	)
}

func (pr *progressReporter) Read(p []byte) (int, error) {
	n, err := pr.r.Read(p)
	pr.progress += int64(n)
	pr.report(int64(n))
	return n, err
}

func (pr *progressReporter) Write(p []byte) (int, error) {
	//fmt.Printf("sdsa %v", pr.totalSize)
	n, err := pr.w.Write(p)
	pr.report(int64(n))
	return n, err
}

func (pr *progressReporter) report(progress int64) {
	if pr.totalSize > 0 {
		if pr.bar == nil {
			pr.initProgressBar()
		}
		pr.bar.Add64(progress) // nolint: errcheck
	} else {
		fmt.Printf("\r%s... %s complete\r", pr.label, humanize.Bytes(uint64(pr.progress)))
	}
}
