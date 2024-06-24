ARG PROJECT_NAME

FROM node:22-bookworm as frontend-builder
COPY package.json ./
RUN npm install
COPY src/ ./src
COPY gulpfile.js ./
RUN npx gulp

FROM ${PROJECT_NAME}-base
COPY configs/docker/web.yaml ./web.yaml
COPY --from=frontend-builder static ./

CMD ["./njudge", "web"]