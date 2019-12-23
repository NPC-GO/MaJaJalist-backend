FROM alpine:latest
LABEL maintainer="wayne900619@gmail.com"
EXPOSE 80 443
COPY ./MaJaJalist/MaJaJaserver /home/MaJaJalist/
COPY ./migrations/* /home/MaJaJalist/migrations/
ENTRYPOINT ["/home/MaJaJalist/MaJaJaserver"]