# RabbitMQ in Golang
`RabbitMQ` is one of the most popular open-source, lightweight message brokers which helps in transferring data based on a messaging protocol. `RabbitMQ` supports multiple messaging protocols, and one of the most commonly used underlying protocol to route messages is `Advanced Messaging Queueing Protocol (AMQP)`.

`AMQP` has three primary entities – namely: `Queues`, `Exchanges` and `Bindings`. `AMQP` being a programmable protocol, provides flexibility to configure these entities to help transfer messages from producer to consumer.

When a publisher publishes a message to `RabbitMQ`, it is first received at an exchange. The exchange routes the message to variously connected queues which is then received by the consumers.

## Set up a RabbitMQ instance locally
To get started, we’ll first need to get a `RabbitMQ` instance running on our machine that we can interact and play with it. The easiest and quickest approach is to use the docker run command and specifying the image name that we want to run.
```
docker run -d --hostname rabbit-in-go --name a-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management 
```
This gets our `RabbitMQ` instance locally with its management interface available at `http://localhost:15672`.