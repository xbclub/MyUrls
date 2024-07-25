
### 项目依赖

#### 1. 启动依赖

- Redis: 7.2.4+

#### 2. 编译依赖

- GO: 1.22.5+

### Docker-Compose

#### 1. 环境变量配置

| 变量名            | 默认值                   | 描述                    |
|----------------|-----------------------|-----------------------|
| WebSite_URL    | http://127.0.0.1:8888 | 本项目 域名地址              |
| REDIS_HOST     | 无                     | redis 地址              |
| REDIS_TYPE     | node                  | redis 类型              |
| REDIS_PASS     | 空                     | redis 密码 `没有不要添加这个变量` |
| REDIS_TLS      | false                 | redis 是否开启tls         |
| ShortKeyLength | 7                     | 短key生成长度              |
| ShortKeyTTL    | 604800                | 短连接有效时间  单位 s         |

#### 2. 启动命令

```bash
docker run -d \
  --name my-urls \
  -p 8888:8888 \
  -e WebSite_URL=http://127.0.0.1:8888 \
  -e REDIS_HOST=127.0.0.1:6379 \
  xbclub/my-urls:latest
```