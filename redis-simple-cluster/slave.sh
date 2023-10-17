#!/bin/bash

# add slots from 1 to 16383 to cluster
myid=$(docker exec redis-node1 redis-cli cluster myid | tr -d '"')
docker exec redis-node2 redis-cli CLUSTER REPLICATE $myid