\c gis;

CREATE TABLE hdb_carpark_information
(
    car_park_no            VARCHAR(50) PRIMARY KEY,
    address                VARCHAR(250),
    x_coord                float,
    y_coord                float,
    car_park_type          VARCHAR(50),
    type_of_parking_system VARCHAR(50),
    short_term_parking     VARCHAR(50),
    free_parking           VARCHAR(50),
    night_parking          VARCHAR(50),
    car_park_decks         integer,
    gantry_height          float,
    car_park_basement      VARCHAR(50)
);

CREATE TABLE hdb_carpark_availability
(
    car_park_no    varchar(50),
    lot_type       varchar(50),
    total_lots     integer,
    lots_available integer
);

COPY hdb_carpark_information
    FROM '/var/lib/postgresql/data/pgdata/sample.csv' DELIMITER ',' CSV HEADER;


SELECT AddGeometryColumn('hdb_carpark_information', 'geom', 4326, 'POINT', 2);
UPDATE hdb_carpark_information
SET geom = ST_SetSRID(ST_MakePoint(y_coord, x_coord), 4326);

SELECT *
FROM hdb_carpark_information i,
     hdb_carpark_availability a
WHERE i.car_park_no = a.car_park_no
  AND lots_available > 0
ORDER BY geom <-> ST_SetSRID(ST_MakePoint(9.18233279, 45.47184631), 4326)
LIMIT 100;



