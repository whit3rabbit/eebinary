## Encrypt/Embed Binary

* Read file into bytes
* Encrypt with AES-256
* Compress with ZLIB
* Generate a seperate Go file with encrypted/embeded binary
* Create new executable based on new Go script

This is a POC test.  

It is currently only working on Unix because memexec doesn't work.  
If you change it on Windows to output the file on Windows it will flag in some Antivirus software.

## Instructions
```yaml
# Build sample executable
cd sample
go build

# Build eebinary
cd ..
go build
eebinary -i sample/sample.exe -o test.go
eebinary -i sample/sample.exe -o test.go -w # For Windows binary

# Build new encrypted executable
set GOARCH=386 # For Windows binary
go build test.go
./test
```