# WGzip
Nothing too special. Just a simple package that will compress/uncompress a single file in gzip standards.

## Usage
`wgzip.Gzip(<SOURCE>, <DESTINATION>)`

Simply takes the file from `SOURCE` and compresses it to the `DESTINATION`. It will add `.gz` to the end of the file.

`wgzip.UnGzip(<SOURCE>, <DESTINATION>)`

Simply takes the file from `SOURCE` and decompresses it to the `DESTINATION`. It will remove `.gz` from the end of the file.

**IMPOTANT**: In both cases `DESTINATION` is a directory not a file.

## TODO
* Add support for multiple files
