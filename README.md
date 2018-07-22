# Gin-Example

Gin boilerplate with ORM, graceful restart support.

## package

1. gin-gonic/gin as web framework
2. jinzhu/Gorm for ORM
3. fvbock/endless for graceful restart

## Run the Application

Use `dep` for package management.

```
brew install dep # install the dep if needed
cd ./gin-example && dep ensure
```

run the code & try to restart

```
go build && ./gin-example # build & start the application

kill -SIGHUP $pid # graceful restart by pid
```
