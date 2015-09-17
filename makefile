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
	@docker run --rm -v "$$PWD":"/srv/app/${SRC_ROOT}" -w "/srv/app/${SRC_ROOT}" \
	--name dockerBuild -t todo-docker godep go build

dockerize: install
	@docker run --rm -v "$$PWD":"/srv/app/${SRC_ROOT}" -w "/srv/app/${SRC_ROOT}" \
	-t todo-docker ./todo ${Mysql}

#Create the test
test: buildDocker
	@docker run -i -v "$$PWD":"/srv/app/${SRC_ROOT}" -w "/srv/app/${SRC_ROOT}" \
	-t todo-docker godep go test 
clean: 

