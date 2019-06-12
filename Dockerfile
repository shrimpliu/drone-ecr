FROM plugins/docker

ADD ecr /bin/
ENTRYPOINT ["/usr/local/bin/dockerd-entrypoint.sh", "/bin/ecr"]