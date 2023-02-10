# syntax=docker/dockerfile:1

# Prepare and run tests for all API integrations (ie. Form3)
FROM golang:1.19-alpine

RUN set -ex; \
    apk update; \
    apk add bash build-base

WORKDIR /test_space
COPY . /test_space

RUN chmod +x /test_space/integrations-entrypoint.sh
ENTRYPOINT ["/test_space/integrations-entrypoint.sh"]

CMD ["go", "test", "-v", "integrations", "integrations/pkg/Form3", "integrations/pkg/Form3/Organisation"]
