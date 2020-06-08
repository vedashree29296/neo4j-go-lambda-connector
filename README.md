# Neo4j-Go-lambda-connector
Connect to Neo4J using golang, serverless and aws lambda

## Installation and Deployments
### Requirements:
- Docker service should be running

- Make sure this code base is in $(GOROOT)/src directory

- Ensure AWS credentials are configured for the particular account/IAM Role

- Install dep package

``` sudo apt install go-dep```

- Install serverless

``` npm install -g serverless```


## Building and Deploying to serverless

- Edit the serverless.yml file according to requirements

- Build the package using make

``` cd neo4j-lambda-connector```

```make clean```

``` make build ```

```serverless deploy```

Or simply run the script:

``` /bin/bash runserverless.sh```

## Troubleshooting

- Deployment issues can be due to permissions set on the bin/ folder. Set the execute permission on the bin folder if it is not set already using:

``` chmod 777 -R bin```
