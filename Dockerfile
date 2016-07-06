FROM scratch
#COPY ui/build/ /build
#COPY ui/index.html /
COPY slowserver /
CMD ["/slowserver"]
