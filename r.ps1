[string]$sourceDirectory = "assets"
[string]$destinationDirectory = "bin"
Copy-item -Force -Recurse $sourceDirectory -Destination $destinationDirectory
go build -o bin/simple.exe simple-sdl2-project/cmd/main ; start .\bin\simple.exe 
