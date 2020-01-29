# stayway-media-api

``` sh
cp config.yaml.example config.yaml
docker-compose up
```

## Test

``` sh
cp config.yaml.example config.test.yaml
sed -i 's|/stayway?|/stayway_test?|' config.test.yaml
echo "create database if not exists stayway_test" | make mysql-cli

docker-compose run --rm app make test

# 一部のテストのみ実行
docker-compose run --rm app make test TARGET=./pkg/application/service
```
