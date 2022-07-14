FROM njudge_base

RUN apt-get install -y git libcap-dev
RUN git clone https://github.com/ioi/isolate.git
WORKDIR /app/isolate
RUN make isolate
RUN make install
WORKDIR /app

COPY configs/judge_docker.json ./judge.json

CMD ["./njudge", "judge"]