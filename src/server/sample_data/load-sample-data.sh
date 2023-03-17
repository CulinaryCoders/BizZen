#!/bin/bash

#  UPDATE VARIABLES WITH APPROPRIATE VALUES
INBOUND_JSON_PATH=./test-addresses2.json
DB_HOST=localhost
DB_PORT=8420
API_ENDPOINT=address

for line in $(cat $INBOUND_JSON_PATH)
do
    curl -d @$line -H "Content-Type: application/json" http://$DB_HOST:$DB_PORT/$API_ENDPOINT
done