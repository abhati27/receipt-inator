
IMAGE_NAME = receipt-inator
PORT = 8080

build:
	docker build -t $(IMAGE_NAME):latest .


run:
	docker run -p $(PORT):$(PORT) $(IMAGE_NAME):latest