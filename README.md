
# Sonus CDR Processor
*Very much in development and incomplete. Do not use yet.*

## Description
Application to read Sonus CDR files parse the data, map the fields to keys and insert in to a database. 
Currently only saving to MySQL other databases on the roadmap.

Can be run as daemon or as command line tool. In daemon mode will process files from directory or specific location. As CLI tool can take filename as argument and will only process that file

## Installation
#### Application
Run the install.sh script and answer the questions. This will create an exectuable for your in the bin dir and will create the tables for you in your MySQL database. You can now copy your binary and your config file to where ever you which to run them from. 

Currently the binary and config file can be anywhere as long as they are in the same directory together. The application looks for the config.json file in current working directory.   

####SQL Tables
in the SQL dir are files to create the three required MySQL tables. You can create the manually or use the install.sh script to create them for you.

Replace username, password and databasename with your values. These should also match what you put in the config file later. 

## Usage  
Can be ran as a command line tool to process a single file, or ran in a daemon mode from a startup script

### options
-h = display helps  
-v  = displays version  
-f = path to file name to parse. Will only parse that file and exit. 
-t = Number of records to commit in a single transaction to the database

### Example  
Process a single file  
`./cdrprocessor -f SFOSBC01.20150331235500.10066FE.ACT`
This will cause the program to not run in daemon mode. It will exit once this file is complete.

`./cdrprocessor -t 500`  
This will start the program in daemon mode with a transaction commit of 500 records at a time.

## Documentation

### Configuration File
Currently the configuration file is a simple json file. This may change in the future. The easiest way is to use this is to copy the shipped config.json.default to config.json and modify the entries to suit your environment. 

#### Fields
__FileDir__  = Directory where the Raw CDR files are stored.  
__DbPort__ = The port of the database server to connect to. for MySQL this is usually 3306  
__DSN__ = The datasource name. This information is used when connecting to the database. This Will include the username and password to use when connecting to the database. The database host port and database name to connect to. See guide: `"<username>:<password>@<host>:<port>/<database>"`
__FileExt__ = The file extension of the files you wish to process. The application will look in the specified Directory for files ending in this file extenstion. So if you want to process *.ACT files use ACT as the value here. 


Exmaple  

    "FileDir" : "/nsfstore/CDR_Files",  
    "DSN"   : "root:root@tcp(127.0.0.1:3306)/test",
    "FileExt" : "ACT"

### Logging examples
still to come.


