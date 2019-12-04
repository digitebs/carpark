# Carpark
API that gives the nearest carpark for a given points in the map. 

## Getting Started
The project is written in _golang_ and can be started using `go build && go run`,
I would recommend using _goland_ from _intellij_ when importing this project manually. 

For simplicity, this project comes with _golang_ base image for easy build and run.

### Postgres Schema
* hdb_carpark_information
* hdb_carpark_availability

### Prerequisites
* docker

### Installing
1. `docker-compose up` will start 2 container one db and one web.
2. Open browser and run the paste the below
    ```
    http://localhost:10000/carparks/nearest?latitude=1.37326&longitude=103.897&page=1&per_page=10
    ```
    You'll get an empty response... :(

3. Ok! lets insert **hdb_carpark_availability** data
    ```
    docker run --env-file web.env carpark_web pg-update 
    ``` 
    Try refreshing the browser. 

## Running the tests
Requires _postgres_ database which will be available once the
 database container is up

```
docker run --env-file web.env carpark_web go test -v ./...
```

#### Updating carpark information (optional)
1. Download the **hdb_carpark_information** on the gov site, place it in the _data_ folder
2. `docker run --env-file web.env carpark_web csv-import` will convert the geo format of the csv

## Build with
* golang
* mux
* docker-compose
* postgis

## References

https://tour.golang.org/

https://medium.com/rungo/error-handling-in-go-f0125de052f0

https://thenewstack.io/understanding-golang-packages/

https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc

https://golang.org/pkg/testing/

https://docs.docker.com/compose/gettingstarted/

https://medium.com/@Umesh_Kafle/postgresql-and-postgis-installation-in-mac-os-87fa98a6814d

https://github.com/go-pg/pg/wiki/Model-Definition

https://github.com/go-pg/pg/wiki/Writing-Queries

https://www.alexedwards.net/blog/organising-database-access

https://medium.com/@felipedutratine/pass-environment-variables-from-docker-to-my-golang-2a967c5905fe
