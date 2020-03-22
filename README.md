# cyticoffee-api

##mysql
docker run --name cyticoffee -p 13306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d

##redis
docker run --name redistesting -d redis

##rabbitmq
https://www.rabbitmq.com/download.html
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management