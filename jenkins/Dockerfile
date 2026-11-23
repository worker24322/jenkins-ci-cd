# Jenkins kèm Go (để test) và Docker CLI (để build, deploy).
FROM jenkins/jenkins:lts-jdk21

USER root
COPY --from=golang:1.25 /usr/local/go /usr/local/go
COPY --from=docker:28-cli /usr/local/bin/docker /usr/local/bin/docker
COPY --from=docker:28-cli /usr/local/libexec/docker/cli-plugins /usr/local/libexec/docker/cli-plugins
ENV PATH="/usr/local/go/bin:${PATH}"

# Cho phép git đọc repo ở /repo dù nó thuộc về user khác (user của máy thật).
RUN git config --system --add safe.directory /repo \
 && git config --system --add safe.directory /repo/.git
