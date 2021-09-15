FROM scratch
 COPY . .
 ENTRYPOINT ["/go-server"]

EXPOSE 80

CMD ["/go-server"]