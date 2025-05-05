# benchmark core ms publish with 10 publishers on subject foo
ms-client bench pub foo --clients 10 --msgs 10000 --size 512

# benchmark core ms subscribe for 4 clients on subject foo
ms-client bench sub foo --clients 5 --msgs 10000

# benchmark core ms request-reply with queuing
## run 4 clients servicing requests
ms-client bench service serve --clients 4 testservice

## run 4 clients making synchronous requests on the service at subject testservice
ms-client bench service request --clients 4 testservice --msgs 20000

# benchmark JetStream asynchronously acknowledged publishing of batches of 1000 on subject foo creating the stream first
ms-client bench js pub foo --create --batch 1000

# benchmark JetStream synchronous publishing on subject foo using 10 clients and purging the stream first
ms-client bench js pub foo --purge --batch=1 --clients=10

# benchmark JetStream delivery of messages from a stream using an ephemeral ordered consumer, disabling the progress bar
ms-client bench js ordered --no-progress

# benchmark JetStream delivery of messages from a stream through a durable consumer shared by 4 clients using the Consume() (callback) method.
ms-client bench js consume --clients 4

# benchmark JetStream delivery of messages from a stream through a durable consumer with no acks shared by 4 clients using the fetch() method with batches of 1000.
ms-client bench js fetch --clients 4 --acks=none --batch=1000

# simulate a message processing time of 50 microseconds
ms-client bench service serve testservice --sleep 50us

# generate load by publishing messages at an interval of 100 nanoseconds rather than back to back
ms-client bench pub foo --sleep=100ns

# remember when benchmarking JetStream
Once you are finished benchmarking, remember to free up the resources (i.e. memory and files) consumed by the stream using 'ms-client stream rm'.

You can get more accurate results by disabling the progress bar using the `--no-progress` flag.