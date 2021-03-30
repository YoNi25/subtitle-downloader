package main

func main() {
	configuration = InitializeConfig()
	colors = InitializeColors()

	input := ReadInputArgs(configuration, colors)

	subtitle := RetrieveShow(SearchSubtitle{
		name:     input.showName.fullname,
		language: input.language,
	})

	DownloadShowsSubtitles(SubtitleToDownload{
		subtitle:  subtitle,
		name:      input.showName.fullname,
		extension: configuration.SubtitleExtension,
		dirPath:   input.dirPath.fullPath,
	})
}
