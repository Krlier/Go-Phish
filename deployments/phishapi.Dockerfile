FROM golang:1.8
ADD app/ /app
WORKDIR /app
RUN update-ca-certificates
RUN git clone https://github.com/elceef/dnstwist.git
RUN go get gopkg.in/matryer/respond.v1
RUN go get github.com/gorilla/mux
RUN go get github.com/agnivade/levenshtein
RUN go run main.go levenshtein.go &
