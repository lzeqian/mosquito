FROM liaomin789/mosquito-base:1.0.0
RUN mkdir -p /application
RUN mkdir -p /application/conf
RUN mkdir -p /application/files
RUN mkdir -p /application/static
RUN mkdir -p /application/views
WORKDIR /application
ADD conf/ /application/conf/
ADD files/ /application/files/
ADD static/ /application/static/
ADD views/ /application/views/
ADD gpm /application/
CMD  /application/gpm
