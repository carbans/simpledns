FROM scratch

COPY simpledns /

EXPOSE 5353/udp
EXPOSE 5353/tcp


ENTRYPOINT ["/simpledns"]


