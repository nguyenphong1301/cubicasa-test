
### Start Project
before start service please run ```go mod download && go mod vendor``` for install required library 
#### via docker
```docker-compose --env-file .env up```
#### run on your local host
you need to:
- edit ```.env``` file to mapping the right database and port.
- create database with name similar your config, default is ```cubicasa_test```
- run ```./data_migrate/public.sql``` file to create database schema. 
- run ```go run main.go``` to start service

### Test
- access ```http://127.0.0.1:{$port}/swagger/index.html``` to see document and test api.

#### default port is 3000 to avoid conflict port in your machine, change to suitable port if you need. 
