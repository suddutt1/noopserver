## Noop server 

### After download 

```sh

go get github.com/gin-gonic/gin

```
### Docker image build instructions


```sh

cp ~/go/src/github.com/suddutt1/noop/*.go .
docker build -t noopserver:2.0.0 .

```

