COMPOSE_FILE = docker-compose.yml
DOCKER_COMPOSE = docker-compose -f $(COMPOSE_FILE)

.PHONY: help run install-prometheus down-prometheus clean

help:
	@echo "Comandos disponíveis:"
	@echo "  make run               - Executa o projeto usando air"
	@echo "  make install-prometheus - Sobe os serviços do Prometheus com Docker Compose"
	@echo "  make down-prometheus   - Derruba os serviços do Prometheus"
	@echo "  make clean             - Remove recursos Docker criados"

run:
	air

install-prometheus:
	$(DOCKER_COMPOSE) up -d

down-prometheus:
	$(DOCKER_COMPOSE) down

clean:
	$(DOCKER_COMPOSE) down --volumes --remove-orphans
	docker image prune -f
	