run:
	 cd deployments/ && docker-compose up -d
build:
	 cd deployments/ && docker-compose up --build
stop:
	 cd deployments/ && docker-compose down -v


