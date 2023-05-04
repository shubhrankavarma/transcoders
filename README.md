
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

```
  GET /transcoders
```



|  Query Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `input_type` | `string` | **(Optional)** Input type of asset |
| `output_type` | `string` | **(Optional)** Output type of asset |
| `operation` | `string` | **(Optional)** Operation of asset |
| `descrition` | `string` | **(Optional)** What is the command for |
| `page_size` | `string` | **(Optional)** Number of asset to list |
| `page` | `string` | **(Optional)** Page Number |
| `updated_by` | `string` |  **(Optional)** Updated By |
| `asset_type` | `string` | **(Optional)** Asset Type of asset |
| `operation` | `string` | **(Optional)** Operation of asset |
| `status` | `string` | **(Optional)** Status of asset |

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
        "description" : string,
        "asset_type : string,
        "operation" : string,
        "template_command": string,
    }
]
````

#### Example URI
````
/transcoders?input_type=hls&output_type=mp4&codec=h264&descriptor=encoding
````

#### Example Response
````
[
    {
        "id": "6439161acbd9de137c273cbf",
        "updated_at": "2023-04-14T09:00:10.223Z",
        "created_at": "2023-04-14T09:00:10.223Z",
        "updated_by": "me",
        "output_type": "mp4",
        "input_type": "hls",
        "status": "active",
        "description" : "encoding",
        "asset_type" : "vedio",
        "operation" : "encoding",
        "template_command": "coming soon"
    }
]
````

#### Handler

|  Name | Location     | Description                |
| :-------- | :------- | :------------------------- |
| GetTranscoder | /handlers/getTranscoder | To list the transcoders |

### Add Item

```
  POST /transcoders
```



|  Request Data| Type     | Description                |
| :-------- | :------- | :------------------------- |
| `input_type` | `string` | **(Optional)** Input type of asset |
| `output_type` | `string` | **(Optional)** Output type of asset |
| `description` | `string` | **(Optional)** What is the command for |
| `updated_by` | `string` | **(Required)** Who has added |
| `template_command` | `string` | **(Required)** The command |
| `operation` | `string` | **(Required)** Operation of asset |
| `asset_type` | `string` | **(Required)** Asset type of asset |

#### Response Type
````
string
````

#### Handler

|  Name | Location     | Description                |
| :-------- | :------- | :------------------------- |
| AddTranscoder | /handlers/addTranscoder | To add the transcoders |


### Update Item

```
  PUT /transcoders
```



|  Request Data| Type     | Description                |
| :-------- | :------- | :------------------------- |
| `input_type` | `string` | **(Optional)** Input type of asset |
| `output_type` | `string` | **(Optional)** Output type of asset |
| `description` | `string` | **(Optional)** What is the command for |
| `updated_by` | `string` | **(Required)** Who has added |
| `template_command` | `string` | **(Required)** The command |
| `operation` | `string` | **(Required)** Operation of asset |
| `asset_type` | `string` | **(Required)** Asset type of asset |

#### Response Type
````
string
````

#### Handler

|  Name | Location     | Description                |
| :-------- | :------- | :------------------------- |
| PutTranscoder | /handlers/putTranscoder | To update using put request |

### Modify Item

```
  PATCH /transcoders
```
|  Query Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `asset_type` | `string` | **(Required)** Asset Type of asset |
| `operation` | `string` | **(Required)** Operation of asset |


|  Request Data| Type     | Description                |
| :-------- | :------- | :------------------------- |
| `input_type` | `string` | **(Optional)** Input type of asset |
| `output_type` | `string` | **(Optional)** Output type of asset |
| `description` | `string` | **(Optional)** What is the command for |
| `updated_by` | `string` | **(Optional)** Who has added |
| `template_command` | `string` | **(Optional)** The command |
| `operation` | `string` | **(Optional)** Operation of asset |
| `asset_type` | `string` | **(Optional)** Asset type of asset |

#### Response Type
````
string
````

#### Handler

|  Name | Location     | Description                |
| :-------- | :------- | :------------------------- |
| PatchTranscoder | /handlers/patchTranscoder | To update using patch request |

### Delete Item

```
  DELETE /transcoders
```



|  Query Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `asset_type` | `string` | **(Required)** Asset Type of asset |
| `operation` | `string` | **(Required)** Operation of asset |


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
