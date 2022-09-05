### Prerequisite
Create a MongoDB atlas account, is free and there you will get the connection string to fill the env file

<a href="https://www.mongodb.com/cloud/atlas/register" target="_blank">MongoDB Sign up</a>

Change .envexample to .env to be able to run the project, and fill the MongoDB connection string with a proper one


### Test
Run the below command to run the tests, and make sure that all tests are passing (I had no time to create the test, but I will)
```
make test
```

### Running
Run the commands below to run the app locally
```
go mod download
```
```
make run-server
```

Run the commands below to run app in docker
```
make create-docker-image
```
```
make run-docker
```

### Api Documentation
Check the link below to see the documentation and examples of usage!

<a href="https://documenter.getpostman.com/view/19131749/VV4wSGqE" target="_blank">Postman Doc</a>