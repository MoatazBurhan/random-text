DOCKER_USERNAME=mburhan
APPLICATION_NAME=random-text
GIT_HASH=$(shell git log --format="%h" -n 1)

build:
	docker build -f docker/Dockerfile --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest .

deploy:
	kubectl create namespace random-text > /dev/null 2>&1 || true
	kubectl apply -f k8s/deployment.yaml -n random-text
	kubectl apply -f k8s/service.yaml -n random-text
	kubectl apply -f k8s/ingress.yaml -n random-text