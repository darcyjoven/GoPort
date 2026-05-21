# 测试方法

config.yaml 新建一个测试用例

```yaml
sso:
    providers:
        - name: provider1
          enabled: true
          verifyURL: http://localhost:8081/api/verify
          verifyText: token
          successText: success
          secretKey: your-secret-key-1
          tokenExpireTime: 7d
          allowAutoRegister: true
          defaultAuthorityId: 888
```

后端服务启动后，运行，查看用户是否建立
`go run .\testsso.go`