# assistant_go

1. start zookeeper and kafka server
zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties
kafka-server-start /usr/local/etc/kafka/server.properties

2. create topics for text and images

kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic textTopic
kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic fileTopic

3. create topics for response text and images

kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic textTopicResponse
kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic fileTopicResponse

3. create producers for text and images

kafka-console-producer --broker-list localhost:9092 --topic textTopic
kafka-console-producer --broker-list localhost:9092 --topic fileTopic

4. install go dependecies
go get -u github.com/confluentinc/confluent-kafka-go/kafka

5. run file and image go file
go run image_consumer.go
go run text_consumer.go
