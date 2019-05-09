# Elasticsearch Golang Service
Simple search app service built in go using elasticsearch. This app was built following the tutorial [How to Build a Search Service with Go and Elasticsearch](https://outcrawl.com/go-elastic-search-service)

## Instalation

Clone this repository
```
$ git clone https://github.com/leogsouza/elasticsearch-go.git
```

Then build and run the services
```
$ docker-compose up -d --build
```

# Usage
The API has two endpoints:
- **POST** ```/document``` - To indexing your documents
- **GET** ```/search?query=xxx``` - To search your documents

For example you can insert some fake documents using the fake-data.json file contained in this repository
```
$ curl -X POST http://localhost:8080/documents -d @fake-data.json -H "Content-Type: application/json"
```

For search documents you can try the following command or make a request using your preferente REST API client

```
$ curl http://localhost:8080/search?query=laboris+tempor+ullamco
{
  "time": "84",
  "hits": "42",
  "documents": [
    {
      "title": "Nisi laboris anim ipsum ullamco Lorem magna velit est ex non tempor irure ut.",
      "created_at": "2019-05-09T00:35:59.626409893Z",
      "content": "In incididunt irure id deserunt labore reprehenderit elit magna...",
    }
    // Other documents
  ]
}
```