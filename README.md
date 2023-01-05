# Date Directory Creator
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
