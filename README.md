# Welcome to Mnemonics

It generate a private key with Mnemonics

## Overview

I have used Revel framework to build the APIs and the Test Cases for them. Revel is not targeted for APIs and has broader functionalities available but I have trimmed down some of those from request processing chain like- session filter, flash store filter. (Changes can be reviewed in-app/init.go)

Revel does not automatically make JSON request attributes available as action arguments, it does this automatically for x-www-form-urlencoded and form-data. To overcome this I have implemented a custom filter in request pipeline which achieves this functionality. Please look into app/init.go line number #69 for more information about this implementation. The implemented api accepts JSON requests as well as x-www-form-urlencoded and form-data requests.

Revel provides logging functionality "out of the box" so nothing has to be done explicitly by us.

Validations are implemented using revel validations.

Id attribute in params is not used anywhere but this api requires it.

Test cases in /tests dir documents the API.

Because we are using revel, the code can not directly be build using "go build", I know this was mentioned in the test description, but thought of going one mile ahead and used Revel. If you wish I can also extract out the API out of revel and build it again without any framework. But I recommend using a framework to avoid re-inventing the wheel for the required boilerplate code. 

## Enhancements required
1. Add more constraints in the password like:- must have a special character etc.
2. Extract out the business logic in services.

### Setup and dependencies installation
1. Install revel
```
go get github.com/revel/cmd/revel
```

2. Make revel command available from everywhere
```
export PATH="$PATH:$GOPATH/bin"
```

3. Install dep dependency manager for go
```
go get -u github.com/golang/dep/cmd/dep
```

4. Need to move the code of this repo to GOPATH src like (Note: will be different for you)
```
gocode/src/github.com/xxxx/mnemonic
```

5. Install dependencies with dep by running below command in project directory
```
dep ensure
```

By running "dep ensure" these will be installed "revel", "go-bip32", "go-bip39".

### Run in development mode
```
revel run $GOPATH/src/<path_to_project>/mnemonic
```

Go to http://localhost:9000/ and you'll see it running

### Run in production mode
```
revel run $GOPATH/src/<path_to_project>/mnemonic prod
```

Go to http://localhost:9000/ and you'll see it running

### API sample

```
POST /private_key

Content-Type application/json

BODY JSON
```

Sample request body
```
{
    "id": "123",
    "password": "sdfsdfsdf"
}
```

Sample response
```
{
    "key": "xprv9wshmRFzFW3MURnzWJPU8B67EKJmqjkF4TGJ9JGKqeKMCoEMmWUNngRWujk5Dq8Rsbt3JDw2nXQN2zj5Sk7ycqmPXgJgmKW5mCLhe2dRddU",
    "mnemonic": "brisk tortoise culture tumble pistol weekend section honey art throw topple goose item script doctor social swallow trigger garment govern kid host ecology hollow"
}
```

### Run tests
```
revel test $GOPATH/src/<path_to_project>/mnemonic dev
```

You can also run tests from browser by running this application in development mode and navigation to http://localhost:9000/@tests

### Build
```
revel package $GOPATH/src/<path_to_project>/mnemonic
```
For more information please visit [Deployment](https://revel.github.io/manual/deployment.html).

### Logging
Revel automatically generates logs for requests and they are written to mnemonics-request.log, logs for warning and errors are written to /log dir of this project.

I have not logged the request params and responses as the both the things should be secure and should not be preserved on the server and to comply with various security standards.

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory
    tests/            Test suites
