FROM debian:11-slim
COPY hello-world-server /usr/local/bin
CMD [ "/usr/local/bin/hello-world-server" ]
EXPOSE 8000
