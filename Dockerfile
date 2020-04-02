FROM golang:alpine as build
COPY . /go/src
RUN cd /go/src \
	&& go build -o app \
	&& cp app /bin/

FROM alpine
COPY --from=build /bin/app /bin
EXPOSE 8080
CMD ["/bin/app"]