# See here for image contents: https://github.com/microsoft/vscode-dev-containers/blob/main/containers/go/.devcontainer/Dockerfile
ARG VARIANT="1"
FROM mcr.microsoft.com/devcontainers/go:1-${VARIANT}

# Install dependencies.
USER vscode

# Install tools
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.58.1
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d
RUN go install github.com/jstemmer/go-junit-report/v2@latest