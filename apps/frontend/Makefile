APP_NAME = polarsquad/example-frontend

all: v2

v2:
	docker build -t $(APP_NAME):v2 --build-arg version=2 .

clean:
	docker rmi $(APP_NAME):v2

.PHONY: all clean v1 v2

publish: v2
	docker push $(APP_NAME):v2