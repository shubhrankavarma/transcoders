
# Transcoders

VOD3Transcoders is a microservice to manage transcoders. This service provides support for CRUD operations on transcoders. Following are the handlers of this service


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`TRANSCODERS_SRV_PORT`  // Port on which the service will run

`TRANSCODERS_HOST` // Host on which the service will run 

`DB_USER` // Database user

`DB_PASS` // Database password

`DB_NAME` // database name 

`DB_URL` // database url 

`TRANSCODERS_COLLECTION` // Collection name

`JWT_TOKEN_SECRET` // JWT token secret 


## API Reference
|  Headers | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required** |

### Get all items

```http
  GET /transcoders
```



|  Query Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `input_type` | `string` | **(Optional)** Input type of asset |
| `output_type` | `string` | **(Optional)** Output type of asset |
| `codec` | `string` | **(Optional)** Codec of asset |
| `descriptor` | `string` | **(Optional)** What is the command for |
| `page_size` | `string` | **(Optional)** Number of asset to list |
| `page` | `string` | **(Optional)** Page Number |

#### Response Type
````
[
    {
        "id": string,
        "updated_at": date,
        "created_at": date,
        "updated_by": string,
        "output_type": string,
        "input_type": string,
        "status": string,
        "codec": string,
        "multi_audio": bool,
        "multi_caption": bool,
        "descriptor": string,
        "template_command": string
    }
]
````

#### Example URI
````http
/transcoders?input_type=hls&output_type=mp4&codec=h264&descriptor=encoding
````

#### Example Response
````http
[
    {
        "id": "6439161acbd9de137c273cbf",
        "updated_at": "2023-04-14T09:00:10.223Z",
        "created_at": "2023-04-14T09:00:10.223Z",
        "updated_by": "me",
        "output_type": "mp4",
        "input_type": "hls",
        "status": "active",
        "codec": "h264",
        "multi_audio": true,
        "multi_caption": false,
        "descriptor": "encoding",
        "template_command": "coming soon"
    }
]
````

#### Handler

|  Name | Location     | Description                |
| :-------- | :------- | :------------------------- |
| GetTranscoder | /handlers/getTranscoder | To list the transcoders |

### Add Item

```http
  POST /transcoders
```



|  Request Data| Type     | Description                |
| :-------- | :------- | :------------------------- |
| `input_type` | `string` | **(Required)** Input type of asset |
| `output_type` | `string` | **(Required)** Output type of asset |
| `codec` | `string` | **(Required)** Codec of asset |
| `descriptor` | `string` | **(Required)** What is the command for |
| `updated_by` | `string` | **(Required)** Who has added |
| `multi_audio` | `bool` | **(Required)** Is Multi Audio |
| `multi_caption` | `bool` | **(Required)** Is Multi Caption |
| `template_command` | `string` | **(Required)** The command to transcode |

#### Response Type
````
string
````

#### Handler

|  Name | Location     | Description                |
| :-------- | :------- | :------------------------- |
| AddTranscoder | /handlers/addTranscoder | To add the transcoders |


### Update Item

```http
  PUT /transcoders
```



|  Request Data| Type     | Description                |
| :-------- | :------- | :------------------------- |
| `input_type` | `string` | **(Required)** Input type of asset |
| `output_type` | `string` | **(Required)** Output type of asset |
| `codec` | `string` | **(Required)** Codec of asset |
| `descriptor` | `string` | **(Required)** What is the command for |
| `updated_by` | `string` | **(Required)** Who has added |
| `multi_audio` | `bool` | **(Required)** Is Multi Audio |
| `multi_caption` | `bool` | **(Required)** Is Multi Caption |
| `template_command` | `string` | **(Required)** The command to transcode |

#### Response Type
````
string
````

#### Handler

|  Name | Location     | Description                |
| :-------- | :------- | :------------------------- |
| PutTranscoder | /handlers/putTranscoder | To update using put request |

### Modify Item

```http
  PATCH /transcoders
```
|  Query Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `input_type` | `string` | **(Required)** Input type of asset |
| `output_type` | `string` | **(Required)** Output type of asset |
| `codec` | `string` | **(Required)** Codec of asset |
| `descriptor` | `string` | **(Required)** What is the command for |


|  Request Data| Type     | Description                |
| :-------- | :------- | :------------------------- |
| `updated_by` | `string` | **(Optional)** Who has added |
| `multi_audio` | `bool` | **(Optional)** Is Multi Audio |
| `multi_caption` | `bool` | **(Optional)** Is Multi Caption |
| `template_command` | `string` | **(Optional)** The command to transcode |

#### Response Type
````
string
````

#### Handler

|  Name | Location     | Description                |
| :-------- | :------- | :------------------------- |
| PatchTranscoder | /handlers/patchTranscoder | To update using patch request |

### Delete Item

```http
  DELETE /transcoders
```



|  Query Prameters| Type     | Description                |
| :-------- | :------- | :------------------------- |
| `input_type` | `string` | **(Required)** Input type of asset |
| `output_type` | `string` | **(Required)** Output type of asset |
| `codec` | `string` | **(Required)** Codec of asset |
| `descriptor` | `string` | **(Required)** What is the command for |


#### Response Type
````
string
````

#### Handler

|  Name | Location     | Description                |
| :-------- | :------- | :------------------------- |
| DeleteTranscoder | /handlers/deleteTranscoder | To delete the transcoder |


## Requirements
### Ports Requirements

     This service uses the following ports
     1. 8001 - for the service

### Communication requirements

    This service should be reachable from AmagiNow backend.

### CPU and memory resource requirements

     cpu 1 memory 1Gi

### Building and runnning the service

    1. To build the service
        go build -o vod3customers
    2. To run the service
        ./vod3customers
    3. To run the service in docker
        docker build -t vod3customers .
        docker run -p 8001:8001 vod3customers
    4. To run the service in docker-compose
        docker-compose up
    5. for local testing please clone the application and run  
        go run main.go
   
### Liveness and readiness probe

    1. Liveness  and rediness probe is configured to check if the service is running on port 8001
        http://<host>>:8001/
  
### Metrics

    Prometheus metrics are exposed on the following endpoint
    http://<host>>:8001/metrics

### Default replicas

     Please use 1 replica in preprod and 2 replicas in prod

### Required Services

     Need MongoDB service to be running
     MongoDBURL, user and password should be provided as environment variables

### Scaling requirements

    No specific scaling requirements are identified
