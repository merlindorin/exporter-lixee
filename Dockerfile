FROM scratch

COPY exporter-lixee /usr/bin/exporter-lixee

ENTRYPOINT [ "/usr/bin/exporter-lixee" ]

CMD ["serve"]

