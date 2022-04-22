﻿# gRPC_url-cutter
## gRPC 
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
