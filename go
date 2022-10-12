#!/bin/bash
# Bash script to print the strings “none”,” two”,” three”,” four”, and “five” on the screen with each appearing on a separate line:
#       - Create a text file that has each of these on a separate line and the file name starts with today's date “yyyy-mm-dd-file.txt”
#       - Note; don’t add today’s date manually
# Date - $(date +%F)

printf "%s\n" "one" "two" "three" "four" "five" | tee $(date +%F)-file.txt

