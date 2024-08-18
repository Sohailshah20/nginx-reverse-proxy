run :
	docker compose up

stop:
	docker compose down

restart-server:
	@echo "building the go server"
	docker compose build go-server
	@echo 'restarting the go server'
	docker compose up -d --no-deps go-server



.PHONY: build