FROM plugins/docker

ADD ecr /bin/
ENTRYPOINT /bin/ecr