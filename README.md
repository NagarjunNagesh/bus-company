## Clean Architecture

This is an implementation of Clean Architecture for the pack-and-go application. 

* Independent of Web layer. The business domain is unaware of the existence of the web layer. As such we could swap the web API layer with a messaging broker. 
* Independent of Database layer. The repository layer is dependent on the domain layer's repository interface. As such we could swap the underlying repository with anything. 


The project has 4 layers

* Entity layer
* Domain layer
* Repository layer
* Delivery layer


### The Architectural Diagram

![alt text](https://github.com/NagarjunNagesh/bus-company/blob/main/resources/readme/High%20Level%20Architecture.png)

This architecure lets us to design our application to be independent of the low level components (Repository, Delivery layer) from our high level components (Entity, Domain layer).

Keeping the high level component clean and independent of the low level components lets us swap these components independent of each other. 

## Start the Server

**Building the application**

`go build -o main cmd/github-com-nagarjun-nagesh-bus-company-server/main.go`

**Running the application**

`go run cmd/github-com-nagarjun-nagesh-bus-company-server/main.go --scheme http --port=8080;`

### Base URL

The base url for the packandgo server is as follows:

`localhost:8080/v1/`

## Low Level Architecture

### Cities Architecture

The file **cities.txt** is read every 5 minutes. The time can be configured in the **config/config.go** file using the **ReloadTime** constant. 

The `go-swagger` auto-generated codes let us modify one of its files safely without the fear of it being modified. Refer **Autogenerated Codes** section for more information. 

We have used that file to call the repository to populate the cities repository object which is used to populate the trips. 

![alt text](https://github.com/NagarjunNagesh/bus-company/blob/main/resources/readme/cities-low-level-architecture.png)

Clean architectural design has been used inorder to revert the dependencies (DIP) between the domain layer and the repository layer. 

The repository interface is called from the configuration layer when the system is compiling the code to populate the cities object. And the configuration layer which is autogenerated by the `go-swagger` also has the function to populate the cities object every X seconds. 

#### Modifying the cities file

After modifying the cities file, We can do one of the following:

* Wait for 5 minutes for the cities object to be repopulated. 
* Restart the server.

### Trips

* The trip web layer (Handler class) is dependant on the domain layers __entity__ and __use case interface__ layers inorder to convert the relavent responses from the domain layer to web layer models. 
* The trip repository implementation layer is dependant on the domain layers __repository interface__ layer inorder to derive its implementation details. 
* The domain layer is responsible in this case for the successful retrival of the data and passing it on to the web layer.

The important information is that the web layer is oblivious of the repository layer and the repository layer is oblivious of the web layer. All of these layers are connected through a high level domain layer. 

It makes the layers independant and maintainable. We can swap the repository layer with a database or the web layer with a messaging broker or to a command line interface without touching the domain layer or the repository. 

### Trip Architectural Diagram

![alt text](https://github.com/NagarjunNagesh/bus-company/blob/main/resources/readme/trips-low-level-architecture.png)

### Error Scenarios

Detailed below are the error scenarios which could occur.

#### Add Trips Errors

The cities and the dates would have to be valid inorder to be added. 

1. The origin city would have to be valid, else it would throw an error `origin city invalid`
1. The destination city would have to be valid, else it would throw an error `destination city invalid`
1. The dates would have only contains these values seperated by comma `'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'` else it would throw an error `invalid date value or repeated dates 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'`
1. The dates should not contain repeated values else it throws the same error as above.
1. The Origin ID and the Destination ID being the same would thrown an error `destination id cannot be the same as origin id`
1. Price should be greater than 0 else it would throw an error `the price of a trip cannot be zero or less`

*Inorder for the city to be valid, The cities text file would have to be non empty and the number is the line number present in the city text file.*

#### Fetch Trips or Fetch all trips error

The fetch trips and fetch all trips error scenarios occurs due in the following scenarios:

1. The Origin ID cannot be parsed in to a city name which would throw and error `"cannot find origin city for the trip id %s`. The format *%s* is the trip ID. 
1. The Destination ID cannot be parsed in to a city name which would throw and error `"cannot find destination city for the trip id %s`. The format *%s* is the trip ID. 
1. **Only applicable for fetch trip** When the trip ID requested is not found then the following error is thrown `no content`.

## Autogenerated Codes

> Prerequisite to use swagger commands below
https://goswagger.io/install.html

The code is written on top of auto generated code `go-swagger`. The `go-swagger` generates the cmd, models and restapi folders. However the **configure_github_com_nagarjun_nagesh_bus_company** is the only go file which is **safe to be modified** without the fear of being overwritten. 

Inorder to modify the model or API endpoints, we should just modify the **swagger.yaml** file and run the following commands.

`swagger validate ./swagger.yaml`

The above command would tell us if the swagger yaml file is valid. If the swagger yaml file is valid then please proceed with the generation of the code. 

`swagger generate server -A github.com/NagarjunNagesh/bus-company -f ./swagger.yaml`

*Note that the **restapi/configure_bus_company.go** file would not be modified by executing the above command. So it is safe to generate the swagger command as many times as necessary*

# Swagger generated files
1. cmd*
1. restapi*
1. models*

# Golang Test

We are PackAndGo, a small bus company. We want to create a REST API that helps us manage the trips that we offer.

`go test ./...`

## Goal

Your goal is to **build a REST API** with the following characteristics:

| Method | Endpoint  | Description          |
|--------|-----------|----------------------|
| GET    | /trip     | List all trips       |
| POST   | /trip     | Add a new trip       |
| GET    | /trip/:id | Get trip with ID :id |

### REST API Description

We would like the trips to be obtained in the following format:

```json
{
    origin: "Barcelona",
    destination: "Seville",
    dates: "Mon Tue Wed Fri",
    price: 40.55
}
```

Whereas to add a trip, we would send the following:

```json
{
    originId: 2,
    destinationId: 1,
    dates: "Sat Sun",
    price: 40.55
}
```

The **trip ID** should be added automatically. Each trip should have a unique trip ID.

The **list of cities** is in a text file, *cities.txt*, but perhaps we will change that in the future as our company grows. Right now, every line in the text file is a city. The first line corresponds to cityId=1, the second to cityId=2, etc.

### General guidelines

We want you to build a REST API that works with our current needs, but that can be ready to **change easily in the future**, without having to rewrite the whole system or fearing that something will break. So please, try to make it as future-proof as possible.

Use the packages that you think are suitable for the job, as well as the **architecture and code structure** that makes most sense from your point of view. Feel free to move, split, etc. the provided files into the files and folders of your choice.

You can also challenge and *change the proposed API structure if you feel it is necessary*, as long as you give a reason why you have decided to do things in a different way.