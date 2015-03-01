# zniconv

Package zniconv provides a Reader to convert the charset of data.
It wraps an io.Reader, and converts the data read from the io.Reader to the target charset.
The actual conversion is done by the glibc iconv.

See main/zniconv.go for example.

