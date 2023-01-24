ARG PROJECT_NAME
FROM ${PROJECT_NAME}-base

COPY configs/docker/glue_docker.json ./glue.json

CMD ["./njudge", "glue"]