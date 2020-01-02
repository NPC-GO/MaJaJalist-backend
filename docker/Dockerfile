FROM alpine:latest
LABEL maintainer="wayne900619@gmail.com"
EXPOSE 80 443
COPY ./MaJaJalist/MaJaJaserver /home/MaJaJalist/
COPY ./database/migrations/* /home/MaJaJalist/migrations/
WORKDIR /home/MaJaJalist
ENTRYPOINT ["MaJaJaserver"]