# slowserver

Simple webserver echoing back the request information. 

### Usage

```
Usage of slowserver:
  -d int
    	Delay before responding
  -p string
    	Port to listen on (default "8080")
  -r string
    	response so answer with
```


### Example using curl to make the request with -r asdf

```
$ curl localhost:8080
asdf
```
