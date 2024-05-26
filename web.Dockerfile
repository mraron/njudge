ARG PROJECT_NAME
FROM ${PROJECT_NAME}-base

COPY configs/docker/web.yaml ./web.yaml

CMD ["./njudge", "web"]