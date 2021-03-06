# terraform-provider-kafka-connect

[Terraform](https://www.terraform.io) provider for [Kafka Connect](https://docs.confluent.io/current/connect/intro.html)

## Installing

Install latest version:
```
go get -u github.com/wakeful/terraform-provider-kafka-connect
```
or download binary from [release](https://github.com/wakeful/terraform-provider-kafka-connect/releases) page.

Register plugin in your local `.terraformrc`:
```
cat >> ~/.terraformrc <<EOF
providers {
  kafkaconnect = "${GOPATH}/bin/terraform-provider-kafka-connect"
}
EOF
```

## Example

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
