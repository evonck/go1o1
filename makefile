# Makefile
#
# Targets (see each target for more information):
#   dockerize:    		builds, tests, then makes a Docker image for each binary
#	  clean:			  		removes build artifacts (aka binaries)
#	  clean-all:	  		cleans, then removes the artifact for build/container

buildDocker: 
	@docker build -t todo-docker docker

#Create the binary inside of the docker
#mount the binary on the folder call todo
install: buildDocker
	@docker run --rm -v "$$PWD":"/srv/app/projet/${SRC_ROOT}" -w "/srv/app/projet${SRC_ROOT}" \
	-t todo-docker 

dockerize: install
	@docker run --rm -v "$$PWD":"/srv/app/projet/${SRC_ROOT}" -w "/srv/app/projet/${SRC_ROOT}" \
	-p 8081:8081 -t todo-docker ./todo "${MYSQL}"

#Create the test
test: install
	@docker run --rm -v "$$PWD":"/srv/app/projet/${SRC_ROOT}" -w "/srv/app/projet/${SRC_ROOT}" \
	-e MYSQL_TEST_ENV="${MYSQL_TEST_ENV}" -t todo-docker godep go test -v ./...
	
clean: 
	rm -f todo

