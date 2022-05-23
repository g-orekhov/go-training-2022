# Test Routes

## Ping

`http://localhost:8081/ping`

## Events List

GET `http://localhost:8081/v1/events`

## Event (by ID)

GET `http://localhost:8081/v1/events/{id}`

## Event (by Geolocation)

GET `http://localhost:8081/v1/events/geo/{lat},{long},{dist}`

 lat, long - decimals
 dist - distance in meters

## Create Event

JSON structure

    "name": " ",
    "short_descr": " ",
    "description": " ",
    "images": [" ", ...],
    "preview_img": " "

POST `http://localhost:8081/v1/events`

## Update Event

PUT `http://localhost:8081/v1/events/{id}`

## Delete Event

DELETE `http://localhost:8081/v1/events/{id}`

## Users List

`http://localhost:8081/v1/users`

## User (by ID)

`http://localhost:8081/v1/users/{id}`