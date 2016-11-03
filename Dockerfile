FROM scratch

COPY app /
COPY sites.csv /

ENTRYPOINT ["/app"]
