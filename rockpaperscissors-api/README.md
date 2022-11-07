# Rock, Paper, Scissors API

## Software

This project has been developed with the following versions:

- go version go1.19.2 darwin/amd64
- GNU Make 3.81

## How to build the project

```sh
make
```

## How to run the project

```sh
make run
```

## How to test the project

Without coverage (just the tests):

```sh
make test
```

Open coverage window when finished

```sh
make coverage
```

*NOTE: The coverage will display onluy the coverage for the packages with at least one \*_test.go file inside*

## TODO:

- [ ] Explain project structure
- [ ] Add more testing to the project and increase coverage
- [ ] Add integration testing (API calls)
- [ ] Add cross compile options
- [ ] Add Dockerfile and docker-compose
- [ ] Display player statistics
- [ ] Use a real database
- [ ] Separete database logic and websocket logic
- [ ] Separate domain logic and websocket logic
- [ ] Refactor the code
- [ ] Remove unused comments
- [ ] Improve README
...
