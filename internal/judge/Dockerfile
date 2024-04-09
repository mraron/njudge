ARG PROJECT_NAME

FROM ubuntu:22.04 as judge_deps
RUN apt-get update && apt-get install -y wget gcc g++ git build-essential libcap-dev
WORKDIR /app
RUN git clone --depth 1 --branch v1.10.1 https://github.com/ioi/isolate.git
WORKDIR /app/isolate
RUN make isolate

FROM ${PROJECT_NAME}-base

COPY --from=judge_deps /app/isolate /app/isolate
WORKDIR /app/isolate
RUN make install

COPY --from=julia:1.10.1 /usr/local/julia/ /usr/local/julia
COPY --from=nimlang/nim:1.6.18 /usr/bin/nim /usr/bin/nim
RUN ln -s /usr/local/julia/bin/julia /usr/local/bin/julia

WORKDIR /app
COPY configs/docker/judge_docker.json ./judge.json

CMD ["./njudge", "judge"]