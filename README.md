# subtitle-downloader
Download subtittle from Addic7ed thanks to https://github.com/matcornic/addic7ed

## Build
To build sources to an executable, go to `src` directory, then run 

```
go build -o ../bin/subtitle-downloader .
```

## Usage
Go to `bin` directory, then run

```
./subtitle-downloader
```

### Usage Example
![demo](doc/screenshots/demo.gif)

### Warning cases

#### Unexpected Directory Input
![Unexpected Directory Input](doc/screenshots/warnings/unexpected-directory-input.png)

The input received for the Directory Path does not exist on the list.
In this case, the default directory path will be set

#### Unexpected Language Input
![Unexpected Language Input](doc/screenshots/warnings/unexpected-language-input.png)

The input received for the Language does not exist on the list.
In this case, the default language path will be set

#### Mismatch subtitle versions
![Mismatch subtitle versions](doc/screenshots/warnings/mismatched-subtitle-versions.png)

Each subtitle is linked to a version. This version is contained on the Episode name.
When the version contained in the Episode name seems to be different from the found one, 
this message is displayed to let you check the versions' compatibility

#### Missing Targeted Directory
![Missing Targeted Directory](doc/screenshots/warnings/missing-targeted-directory.png)

If the directory does not exist, it will be created automatically before downloading subtitles.

### Failure cases
#### Episode name is invalid
![Episode Name is invalid](doc/screenshots/errors/unable-to-parse-episode-name.png)

We are waiting a specific format for the Episode name. In fact the Episode name should look like this
```
<TVShowName>.S<SEASON>E<EPISODE>-<VERSION>
```

<b>Example</b>: `The.Falcon.and.The.Winter.Soldier.S01E03.Power.Broker-TOMMY`

#### Episode Not found
![Episode Not Found](doc/screenshots/errors/episode-not-found.png)

The Episode was not found on the addic7ed servers.

#### Subtitle Not Found For Language
![Subtitle Not Found For Language.png](doc/screenshots/errors/subtitle-not-found-for-language.png)

The Episode has no subtitles for the chosen language
