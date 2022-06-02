compose:
	docker compose up

api:
	docker build -t wape-api ./api

parser:
	docker build -t wape-parser ./parser

beanstalkd:
	beanstalkd -l 127.0.0.1 -p 11300