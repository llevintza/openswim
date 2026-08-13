# Dev-container image: Go toolchain for apps/api.
# Host machines should not need a local Go install — open this repo in a Dev Container.
FROM mcr.microsoft.com/devcontainers/go:1.22-bookworm

RUN apt-get update \
  && apt-get install -y --no-install-recommends curl ca-certificates \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace
