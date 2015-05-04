#!/bin/bash
OS=`uname`
#if [ $OS == "Linux" ]
#    then
#fi

#copy binary to system location
mkdir -p /usr/local/cargo
cp cargo /usr/local/cargo

#copy config to /etc
mkdir /etc/cargo
cp cargo.conf.default /etc/cargo/cargo.conf

#prompt user to install database tables
echo -n "Do you wish to have me install the database tables? [y/n]:"
read DBTABLES

if [ $DBTABLES == y ]
    then 

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
fi

echo "Install process is complete. Now edit /etc/cargo/cargo.conf to contain your confiruration parameters"
