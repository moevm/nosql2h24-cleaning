FROM denoland/deno:alpine-2.0.4 AS develop

COPY package.json .
RUN deno install
COPY . .

FROM develop AS builder

RUN deno task build

FROM nginx:1.25.5 AS deploy

COPY --from=builder /dist /usr/share/nginx/html
COPY default.conf /etc/nginx/conf.d/default.conf

CMD [ "nginx", "-g", "daemon off;" ]