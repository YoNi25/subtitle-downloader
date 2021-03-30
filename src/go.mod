module main

go 1.15

replace utils => ./utils

replace input => ./input

replace downloader => ./downloader

require (
	downloader v0.0.0-00010101000000-000000000000 // indirect
	github.com/matcornic/addic7ed v0.2.0
	input v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)
