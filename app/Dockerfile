FROM debian:bullseye as builder

ARG NODE_VERSION=22.9.0



RUN apt-get update; apt install -y curl
RUN curl https://get.volta.sh | bash
ENV VOLTA_HOME /root/.volta
ENV PATH /root/.volta/bin:$PATH
RUN volta install node@${NODE_VERSION}

#######################################################################

RUN mkdir /app
WORKDIR /app

ENV NODE_ENV production

COPY . .

# https://stackoverflow.com/a/72323758
RUN chown -R root:root .

RUN npm ci && npm run build
# RUN npm run build
FROM debian:bullseye
EXPOSE 5000
EXPOSE 3000

LABEL fly_launch_runtime="nodejs"

COPY --from=builder /root/.volta /root/.volta
COPY --from=builder /app /app

WORKDIR /app
ENV NODE_ENV production
ENV PATH /root/.volta/bin:$PATH

CMD [ "npm", "run", "start" ]
