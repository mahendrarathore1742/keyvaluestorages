# key value store

This is Simple Key value store web service with a subscription feature to pull changes happening to any of the keys. Also has a CLI client which supports the set value and get value along with watch operation to watch changes.

## Project Overview

This project can be divided into steps:
- Build the key value store server and write endpoints for GET and SET operations.
- Setup a real-time connection with client to send live changes.
- Build a CLI client that consumes the web service and facilitates the required operations.

### The CLI client

Going with [spf13/cobra](https://github.com/spf13/cobra) is an easy choice when building a CLI application. I could easily add stuff like persistent flags and such with the help of it.

The code can be found in `/Keyvaluecli`.

## How to Use This application

The requirement to run the storage server is Docker. It can be run as follows,
```sh
# Bring up the Kafka containers
docker-compose up --build

# Run the server
# --network="host" allows the container to access host's network
# this means that it can connect to Kafka easily as compared to
# setting up a bridge network inside Docker
docker run -it --rm --network="host" mahendrarathore/keyvaluestorage:latest
```

Once the server is up and running, we can make requests using `curl`,
```sh
# SET operation
curl localhost:3000/set -d '{"Key": "Name", "Value": "Mahendra"}'

# GET operation
curl localhost:3000/get -d '{"Key": "Name"}'

# GET all key-value pairs
curl localhost:3000/getall
```

To install the CLI client
 
```

go install github.com/mahendrarathore1742/keyvaluestorage/keyvaluecli@latest


```

Once installed, use the CLI as follows,
```sh
# GET all keys
keyvaluecli --server-address localhost:3000 get --all

# GET a given key-value pair
keyvaluecli --server-address localhost:3000 get --key "Name"

# SET a given key-value pair
keyvaluecli --server-address localhost:3000 set --key "Name" --value "Mahendra"

# Watch changes
keyvaluecli --server-address localhost:3000 watch
```

## References

- https://github.com/segmentio/kafka-go/tree/main/examples
- https://github.com/gorilla/websocket/tree/master/examples
- https://github.com/wurstmeister/kafka-docker
- https://github.com/conduktor/kafka-stack-docker-compose
- https://cobra.dev/