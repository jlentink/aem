.PHONY: all build start shell test

CONTAINER_VERSION?=4.3.2
CONTAINER_NAME?=aem-dispatcher
TESTNODE_VERSION?=

all: clean build

clean:
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep ${CONTAINER_NAME}`

build:
	docker build . --tag ${CONTAINER_NAME}:${CONTAINER_VERSION}

shell:
	docker run -it --rm --entrypoint "/bin/bash" \
	-v /Volumes/codespace/projects/A/Azko/Platform/dispatcher/src/conf:/usr/local/apache2/conf/conf \
	-v /Volumes/codespace/projects/A/Azko/Platform/dispatcher/src/conf.d:/usr/local/apache2/conf/conf.d \
	-v /Volumes/codespace/projects/A/Azko/Platform/dispatcher/src/conf.dispatcher.d:/usr/local/apache2/conf/conf.dispatcher.d \
    -e DISP_ID=dispatcher1euwest1 \
    -e AUTHOR_IP=host.docker.internal \
    -e AUTHOR_PORT=4502 \
    -e AUTHOR_DEFAULT_HOSTNAME=host.docker.internal \
    -e AUTHOR_DOCROOT=/var/www/author \
    -e PUBLISH_IP=host.docker.internal \
    -e PUBLISH_PORT=4503 \
    -e PUBLISH_DEFAULT_HOSTNAME=host.docker.internal \
    -e PUBLISH_BETA_HOSTNAME= \
    -e PUBLISH_DOCROOT=/var/www/html \
    -e LIVECYCLE_IP=127.0.0.1 \
    -e LIVECYCLE_PORT=8080 \
    -e LIVECYCLE_DEFAULT_HOSTNAME=host.docker.internal \
    -e LIVECYCLE_DOCROOT=/var/www/lc \
    -e CRX_FILTER=deny \
    -p 8080:80 \
    --name dispatcher \
	${CONTAINER_NAME}:${CONTAINER_VERSION}

#shell:
#	docker run -it --rm --entrypoint "/bin/sh" \
#	-v /Volumes/codespace/projects/A/Azko/Platform/dispatcher/src/conf.d:/etc/httpd/conf.d \
#	-v /Volumes/codespace/projects/A/Azko/Platform/dispatcher/src/conf.dispatcher.d:/etc/httpd/conf.dispatcher.d \
#	${CONTAINER_NAME}
run:
	docker run --rm  \
	-v /Volumes/codespace/projeenvcts/A/Azko/Platform/dispatcher/src/conf.d:/etc/httpd/conf.d \
	-v /Volumes/codespace/projects/A/Azko/Platform/dispatcher/src/conf.dispatcher.d:/etc/httpd/conf.dispatcher.d \
	${CONTAINER_NAME}

tag:
	docker tag ${CONTAINER_NAME}:${CONTAINER_VERSION} jlentink/${CONTAINER_NAME}:${CONTAINER_VERSION}

push: tag
	docker push jlentink/${CONTAINER_NAME}:${CONTAINER_VERSION}
