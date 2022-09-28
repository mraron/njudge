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

WORKDIR /app
COPY configs/judge_docker.json ./judge.json

CMD ["./njudge", "judge"]