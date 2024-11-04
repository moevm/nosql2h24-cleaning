FROM denoland/deno:alpine-2.0.4 AS develop

COPY package.json deno.lock ./
RUN deno install
COPY . .

FROM develop as builder

RUN deno task build

FROM nginx:1.25.5 as deploy

COPY --from=builder /dist /usr/share/nginx/html
CMD [ "nginx", "-g", "daemon off;" ]