# Grocery Store Produce API

## API Endpoints & Examples

Below is a list of the current API endpoints and some examples on how to interact with them.

### GET

**/produce/getall**

This returns a JSON object of all the produce items in the DB.
Example Usage (CURL):

```
curl -X "GET" http://localhost:8080/produce/getall
```
Or you can paste the URL into a browser.

**/produce/getitem**

This returns the JSON object of one or more produce items. Pass in the produce code using the 'code' URL paramater, repeat this variable to get multiple produce items

Example Usage (CURL) for one produce item:

```
curl -X "GET" "http://localhost:8080/produce/getitem?code=A12T-4GH7-QPL9-3N4M"
```

Example Usage (CURL) for two produce items:
```
curl -X "GET" "http://localhost:8080/produce/getitem?code=A12T-4GH7-QPL9-3N4M&code=E5T6-9UI3-TH15-QR88"
```
Or, of course, you may paste these URL's into a browser instead.

### POST

**/produce/add**

This allows the caller to add one or more produce items. The API expects a JSON payload.

Example (CURL) for adding "Banana" and "Red Onion" produce items:

```
curl -X "POST" http://localhost:8080/produce/add -H "Accept: application/json" -H "Content-Type: application/json" -d '{"Produce":[{"Name":"Banana","ProduceCode":"B123-4G2A-QPL9-3N2A","Price":1.19},{"Name":"Red Onion","ProduceCode":"1234-ABCD-QPL9-3N4M","Price":2.24}]}'
```

On the event of an error (such as using a code that doesn't follow the correct format structure as seen in the below example), you will get a response such as:

```
Could not add the item: B123-4G2A-QPL
Could not add the item: 1234-ABCD-QP4M
```

To create your own JSON objects and to compact them into one line, a good tool to use is: https://codebeautify.org/jsonminifier

### DELETE

**/produce/delete**
This allows the caller to delete one or more produce items. Per the specifications, the expected URL parameter is "Produce Code". When using curl, the space between "Produce" and "Code" should be the string "%20".

Example (CURL) to delete one item:

```
curl -X "DELETE" "http://localhost:8080/produce/delete?Produce%20Code=TQ4C-VV6T-75ZX-1RMR"
```

Example (CURL) to delete two items:

```
curl -X "DELETE" "http://localhost:8080/produce/delete?Produce%20Code=TQ4C-VV6T-75ZX-1RMR&Produce%20Code=A12T-4GH7-QPL9-3N4M"
```

## How to Run

The simplest way to run the app is to pull the latest image from docker:

```
docker run --rm -p 8080:8080 murphylsou/grocery-store-api
```

You can also specify a custom port by setting a PORT environment variable for the container:

```
export PORT=3030
docker run --rm -p $PORT:$PORT --env PORT=$PORT murphylsou/grocery-store-api
```

To run locally without a docker container, in the root directory of the project:

```
go build && ./grocery-store-api
```

Or to build and run a docker container locally:

```
docker build -f Dockerfile -t grocery-store-api .
docker run -it -p 8080:8080 grocery-store-api
```