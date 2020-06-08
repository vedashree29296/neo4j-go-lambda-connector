FROM lambci/lambda-base:build

RUN yum makecache fast; yum clean all && yum -y update && yum -y upgrade; yum clean all && \
    yum install -y yum-plugin-ovl; yum clean all && yum -y groupinstall "Development Tools"; yum clean all

RUN yum -y install gcc gcc-c++ make pkgconfig wget
RUN openssl version -a
RUN wget https://github.com/neo4j-drivers/seabolt/releases/download/v1.7.4/seabolt-1.7.4-Linux-centos-7.tar.gz
RUN tar zxvf seabolt-1.7.4-Linux-centos-7.tar.gz --strip-components=1 -C /
RUN mkdir -p /opt/lib
RUN cp -r /usr/local/lib64/* /opt/lib