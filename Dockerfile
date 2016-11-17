FROM golang:1.6

# Install beego and the bee dev tool
RUN go get github.com/dburbridge/uruptime

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["uruptime"]
