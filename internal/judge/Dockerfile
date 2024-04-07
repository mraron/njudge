ARG PROJECT_NAME

FROM ubuntu:22.04 as judge_deps
RUN apt-get update && apt-get install -y wget gcc g++ git build-essential libcap-dev
WORKDIR /app
RUN git clone --depth 1 --branch v1.10.1 https://github.com/ioi/isolate.git
WORKDIR /app/isolate
RUN make isolate

WORKDIR /
RUN mkdir languages

WORKDIR /languages
RUN wget https://julialang-s3.julialang.org/bin/linux/x64/1.8/julia-1.8.4-linux-x86_64.tar.gz && \
    tar zxvf julia-1.8.4-linux-x86_64.tar.gz && \
    rm julia-1.8.4-linux-x86_64.tar.gz

RUN wget https://nim-lang.org/download/nim-1.6.10-linux_x64.tar.xz && \
    tar xvf nim-1.6.10-linux_x64.tar.xz && \
    rm nim-1.6.10-linux_x64.tar.xz

FROM ${PROJECT_NAME}-base

COPY --from=judge_deps /app/isolate /app/isolate
WORKDIR /app/isolate
RUN make install

COPY --from=judge_deps /languages/* /languages/
RUN ln -s /languages/julia-1.8.4/bin/julia /usr/local/bin/julia && \
    ln -s /languages/nim-1.6.10/bin/nim /usr/bin/nim

WORKDIR /app
COPY configs/docker/judge_docker.json ./judge.json

CMD ["./njudge", "judge"]