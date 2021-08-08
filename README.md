# goGrep

goGrep is a golang implemention of linux grep command.

### Usage

```
goGrep "pattern" file.txt
This matches and print all the occurence of term "pattern" and print it on the screen.

goGrep -r "pattern" SearchFolder
This matches and print all the occurence of term "pattern" present in the files in the folder SearchFoler

goGrep -i "PATTERN" file.txt
This matches and print the occurence irrespective of the case of the match

goGrep -o output.txt "pattern" file.txt
This usage outputs the results into the file output.txt instead of stdout

goGrep -c "pattern" file.txt
This prints the number of matches found
```
