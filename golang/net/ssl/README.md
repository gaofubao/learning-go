
```shell
# 创建 CA 证书
cfssl print-defaults config > ca-config.json
cfssl print-defaults csr > ca-csr.json
cfssl gencert -initca ca-csr.json | cfssljson -bare ca/ca -

# 创建 Server 端证书
cfssl print-defaults csr > server.json
cfssl gencert -ca=ca/ca.pem -ca-key=ca/ca-key.pem -config=ca-config.json -profile=www server.json | cfssljson -bare server/server

# 创建 Client 端证书
cfssl print-defaults csr > client.json
cfssl gencert -ca=ca/ca.pem -ca-key=ca/ca-key.pem -config=ca-config.json -profile=client client.json | cfssljson -bare client/client

# 验证
cfssl certinfo -cert ca/ca.pem
openssl x509 -in ca/ca.pem -text -noout
```

单向认证

双向认证
