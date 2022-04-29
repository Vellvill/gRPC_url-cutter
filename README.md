# gRPC_url-cutter server
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
2. ```docker-compose run -p 8080:8080 server [-FLAGS]```
2.1. If you want to use your own database just edit dsn in ```config.yaml```
3 Try to add new links in database. For gRPC client you can use evans https://github.com/ktr0731/evans or you can use gRPC_url-cutter client.

# gRPC_url-cutter client
## Implementation
Client implements generated interface by protocol buffers.
```golang
type URLShortenerClient interface {
	Create(ctx context.Context, in *CreateURLRequest, opts ...grpc.CallOption) (*CreateURLResponse, error)
	Get(ctx context.Context, in *GetURLRequest, opts ...grpc.CallOption) (*GetURLResponse, error)
}
```
You can send requests directly to server by using flags ```-add``` and ```-get```.
## How to use
1. start your server with docker and make sure that you used flags for repository that you want to use.
2. ```run ./internal/client [-FLAG] ["VALUE"]```
3. get an answer.
