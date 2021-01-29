#!/usr/bin/env bash

ln -s /usr/local/apache2/ /etc/httpd
ln -s /usr/local/apache2/modules/ /etc/httpd/modules
cp /assets/4.3.3/dispatcher-apache2.4-4.3.3.so /etc/httpd/modules/
cp /assets/4.3.3/dispatcher-apache2.4-4.3.3.so /etc/httpd/modules/
cp /assets/4.3.3/conf/mime.types /etc
chmod 755 /etc/httpd/modules/dispatcher-apache2.4-4.3.3.so
ln -s /usr/local/apache2/logs /etc/httpd/logs
mkdir -p /etc/httpd/conf
mv /etc/httpd/httpd.conf /etc/httpd/conf
ln -s /etc/httpd/conf/httpd.conf /etc/httpd/httpd.conf
mkdir -p /usr/local/apache2/conf.modules.d
ln -s /usr/local/apache2/conf.modules.d /etc/httpd/conf.modules.d
cp /assets/4.3.3/conf/modules.d/* /usr/local/apache2/conf.modules.d
ln -s /usr/local/apache2/modules/dispatcher-apache2.4-4.3.3.so /usr/local/apache2/modules/mod_dispatcher.so
addgroup apache
useradd -g apache apache
mkdir -p /var/www/html \
  /var/www/author \
  /mnt/var/www/default \
  /var/www/author \
  /var/www/lc
