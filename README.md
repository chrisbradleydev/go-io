# go-io

Clone the repository.

```sh
git clone https://github.com/chrisbradleydev/go-io.git .
```

Build Docker image.

```sh
docker build -t go-io .
```

Run Docker container.

```sh
docker run -v $(pwd):/app -v /app/tmp go-io
```
