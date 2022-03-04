#!/bin/bash
sed -i "s/RUN wget -nc /#RUN wget -nc /g" Dockerfile
sed -i "s/echo /#echo /g" Dockerfile
sed -i "s/tar -zxf \/tmp\/bitnami/#tar -zxf \/tmp\/bitnami/g" Dockerfile
sed -i "s/rm -rf \/tmp\/bitnami/#rm -rf \/tmp\/bitnami/g" Dockerfile
