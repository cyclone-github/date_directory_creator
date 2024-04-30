[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=cyclone-github&repo=date_directory_creator&theme=gruvbox)](https://github.com/cyclone-github/)
# Cyclone's Date Directory Creator
### CLI program to create directories by year / month / day

- version:        0.3.1
- build date:     2023-01-05-1600
- written by:     cyclone

### Usage Examples:
- (create directories for all days of 2025)

date_directory_creator.bin -year 2025

- (create directories for all days of 2020-2030)

date_directory_creator.bin -year 2020-2030

- (create directories for all sunday's of 2025)

date_directory_creator.bin -year 2025 -day sunday

- (create directories for monday-friday of 2020-2030)

date_directory_creator.bin -year 2020-2030 -day monday-friday

- Program will check for and not overwrite existing directories, but use with caution.

### Compile from source:
- If you want the latest features, compiling from source is the best option since the release version may run several revisions behind the source code.
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/date_directory_creator.git`
  - `cd date_directory_creator`
  - `go mod init date_directory_creator`
  - `go mod tidy`
  - `go build -ldflags="-s -w" .`
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt
