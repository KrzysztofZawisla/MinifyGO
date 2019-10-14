# MinifyGO
Tool for minifying your JavaScript code.

## Important
The program requires installed globally [minify package](https://www.npmjs.com/package/minify) from npm.

## Intruction
1. Make next to the program "minify.json" file.  
2. Create a list with files to minify.  
For example:
```json
{
  "files": [
    "dist/main.js",
    "dist/script.js"
  ]
}
```
3. Run the program.

## Flags
-outputToTheSameFile - When you set this flag to the true program will save minified code to the same file. When you set this flag to the false program will create a new file with extension .min.js next to the origin file.

## Additional info  
### Working environment  
Windows 10 Professional 64-bit version 1903 build 18362.418  
Version of Golang: 1.12.9 windows/amd64
### Author
Krzysztof Zawisła
