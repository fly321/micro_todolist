# 微服务项目
## user-service
> config.ini
```env
[service]
AppMode = debug
HttpPort = :4000

[mysql]
Driver = mysql
DbHost = host.docker.internal
DbPort = 3306
DbUser = root
DbPassword = 123456
DbName = micro_todo_list
Charset = utf8mb4
MaxIdleConns = 10
MaxOpenConns = 100

[rabbitmq]
RabbitMQ = amqp
RabbitMQUser = guest
RabbitMQPassWord = guest
RabbitMQHost = localhost
RabbitMQPort = 5672

[etcd]
EtcdHost = localhost
EtcdPort = 2379

[server]
UserServiceAddress = 127.0.0.1:8082
TaskServiceAddress = 127.0.0.1:8083

[redis]
RedisHost = localhost
RedisPort = 2379
RedisUsername = micro_todolist
RedisPassword =
```