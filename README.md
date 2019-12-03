###### Stack
* golang
* mux
* docker-compose
* postgis

###### Prerequisite
* docker

###### Run
1. `docker-compose up` will start 2 container one db and one web
2. open browser and run the paste the below
    ```
    http://localhost:10000/carparks/nearest?latitude=1.37326&longitude=103.897&page=1&per_page=10
    ```
    you'll get an empty response...

3. Ok lets insert carpark availability data: `docker run --env-file web.env carpark_web pg-update `

###### Updating (optional)
1. Download the carpark information on the web, place it in the data folder
2. `docker run --env-file web.env carpark_web csv-import` will convert the geo format of the csv

###### Reference

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
