.PHONY: run stop shutdown

run:
	docker compose --env-file .env up -d --build

stop:
	docker compose --env-file .env down

shutdown:
	docker compose --env-file .env down --volumes
