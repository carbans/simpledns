FROM scratch

COPY simpledns /

EXPOSE 53/udp
EXPOSE 53/tcp


ENTRYPOINT ["/simpledns"]


