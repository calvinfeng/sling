# About Sling
Sling messaging mimics Slack

# Usage
Set up your local SQL database, you will need to install Postgres if you don't already have it. This link does a great job of doing this.

http://postgresguide.com/setup/install.html
 
Add jcho as a superuser.

Upen the psql envirnment to make changes.
```shell
$ sudo -u postgres psql
psql (10.9 (Ubuntu 10.9-0ubuntu0.18.04.1))
Type "help" for help.

postgres=#
 ```
Run the following commands to add the superuser jcho.
```shell
postgres=# create role jcho superuser login;
postgres=# alter user jscho with password 'jcho'
postgres=# create database sling with owner=jcho; 
```

Compile and run the server code:
```shell
go install
sling runmigrations
```

Build the frontend:
```unix
cd frontend/
npm run build
```

Finally, run the server:
```unix
sling runserver
```

Open localhost:8888 on your browser to begin chatting!

