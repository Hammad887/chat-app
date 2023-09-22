include Makefile.variables

.PHONY: format todo test check prepare
## prefix before other make targets to run in your local dev environment
local: | quiet
	@$(eval DOCKER_RUN= )
	@rm -rf tmp vendor
	@mkdir -p tmp
	@touch tmp/dev_image_id
quiet: # this is silly but shuts up 'Nothing to be done for `local`'
	@:

prepare: tmp/dev_image_id
tmp/dev_image_id:
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	@docker build -t ${DEV_IMAGE} -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id

clean:
	@rm -rf tmp vendor

vendor: prepare
	${DOCKER_RUN} bash -c 'go mod vendor && chmod -R 777 vendor'

format: vendor
	${DOCKER_RUN} bash ./scripts/format.sh

check: format
	${DOCKER_RUN} bash ./scripts/check.sh

codegen: prepare
	${DOCKER_RUN} bash ./scripts/swagger.sh

test: check db_prepare
	${DOCKER_TEST} bash ./scripts/test.sh

db_stop:
	bash ./scripts/db_stop.sh

build:
	bash ./scripts/build.sh

db_start: db_stop
	@docker run --name mysql_db -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3307:3306 -d mysql:latest

db_prepare: db_start
	@docker cp chat_app.sql mysql_db:chat_app.sql
	@echo "Executing databases...wait for 15 seconds"
	@sleep 15
	@docker exec -i mysql_dbsh -c 'mysql -uroot < chat_app.sql'

help:
	@echo
	@echo 'Usage: make COMMAND'
	@echo
	@echo 'Commands:'
	@echo '  build           		Compile project.'
	@echo '  check           		Run linters.'
	@echo '  format          		Format source code.'
	@echo '  codegen         		Generate code.'
	@echo '  test            		Run test case'
	@echo '  prepare         		build dev container'
	@echo '  clean       			remove dev temp folder'
	@echo
	