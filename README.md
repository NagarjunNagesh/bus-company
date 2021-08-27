## Start the Server

`go run cmd/bus-company-server/main.go --scheme http --port=8080;`

## Microservice Explanation

> Prerequisite to use swagger commands below
https://goswagger.io/install.html

The code is written on top of auto generated code `go-swagger`. The `go-swagger` generates the cmd, models and restapi folders. However the **restapi/configure_bus_company.go** is the only go file which is **safe to be modified** without the fear of being overwritten. 

Inorder to modify the model or API endpoints, we should just modify the **swagger.yaml** file and run the following commands.

`swagger validate ./swagger.yaml`

The above command would tell us if the swagger yaml file is valid. If the swagger yaml file is valid then please proceed with the generation of the code. 

`swagger generate server -A bus-company -f ./swagger.yaml`

*Note that the **restapi/configure_bus_company.go** file would not be modified by executing the above command. So it is safe to generate the swagger command as many times as necessary*

# Swagger generated files
1. cmd*
1. restapi*
1. models*

# Golang Test

We are PackAndGo, a small bus company. We want to create a REST API that helps us manage the trips that we offer.

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