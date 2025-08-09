package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"rpp-parser/pkg/lexer"
	"rpp-parser/pkg/parser"
)

//NOTE: My reaper project represent a node which contains some options and other
// nodes. There are:
// 1. MEDIA ITEMS - which is potentially other reaper projects
// 2. MARKERS - which are potentially regions
// 3. TRACKS
//
// Every option stats off a new line, first we have a name and then list of params
// Every item inclosed with "<" and ">" and potentially contains it own options
// and can spread across multiple lanes

func main() {

	f, _ := os.Create("baseproject.log")

	defer f.Close()

	logger := slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	fileBytes, err := os.ReadFile("/home/butbkadrug/Music/REAPER/Your Love Never Fails - Bb - 114 BPM/Your Love Never Fails - Bb - 114 BPM.rpp")

	if err != nil {
		log.Fatal(err)
	}

	slog.Debug("file sucessfully read", "bytes", len(fileBytes))

	fileString := string(fileBytes)

	l := lexer.NewLexer(fileString)
	p := parser.NewParser(l)
	proj := &parser.Project{}
	p.ParseProject(proj)

	slog.Debug("parsing finished", "proj", proj)

	for i, track := range proj.Tracks {
		fmt.Println(i+1, track.Name)
		for _, item := range track.Items {
			fmt.Println("\t", item.Position, item.Name, item.SourceSection.Source.File)
		}
	}

}
