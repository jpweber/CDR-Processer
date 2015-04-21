
# CDR Server web service

## Description
The CDR server is a web service written in go to serve raw CDR files via a basic web API.
This service is designed to run on the CDR Loader server so that it can serve the files that the loader is collecting from the DSIs and the NBSes.

## Instalation
Right now the is only a Go binary and an init script. These are copied in to their respective directories. Which currently is /etc/init.d and /home/cdr/cdr_server/
> this will need to change. Expecting to make a .deb package and be able to install using apt-get from our internal repo if I could every figure where I was on that. 

## Usage  
just start the daemon  
`sudo service cdr_server start`

### Example  
`sudo service cdr_server start`  
`sudo service cdr_server stop`  
`sudo service cdrs_erver restart`  

## Documentation
cdr_server is ran as user cdr. Otherwise there are permission problems with zipping the files.  
a pid file is written to `/tmp` This pid file is created and cleaned up via the init script. This pid file is also how the service is monitored via watchcat. 

### options
None

### API
API Supports GET and PUT HTTP methods currently.  
Daemon runs on port 8000  
URL = hostname:8000/resource/`<command>` 

Currently you can 
- fetch a file from the server. 
- Get an MD5 hash checksum of a specific file 
-  zip a specified file. Currently bzip2 is the zip that is implemented for legacy reasons.

#### Get a cdr file
- Make a `GET` request with  just the path/filename. See example below
		`http://loader01:8000/CMHNBS1/172.31.10.72/101F85D.ACT`
- This will return a raw text stream of the CSV file. No compression or fancy tricks.
-  If the file is not found a `404 Not Found` will be returned. 
	
#### Get an MD5 Hash of a file
- Make a `GET` request with the path/filename and `hash` as the command. See example below
		`http://loader01:8000/CMHNBS1/172.31.10.72/101F85D.ACT/hash`
- This will return the md5 hash of the requested resource.
- If the file is not found a `404 Not Found` will be returned. 

#### Zip a file.
- Make a `PUT` request with the path/filename and `zip` as the command. See example below
		`http://loader01:8000/CMHNBS1/172.31.10.72/101F85D.ACT/zip`
- This will return a `200 ok` if the zip succeeds. Otherwise it will return a `304 Not Modified` if the file is not able to be zipped. 

### Logging examples
```
Mar 13 13:18:10 loader01 CDR_SERVER[59494]: GET - /CMHNBS1/172.31.10.72/1020D1D.ACT
Mar 13 13:18:10 loader01 CDR_SERVER[59494]: Returning the file /CMHNBS1/172.31.10.72
Mar 13 13:18:32 loader01 CDR_SERVER[59494]: PUT - /CMHNBS1/172.31.10.72/1020D1D.ACT/zip
Mar 13 13:18:32 loader01 CDR_SERVER[59494]: zipping /home/cdr/ready_cdrs/CMHNBS1/172.31.10.72/1020D1D.ACT
Mar 13 13:18:35 loader01 CDR_SERVER[59494]: Completed zipping/home/cdr/ready_cdrs/CMHNBS1/172.31.10.72/1020D1D.ACT
```


