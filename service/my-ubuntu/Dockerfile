FROM ubuntu:trusty
RUN apt-get update \
    && apt-get install -y curl \
    && rm -rf /var/lib/apt/lists/*
COPY loop.sh /
CMD ["/loop.sh"]
#CMD /loop.sh
