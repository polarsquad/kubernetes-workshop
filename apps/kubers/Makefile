APP_NAME = polarsquad/kubers

all: build

build:
	docker build -t $(APP_NAME) .

clean:
	docker rmi $(APP_NAME)

.PHONY: all clean build

publish: build
	docker push $(APP_NAME)