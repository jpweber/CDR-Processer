#!/bin/bash

#link file in bin to Current dir
ln -s bin/cdrprocessor ./

#prompt user for app server name (app01 etc) and gatway type (gsx|nbs)
echo -n "What is the hostname or IP of the database server you wish to save records to?: "
read DBHOST

echo -n "What is the username to login with to create tables?: "
read DBUSER

echo -n "What is The Password to use for $DBUSER?: "
read DBPASS

echo -n "What is The Name of the database to create the tables in?: "
read DBNAME

echo "ok I will be creating Starts, Stops and Attempts tables on $DBHOST as $DBUSER in the Database $DBNAME"

`mysql -u $DBUSER -p$DBPASS -h $DBHOST $DBNAME < SQL/stops_table.sql`
`mysql -u $DBUSER -p$DBPASS -h $DBHOST $DBNAME < SQL/starts_table.sql`
`mysql -u $DBUSER -p$DBPASS -h $DBHOST $DBNAME < SQL/attempts_table.sql`

echo "Install process is complete. Now edit config.json.default and save as config.json"
