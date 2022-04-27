# gRPC_url-cutter
## gRPC functions
```golang
rpc Create(CreateURLRequest) returns (CreateURLResponse){}
```
Sends a message as a long link, processes it on the service and returns a short link

```golang
rpc Get(GetURLRequest) returns (GetURLResponse){}
```
Takes a short link and returns a long link from the service.
## Repositories 
```golang
type in-memory struct {
	*sync.RWMutex
	inc          int64
	hashmapShort map[string]*model.Model
	hashmapLong  map[string]*model.Model
}
```
In-memory storage as two mappings, two mappings to avoid non-unique short links for long links
```sql
url(
    id SERIAL PRIMARY KEY NOT NULL ,
    longurl VARCHAR UNIQUE NOT NULL,
    shorturl VARCHAR NOT NULL
)
```
Scheme of postgres repository.
## Flags
```
-migrations creates a new table if it doesn't exist. 
If you are running the service on the local machine for the first time and want to use the database, you need to enable this flag.
```
```
-cache enable using cache instead of postgres
```
## How to run
1. Clone repository
2. docker-compose run -p 8080:8080 server [...flags]
3. Try to add new links in database. For gRPC client you can use evans https://github.com/ktr0731/evans.
