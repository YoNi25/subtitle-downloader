module main

go 1.15

replace utils => ./utils

replace input => ./input

replace downloader => ./downloader

require (
	downloader v0.0.0-00010101000000-000000000000
	github.com/matcornic/addic7ed v0.2.0
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/tools v0.1.0 // indirect
	input v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)
