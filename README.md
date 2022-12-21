# API Template

## Style Guide

We mostly follow [Uber Style Guide](https://github.com/uber-go/guide/blob/master/style.md).

## Architecture & code organization

> Note: would change over time.

```sh
.
├── template.yaml              # AWS SAM Template
├── Makefile                   # Makefile
├── handlers                   # Main application commands & entry point
│   └── <lambda-name>          # Contains 'main.go' lambda handler
├── migrations                 # Database migrations
├── models                     # database entities (including their specific enum interfaces)
├── repositories               # database access layer
├── config                     # configurations and env variable/setup
└── pkg                        # 3rd party lib wrappers and/or utility functions. (Utilities should be extracted)
    ├── utils                  # utility library.
    └── db                     # database library.
        └── migrations         # utilities for database migrations
```

> NOTE: 'dev' instances of database are used for Local Development.

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
* [watchexec](https://github.com/mattgreen/watchexec)

You may need the following for local testing.
* [Golang](https://golang.org)

Install the CLI tools and Docker CE

```sh
brew install watchexec
```

### Testing

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests locally:

```shell
go test -v ./handlers/<lambda-name>/
```

## Engineering Practices

- Code defensively: never make assumption about the data in the database or api response.
- Data Duplication, denormalizing data: document stores are allowed to denormalize data.
    - Duplicated data changes can be enforced.

---
* [AWS Serverless Application Repository](https://aws.amazon.com/serverless/serverlessrepo/)
