include Makefile.variables

prepare: tmp/dev_image_id
tmp/dev_image_id:
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	@docker build -t ${DEV_IMAGE} -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id

tmp/vendor-installed: tmp/dev_image_id go.mod
	@mkdir -p vendor
	${DOCKRUN} go mod tidy
	${DOCKRUN} go mod vendor
	@date > tmp/vendor-installed
	@chmod 644 go.sum || :

format: tmp/vendor-installed
	${DOCKRUN} bash ./scripts/format.sh
codedocs: prepare
	${DOCKRUN} bash ./scripts/swagger.sh
check: format
	${DOCKRUN} bash ./scripts/check.sh
build:
	bash ./scripts/build.sh
test: mysql
	${DOCKTEST} bash ./scripts/test.sh
mysql: db_stop
	docker run -d --name mysql_dev  -p 3306:3306 mysql:latest
db_stop:
	bash ./scripts/db_stop.sh
help:
	@echo
	@echo 'Usage: make COMMAND'
	@echo
	@echo 'Commands:'
	@echo '  build           		Compile project.'
	@echo '  check           		Run linters.'
	@echo '  format          		Format source code.'
	@echo '  codedocs         		docserate code.'
	@echo '  test            		Run test case'
	@echo '  prepare         		build dev container'
	@echo '  clean       			remove dev temp folder'
	@echo
	