docker pull localstack/localstack

docker container run -it -d -p 4566:4566 localstack/localstack

aws --endpoint-url=http://localhost:4566 sns create-topic --name cadastro
