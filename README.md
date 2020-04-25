# lucky-draw
lucky-draw is a program which will pick an item randomly from a pool of items and declare it as the winner.

## Installation
No installation is required. Just download/clone the repository.
## Usage
All you have to do is run the binary _lucky-draw_ in the repository, with the path to a text file (which contains the 
items eligible in the lucky draw) as an argument.

**Example:** 
```go
./lucky-draw path/to/file.txt
```
All the items in the text file must be in separate lines.     
  
**Format of txt file:**
```go
Alpha
Beta
Charlie
``` 
`example-files/example.txt` is an example txt file.
