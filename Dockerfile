FROM httpd:2.4

ADD docker/assets /assets
RUN bash /assets/setup.sh

