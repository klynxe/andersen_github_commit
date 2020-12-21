export MONGO_USER=admin
export MONGO_PASS=admin
export GITHUB_URL=https://api.github.com/repos/klynxe/andersen_github_commit/commits
export PERIOD_SECOND=30

run_local:
	golangci-lint run ./...
	docker-compose up -d --build