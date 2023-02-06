kafka-topics.sh --bootstrap-server kafka:9092 --create --topic boom
kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic boom --from-beginning

kafka-console-producer.sh --bootstrap-server kafka:9092 --topic boom

docker-compose run -e topic=boom -e kafkaURL=kafka -w /golang go go run ./cmd/producer.go



producer
topic=boom kafkaURL=kafka go run ./cmd/producer/producerd.go
consumer
topic=boom kafkaURL=kafka go run ./cmd/consumer/consumerd.go


protco generation:
```bash
protoc -I=./protos/ --go_out=./out ./protos/*.proto
```