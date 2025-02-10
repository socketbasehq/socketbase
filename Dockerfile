FROM oven/bun:1 AS base
WORKDIR /app

FROM base AS install
COPY package.json .
RUN bun install --frozen-lockfile

FROM base AS builder
COPY --from=install /app/node_modules ./node_modules
COPY web ./web
COPY public ./public
COPY tsconfig*.json ./
COPY vite.config.ts ./
COPY package.json ./
COPY index.html ./

ENV NODE_ENV=production
RUN bun run build

FROM golang:1.23-alpine AS server_base
WORKDIR /app
COPY . .

RUN go build -o socketbase ./cmd/...

FROM server_base AS server
COPY --from=builder /app/dist .
COPY --from=server_base /app/socketbase .

CMD ["./socketbase", "start"]
