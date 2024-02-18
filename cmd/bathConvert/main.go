package main

import (
	"batchConvert/pkg/converter"
	"batchConvert/pkg/manifest"
	"flag"
	"fmt"
)

func main() {
	// prompt welcome message
	fmt.Printf("Batch Convert v%s\n", manifest.VersionName)
	fmt.Println("Note: default options was set to hevc_amf for high quality encoding which meets my need, " +
		"specify your own args for your need. (use -vc -ac -e, see help by -h).")
	// base config
	input := flag.String("i", "./", "input path")
	output := flag.String("o", "./", "output path")
	// args config
	VideoCodec := flag.String("vc", string(converter.VideoCodecHevcAmf), "video encoding codec")
	AudioCodec := flag.String("ac", string(converter.AudioCodecCopy), "audio encoding codec")
	extraFlag := flag.String("e", "", "extra flags")
	dryRun := flag.Bool("d", false, "dry run")
	flag.Parse()
	// applies
	opts := converter.DefaultOptions
	opts = append(opts, converter.WithRawFlags(*extraFlag))
	cfg := converter.NewConfig(converter.VideoCodec(*VideoCodec), converter.AudioCodec(*AudioCodec), opts...)
	if err := converter.NewConverter(*input, *output, *cfg, *dryRun).Convert(); err != nil {
		panic(err)
	}
}
