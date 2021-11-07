FROM golang:1.17-buster@sha256:be7aa81b44dc85ddf4008bc5f3d5a5acfca8517620d0c4a393601c8e0495fb05

ENV MOD_VERSION 0.8
ENV MOD_CHECKSUM 65d5b5410d0b11dec3427d9363426add67bde60ddbd8277f5794c981af4b8ad7

RUN wget -O m.tgz https://github.com/cortesi/modd/releases/download/v${MOD_VERSION}/modd-${MOD_VERSION}-linux64.tgz && \
    echo "${MOD_CHECKSUM} m.tgz" | sha256sum -c - && \
    tar xvfz m.tgz && \
    mv modd-${MOD_VERSION}-linux64/modd /usr/local/bin && \
    chmod +x /usr/local/bin/modd && \
    rm m.tgz

# Default vendoring support
ENV GOPATH /go:/backend

VOLUME ["/backend"]
WORKDIR /backend

ENTRYPOINT ["/bin/bash"]