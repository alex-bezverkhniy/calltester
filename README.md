# CallTester

<!--toc:start-->
- [CallTester](#calltester)
  - [Getting Started](#getting-started)
    - [Manual Installation](#manual-installation)
      - [Requirements](#requirements)
  - [Features](#features)
  - [Example](#example)
  - [Contributing](#contributing)
<!--toc:end-->

CallTester CLI is a user-friendly, modern command-line multipurpose client. It makes calling and testing modern APIs as simple as running one command.
Supports HTTP and Kafka-based APIs.



## Manual Installation

### Requirements
- Go version `1.23`
- Make (GNU Make)

You can install CallTester CLI by using the following commands:
Clone the repository: 
```bash
git clone git@github.com:alex-bezverkhniy/calltester.git
```
Build the binary:
```bash
make build
```
Make a symbolic link to the binary or add `./bin` folder to the `$PATH` variable
```bash
sudo ln -s $(pwd)/bin/calltester /usr/local/bin

```

## Usage

### Work with Kafka
You can use CallTester CLI to work with Kafka by using the following commands:

#### Commands
```
Usage:
  calltester kafka [flags]
  calltester kafka [command]

Available Commands:
  pub         Publish message to Kafka
  sub         Subscribe to Kafka topic

Flags:
  -h, --help           help for kafka
      --host string    Hostname of the Kafka broker (default "localhost")
      --port int       Port of the Kafka broker (default 9092)
      --topic string   Kafka topic name (default "test.topic")
```

#### Produce message
```bash
calltester kafka pub --host localhost --port 9092 --topic my-topic --data 'simple message'
```

#### Subscribe to topic
```bash
calltester kafka sub --host localhost --port 9092 --topic my-topic
```


### Perform an HTTP call

#### Commands
```
Usage:
  calltester http [flags]
  calltester http [command]

Available Commands:
  delete      send DELETE request
  get         send GET request
  head        send HEAD request
  patch       send PATCH request
  post        send POST request
  put         send PUT request

Flags:
  -d, --data string          Request data (default "{\"test\": \"test\"}")
  -H, --header stringArray   HTTP header (default [Accept: */*])
  -h, --help                 help for http
  -m, --method string        Method of the request (GET, POST, PUT, DELETE, PATCH, HEAD) (default "GET")
  -p, --proxy string         Proxy URL
  -u, --url string           URL to send requests to

Global Flags:
  -v, --verbose   verbose output

```

#### Simple GET call
```bash
calltester http GET https://jsonplaceholder.typicode.com/users
```
or simply

```bash
calltester http https://jsonplaceholder.typicode.com/users
```

### Examples

Defining HTTP headers:
```bash
calltester http https://jsonplaceholder.typicode.com/users -H "Accept: application/json"
```

Sending JSON data:
```bash
calltester http POST https://jsonplaceholder.typicode.com/posts \
-H "Content-Type: application/json" \
-data '
{
  "userId": 1,
  "id": 1,
  "title": "sample title",
  "body": "sample body"
}'
```

## Features

- [x] HTTP API calls (POST, GET, PUT, PATCH, DELETE etc.)
- [x] Kafka subscriber and producer
- [x] Build-in JSON support
- [ ] WebSocket API calls `TODO`
- [ ] GraphQL API calls `TODO`

## Example

## Contributing
