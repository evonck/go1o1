FROM debian:7

RUN apt-get -y update \

&& apt-get install -y build-essential mercurial git wget curl \

#Set up folders

&& mkdir /srv/app/ && mkdir /srv/conf/ && mkdir /srv/logs/ \
# env vars

# go 1.3 tarball
&& wget -qO- http://golang.org/dl/go1.4.2.linux-amd64.tar.gz | tar -C /srv/app -xzf - 

WORKDIR /srv/app/go/src
RUN ./make.bash  

ENV GOBIN /srv/app/projet/bin

RUn mkdir /srv/app/projet

ENV GOPATH /srv/app/projet/

ENV GOROOT /srv/app/go

ENV PATH /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/srv/app/go/bin:/srv/app/projet/bin

RUN go get github.com/tools/godep

CMD godep restore  && godep go build

