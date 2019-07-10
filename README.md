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
go install sling
sling runmigrations
sling runserver
```
Open localhost:8888 on your browser to begin chatting!

# How to Test the Frontend UI Using NPM

The following commands will allow you to host the frontend files on `localhost:3000` using npm. They will NOT connect properly to the backend using these commands, only allow you to test the UI.

```unix
cd frontend/
npm install
npm start
```
Open `localhost:3000` to see the content.

# How to Compile the Frontend (make a new index.bundle.js using Webpack)

The following commands will allow you to host the frontend files on `localhost:8888` using the echo server. This will allow you to proprly connect the frontends' requests with server responses.

Navigate to the frontend folder

```unix
cd frontend/
```
Run all npm installs to check that all tools are availible in your environment [ NOTE: we could be missing some here]

```unix
npm install --save-dev webpack webpack-cli
npm install --save react react-dom
npm install --save-dev @types/react @types/react-dom
npm install --save-dev typescript ts-loader source-map-loader
npm install babel-preset-react
npm install --save-dev css-loade
```
To build the index.bundle.js in the public folder:
 ```unix
 npx webpack
 ```
 You should recieve an output like the following upon success:
```bash
kbowman@FR-1071:~/Workspace/src/sling/sling/frontend$ npx webpack
Hash: 24850e466b3a13a663d1
Version: webpack 4.35.3
Time: 5727ms
Built at: 07/10/2019 10:19:52 AM
          Asset      Size  Chunks             Chunk Names
index.bundle.js  6.26 MiB    main  [emitted]  main
Entrypoint main = index.bundle.js
[./node_modules/css-loader/dist/cjs.js!./src/App.css] 663 bytes {main} [built]
[./node_modules/css-loader/dist/cjs.js!./src/index.css] 517 bytes {main} [built]
[./node_modules/webpack/buildin/global.js] (webpack)/buildin/global.js 472 bytes {main} [built]
[./src/App.css] 1.05 KiB {main} [built]
[./src/App.tsx] 534 bytes {main} [built]
[./src/index.css] 1.06 KiB {main} [built]
[./src/index.tsx] 1.13 KiB {main} [built]
[./src/serviceWorker.ts] 5.37 KiB {main} [built]
    + 502 hidden modules
 ```
