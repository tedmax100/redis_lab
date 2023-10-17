#!/bin/bash

# launch services
docker compose up -d
sleep 10

# add slots from 1 to 16383 to cluster
docker exec -it redis-node1 redis-cli CLUSTER ADDSLOTS $(seq 0 16383)

# add node2 to cluster
docker exec -it redis-node1 redis-cli CLUSTER MEET 172.25.0.11 6379

# let node2 to be a replica node
myid=$(docker exec redis-node1 redis-cli cluster myid | tr -d '"')
docker exec redis-node2 redis-cli CLUSTER REPLICATE $myid