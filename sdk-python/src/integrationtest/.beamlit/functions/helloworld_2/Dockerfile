
FROM python:3.12-slim

ARG UV_VERSION="latest"
RUN apt update && apt install -y curl

# Install uv.
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/
RUN curl -fsSL https://raw.githubusercontent.com/beamlit/toolkit/main/install.sh | BINDIR=/bin sh
WORKDIR /beamlit

# Install the application dependencies.
COPY pyproject.toml /beamlit/pyproject.toml
COPY uv.lock /beamlit/uv.lock
RUN uv sync --no-cache

COPY README.m[d] /beamlit/README.md
COPY LICENS[E] /beamlit/LICENSE
COPY . /beamlit/src

ENV PATH="/beamlit/.venv/bin:$PATH"

ENTRYPOINT ["bl","serve","--port","80","--module","functions.helloworld.helloworld_2"]
