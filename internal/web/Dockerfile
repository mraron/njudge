ARG PROJECT_NAME
FROM ${PROJECT_NAME}-base

COPY configs/docker/web_docker.json ./web.json

CMD ["./njudge", "web"]