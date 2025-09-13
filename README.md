
## Requirements


### Functional Requirements
1. Generate and Store Shorten Url
2. Configure Sorten Url Alias
3. Redirect to Actual Url
4. Expiry

### Non Functional Requirements
1. High Availability
2. Low Latency
3. Scalability
4. Durability
5. Security


## HLD

![System Design HLD](docs/images/Diagram%201.svg)


## Database Design
1. No SQL (Dynamo DB, Casanda )
2. Schema

**Url Mapping**

| short_url   | varchar   |
| ----------- | --------- |
| url         | varchar   |
| created_on  | timestamp |
| expire_on   | timestamp |
| user_id     | uuid      |
| click_count | int       |
**User**

| user_id  | uuid    |
| -------- | ------- |
| email    | varchar |
| name     | varchar |
| password | varchar |

## API Design

```go title="URL Shortening"
Endpoint: /api/v1/shorten
Method: POST
Request:
{
	long_url: string
	alias: string
	expiry_time: int
}
Response:
{
	short_url: string
	long_url: string
	created_on: int
	expire_on: int
}

/*
1. Generate Some Unique ID(Auto Increment,UUID, Hash)
2. Encode using Base62
*/
```


```go title="URL Redirection"
Endpoint: {short_url}
Method: GET
Response:
	Code: 301
	Location: {long_url}
```



