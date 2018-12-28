#!/bin/bash
docker run -d \
-p 9090:8080 \
--name swagger-ui \
-e SWAGGER_JSON=/json/helloworld.swagger.json \
-v /Users/xiaoshenge/workcode/go/src/go-example/rpc/grpc/examples/helloworld/helloworld:/json swaggerapi/swagger-ui