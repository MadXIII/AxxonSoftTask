# HTTP server for proxying HTTP-requests

Server with 3 routes:

```
"/" - get parameters from req.Body and send request to other services with this parameters and response it to client in JSON format;

"/request?id=<uniq hash>" get data from server by ID;

"/response?id=<uniq hash>" get data from server by ID.
```

PS: If Content-length is -1 it's means response service don't have content-length in their Header parameters.
