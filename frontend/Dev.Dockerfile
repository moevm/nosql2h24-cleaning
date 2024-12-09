FROM denoland/deno:alpine-2.0.4

WORKDIR /vue_app

COPY package.json .
RUN deno install
EXPOSE 80

CMD ["deno", "task", "--node-modules-dir=auto", "dev", "--host=0.0.0.0", "--port=80"]