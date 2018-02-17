[Terraform](https://www.terraform.io) provider for [Kafka Connect](https://docs.confluent.io/current/connect/intro.html)

# Example

```hcl
provider "kafkaconnect" {
  url = "http://kafka_connect_url:PORT"
}

resource "kafkaconnect_connector" "readFromFile" {

  name = "readFromFile"

  config = {
    "name" = "local-file-sink"
    "connector.class" = "FileStreamSink"
    "tasks.max" = "1"
    "file" = "/tmp/readFile.txt"
    "topics" = "deliverToTopic"
  }

}
```
