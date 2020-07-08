# temp-container
A container that returns http information to be used as a placeholder app in kubernetes deployments.

build
`docker build -t temp-container:local .`

test
```
docker run -d \
  -it \
  --name temp-container \
  temp-container:local
```

# features
- mostly copied from gobyexample.com
- runs in a google distroless container
- defaults to listen on port 8080

