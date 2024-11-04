FROM denoland/deno:alpine-2.0.4

COPY package.json deno.lock ./
RUN deno install

COPY . .
CMD [ "deno", "task", "dev" ]
