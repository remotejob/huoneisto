all: push

# 0.0 shouldn't clobber any released builds
TAG =1.2cron
PREFIX = gcr.io/jntlserv0/huoneisto_utils

binary: app.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o app

container: binary
	docker build -t $(PREFIX):$(TAG) .

push: container
	gcloud docker push $(PREFIX):$(TAG)

set: push
	 kubectl set image deployment/godocker godocker=$(PREFIX):$(TAG)

clean:
	docker rmi -f $(PREFIX):$(TAG) || true
