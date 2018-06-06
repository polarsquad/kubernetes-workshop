APP_NAME = polarsquad/kubers

all: v1

v1:
	docker build -t $(APP_NAME) .

clean:
	docker rmi $(APP_NAME)

.PHONY: all clean v1

publish: v1
	docker push $(APP_NAME)
