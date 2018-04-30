FROM daocloud.io/golang
RUN mkdir /gopath && mkdir /tmp/file
ENV GOPATH /gopath
ADD . /gopath/
RUN cd /gopath/src/main/ && go build main.go && mv main /go && rm -rf /gopath
EXPOSE 8989
CMD ["./main"]
