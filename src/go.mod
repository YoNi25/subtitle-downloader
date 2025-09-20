module main

go 1.24.0

replace command => ./command

replace downloader => ./downloader

replace input => ./input

replace utils => ./utils

require command v0.0.0-00010101000000-000000000000

require (
	downloader v0.0.0-00010101000000-000000000000 // indirect
	github.com/PuerkitoBio/goquery v1.10.3 // indirect
	github.com/andybalholm/cascadia v1.3.3 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/masatana/go-textdistance v0.0.0-20191005053614-738b0edac985 // indirect
	github.com/matcornic/addic7ed v0.2.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/tkanos/gonfig v0.0.0-20210106201359-53e13348de2f // indirect
	golang.org/x/net v0.44.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	input v0.0.0-00010101000000-000000000000 // indirect
	utils v0.0.0-00010101000000-000000000000 // indirect
)
