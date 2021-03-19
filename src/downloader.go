package main

import (
	"fmt"
	addic7ed "github.com/matcornic/addic7ed"
	"log"
	"os"
)

type SearchSubtitle struct {
	name     string
	language string
}

type SubtitleToDownload struct {
	subtitle  addic7ed.Subtitle
	name      string
	extension string
	dirPath   string
}

func RetrieveShow(searchSubtitle SearchSubtitle) addic7ed.Subtitle {
	c := addic7ed.New()
	showName, subtitle, err := c.SearchBest(searchSubtitle.name, searchSubtitle.language)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ℹ️      TV Show %s found !\n", searchSubtitle.name)
	log.Println("------------------")
	log.Println(showName)          // Output: Shameless (US) - 08x11 - A Gallagher Pedicure
	log.Println(subtitle)          // Output: the best suitable subtitle in English language
	log.Println(subtitle.Version)  // Output: BATV
	log.Println(subtitle.Language) // Output: English
	log.Println("------------------")

	return subtitle
}

func DownloadShowsSubtitles(subtitleToDownload SubtitleToDownload) {
	log.Printf("⬇     Download srt for %s\n", subtitleToDownload.name)
	if _, err := os.Stat(subtitleToDownload.dirPath)
	os.IsNotExist(err) {
		log.Printf("🚧     Missing directory %s. Creating ...\n", subtitleToDownload.dirPath)
		os.MkdirAll(subtitleToDownload.dirPath, 0755)
	}
	subtitle := subtitleToDownload.subtitle
	err := subtitle.DownloadTo(fmt.Sprintf("%s/%s.%s", subtitleToDownload.dirPath, subtitleToDownload.name, subtitleToDownload.extension))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("🎉 Subtitle %s/%s.%s downloaded\n", subtitleToDownload.dirPath, subtitleToDownload.name, subtitleToDownload.extension)
}
