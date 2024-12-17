LOCAL_REGISTRY=localhost:5000

run-registry:
	docker run -d -p 5000:5000 --name registry registry:2.7

build:
	docker build -t secretly:latest .

push-local:
	docker tag secretly localhost:5000/secretly:latest
	docker push localhost:5000/secretly:latest


go-build:
	go build -o secretly

apply:
	kubectl apply -f deploy/deployment.yaml

delete:
	kubectl delete -f deploy/deployment.yaml