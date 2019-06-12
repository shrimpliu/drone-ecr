# DRONE-CI AWS ECR插件

- 使用

```yaml
kind: pipeline
name: default
steps:
  - name: publish
    image: droneplugin/ecr
    privileged: true
    settings:
      repo: {Your-Repo-Name}
      region: {Your-AWS-Region}
      access_key: {Your-AWS-AccessKey}
      secret_key: {Your-AWS-SecretKey}  
```

其他配置参照[drone-docker](http://plugins.drone.io/drone-plugins/drone-docker/)