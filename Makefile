run:
	go build -o main && ./main

build: ## Build the docker image
	docker build -t registry/app:latest .

push: ## Push the docker image
	docker push registry/app:latest .

test:
	docker run registry/app:latest .

