FROM scratch

COPY ./tcp2stdout /
ENTRYPOINT [ "/tcp2stdout" ]
