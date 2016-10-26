all: push

# 0.0 shouldn't clobber any released builds
TAG =1.2cron
PREFIX = remotejob/huoneisto_utils

binary: app.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o app

container: binary
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

set: 
	 kubectl set image deployment/ huoneisto_utils huoneisto_utils=$(PREFIX):$(TAG)

clean:
	docker rmi -f $(PREFIX):$(TAG) || true
