# Get github commits and save to DB every set delay
## Run local
export GITHUB_TOKEN=XXXXX
make run_local


## ENV
MONGO_USER=admin

MONGO_PASS=admin

GITHUB_URL=https://api.github.com/repos/klynxe/andersen_github_commit/commits

GITHUB_TOKEN=XXXXX

PERIOD_SECOND=30