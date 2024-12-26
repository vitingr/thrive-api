run:
	air

install-prometheus:
	docker-compose -f docker-compose.yml up -d

down-prometheus:
	docker-compose -f docker-compose.yml down
