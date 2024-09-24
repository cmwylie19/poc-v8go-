.PHONY: create delete build import k8s-deploy run-all k8s-events k8s-logs k8s-delete redeploy k8s-exec

delete:
	k3d cluster delete k3s-default
	# docker rmi poc-v8go

create:
	k3d cluster create k3s-default

build:
	docker build -t poc-v8go . -f Dockerfile.go

import:
	k3d image import poc-v8go -c k3s-default

k8s-deploy:
	kubectl apply -f k8s

k8s-events:
	kubectl get ev --sort-by='.lastTimestamp'

k8s-logs:
	kubectl logs -l run=controller

k8s-delete:
	kubectl delete po controller --force
	kubectl delete cm module
	
docker-exec:
	docker run -it --entrypoint sh poc-v8go  
	   
redeploy: k8s-delete k8s-deploy

run-all: delete create build import k8s-deploy

k8s-exec:
	 kubectl exec -it controller -- sh
