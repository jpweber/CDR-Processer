
# Sonus CDR Processor
*Very much in development and in complete. Do not use yet.*

## Description
Application to read Sonus CDR files parse the data, map the fields to keys and insert in to a database. 
Currently only saving to MySQL other databases on the roadmap.

Can be run as daemon or as command line tool. In daemon mode will process files from directory or specific location. As CLI tool can take filename as argument and will only process that file

## Instalation
Thinking maybe just a simple make file or shell script. Ideally an ubuntu package as well

## Usage  
Can be ran as a command line tool to process a single file, or ran in a daemon mode from a startup script

### options
-h = display helps  
-v  = displays version  
-f = path to file name to parse. Will only parse that file and exit. 
-t = Number of records to commit in a single transaction to the database

### Example  
Process a single file  
`./main -f SFOSBC01.20150331235500.10066FE.ACT`
This will cause the program to not run in daemon mode. It will exit once this file is complete.

`./main -t 500`  
This will start the program in daemon mode with a transaction commit of 500 records at a time.

## Documentation

### Configuration File


### Logging examples



