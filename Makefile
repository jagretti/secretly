LOCAL_REGISTRY=localhost:5000

run-registry:
	docker run -d -p 5000:5000 --name registry registry:2.7

rm-registry:
	docker container rm registry

build:
	docker build -t secretly:latest .

push-local:
	docker tag secretly $(LOCAL_REGISTRY)/secretly:latest
	docker push $(LOCAL_REGISTRY)/secretly:latest


go-build:
	go build -o secretly

apply:
	kubectl apply --force -f deploy/deployment.yaml

delete:
	kubectl delete -f deploy/deployment.yaml
