FROM ubuntu:22.04 as isolate_builder
RUN apt-get update && apt-get install -y gcc g++ git build-essential libcap-dev
WORKDIR /app
RUN git clone https://github.com/ioi/isolate.git
WORKDIR /app/isolate
RUN make isolate

FROM njudge_base

COPY --from=isolate_builder /app/isolate /app/isolate
WORKDIR /app/isolate
RUN make install

WORKDIR /
RUN mkdir languages

WORKDIR /languages
RUN wget https://julialang-s3.julialang.org/bin/linux/x64/1.8/julia-1.8.4-linux-x86_64.tar.gz && \
    tar zxvf julia-1.8.4-linux-x86_64.tar.gz && \
    ln -s /languages/julia-1.8.4/bin/julia /usr/local/bin/julia && \
    rm julia-1.8.4-linux-x86_64.tar.gz

RUN wget https://nim-lang.org/download/nim-1.6.10-linux_x64.tar.xz && \
    tar xvf nim-1.6.10-linux_x64.tar.xz && \
    ln -s /languages/nim-1.6.10/bin/nim /usr/bin/nim && \
    rm nim-1.6.10-linux_x64.tar.xz

WORKDIR /app
COPY configs/docker/judge_docker.json ./judge.json

CMD ["./njudge", "judge"]