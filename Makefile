.PHONY: up
up:
	./scripts/compose.sh up

.PHONY: down
down:
	./scripts/compose.sh down

.PHONY: clean
clean:
	./scripts/compose.sh clean
