# subtitle-downloader
Download subtittle from Addic7ed

## Build
To build sources to an executable, go to `src` directory, then run 

```
go build -o ../bin/subtitle-downloader .
```

## Usage
Go to `bin` directory, then run

```
./subtitle-downloader --showStr 'FileNameOfTheTvShow' --dirPath 'WhereToDownloadSubtitle' --language 'SubtitleLanguage'
```

### Explanation
| Parameter 	| Type   	| Optional                      	| Description                                                                                                                         	| Example                                                               	|
|-----------	|--------	|-------------------------------	|-------------------------------------------------------------------------------------------------------------------------------------	|-----------------------------------------------------------------------	|
| showStr   	| string 	| false                         	| The TV show episode that you wanted to download subtitles. <br /> Commonly, the value is just the file name of your TV Show episode 	| The.Falcon.and.the.Winter.Soldier.S01E01.WEBRip.x264-PHOENiX[eztv.re] 	|
| dirPath   	| string 	| true <br />Default : `.`      	| The Directory where the .srt file is going to be download. If the directory does not exist, it will be created                      	| ~/Series/The Falcon and the Winter Soldier/S01                        	|
| language  	| string 	| true <br />Default : `French` 	| The language for which you are looking for subtitles.                                                                               	| French                                                                	|