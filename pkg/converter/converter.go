package converter

import "C"
import (
	"batchConvert/pkg/command"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type Converter struct {
	Input  string
	Output string
	Config Config
	DryRun bool
}

func NewConverter(input, output string, param Config, dryRun bool) *Converter {
	return &Converter{
		Input:  input,
		Output: output,
		Config: param,
		DryRun: dryRun,
	}
}

func (c Converter) Convert() error {
	// check input path type
	info, err := os.Stat(c.Input)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return c.batchConvert()
	} else {
		return c.convert(c.Input, c.Output)
	}
}

func (c Converter) batchConvert() error {
	log.Printf("[batch convert] Reading from path: %s", c.Input)
	var videos []string
	// iterate the folder, add video to videos
	err := filepath.WalkDir(c.Input, func(path string, d fs.DirEntry, err error) error {
		// skip dir
		if !d.IsDir() {
			// check file extension
			if VideoExtension.Match(path) {
				videos = append(videos, path)
			}
		}
		return nil
	})
	for i, video := range videos {
		log.Printf("convert %d of %d", i+1, len(videos))
		filename := filepath.Base(video)
		if err = c.convert(video, filepath.Join(c.Output, filename)); err != nil {
			return err
		}
	}
	return nil
}

func (c Converter) convert(in, out string) error {
	log.Printf("converting: %s to %s", in, out)
	// build the ffmpeg command
	cmd := c.Config.Command(in, out)
	return command.NewCommand(cmd, c.DryRun).ExecuteStream(os.Stdout)
}
