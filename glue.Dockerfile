ARG PROJECT_NAME
FROM ${PROJECT_NAME}-base

COPY configs/docker/glue.yaml ./glue.yaml

CMD ["./njudge", "glue"]