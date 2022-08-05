## 模拟HTTP故障

本文主要介绍如何使用 `Chaosd` 模拟HTTP故障.
该功能通过使用Linux网络组件模拟HTTP故障，支持通过命令行模式或服务模式创建实验。
### 使用命令行模式创建实验

在创建HTTP故障实验前，可运行以下命令行查看 `Chaosd` 支持的HTTP故障类型
```bash
chaosd attack http help
```
输出结果如下所示：
```bash
HTTP attack related commands 

 Usage:
  chaosd attack http [command]

Available Commands:
  abort       abort selected HTTP connection
  config      attack with config file
  delay       delay selected HTTP Package
  request     request specific URL, only support GET

Global Flags:
      --log-level string   the log level of chaosd. The value can be 'debug', 'info', 'warn' and 'error'
      --uid string         the experiment ID

Use "chaosd attack http [command] --help" for more information about a command.

```

目前 Chaosd 支持模拟HTTP故障类型 `abort`  `config`  `delay`  `request` .

#### 使用命令行模式模拟HTTP中断

通过运行HTTP中断命令，查看模拟HTTP中断场景支持的配置。

##### 模拟HTTP中断命令

具体命令如下所示：

```bash
chaosd attack http abort help
```

输出结果如下所示：

```bash
abort selected HTTP connection 

 Usage:
  chaosd attack http abort [flags]

Flags:
  -c, --code string         Code is a rule to select target by http status code in response.
  -m, --method string       HTTP method
      --path string         Match path of Uri with wildcard matches.
      --port int32          The TCP port that the target service listens on.
  -p, --proxy-ports uints   composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports (default [])
  -t, --target string       HTTP target: Request or Response

Global Flags:
      --log-level string   the log level of chaosd. The value can be 'debug', 'info', 'warn' and 'error'
      --uid string         the experiment ID

```
##### 模拟HTTP中断相关配置说明

相关配置说明如下所示：

| 配置项 | 配置缩写 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | --- | ------------ | --- |
| code |  `c`  | `string` | `""` | 代码是通过http状态码响应选择目标的规则。 | false |
| method |  `m`  | `string` | `""` | HTTP 方法 | false |
| path |  无  | `string` | `""` | Uri 的匹配路径基于通配符匹配。 | false |
| port |  无  | `int32` | `0` | 目标服务侦听的 TCP 端口。 | false |
| proxy-ports |  `p`  | `uintSlice` | `[]` | 由 HTTP 连接的端口之一组成，我们只会使用 proxy_ports 内的端口来攻击 HTTP 连接 | true |
| target |  `t`  | `string` | `""` | HTTP 目标: Request or Response | true |
#### 使用命令行模式模拟HTTP故障配置

通过运行HTTP故障配置命令，查看模拟HTTP故障配置场景支持的配置。

##### 模拟HTTP故障配置命令

具体命令如下所示：

```bash
chaosd attack http config help
```

输出结果如下所示：

```bash
attack with config file 

 Usage:
  chaosd attack http config [flags]

Flags:
  -p, --file path string   Config file path.

Global Flags:
      --log-level string   the log level of chaosd. The value can be 'debug', 'info', 'warn' and 'error'
      --uid string         the experiment ID

```
##### 模拟HTTP故障配置相关配置说明

相关配置说明如下所示：

| 配置项 | 配置缩写 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | --- | ------------ | --- |
| file path |  `p`  | `string` | `""` | 配置文件路径. | true |
#### 使用命令行模式模拟HTTP时延

通过运行HTTP时延命令，查看模拟HTTP时延场景支持的配置。

##### 模拟HTTP时延命令

具体命令如下所示：

```bash
chaosd attack http delay help
```

输出结果如下所示：

```bash
delay selected HTTP Package 

 Usage:
  chaosd attack http delay [flags]

Flags:
  -c, --code string         Code is a rule to select target by http status code in response.
  -d, --delay time string   Delay represents the delay of the target request/response.
  -m, --method string       HTTP method
      --path string         Match path of Uri with wildcard matches.
      --port int32          The TCP port that the target service listens on.
  -p, --proxy-ports uints   composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports (default [])
  -t, --target string       HTTP target: Request or Response

Global Flags:
      --log-level string   the log level of chaosd. The value can be 'debug', 'info', 'warn' and 'error'
      --uid string         the experiment ID

```
##### 模拟HTTP时延相关配置说明

相关配置说明如下所示：

| 配置项 | 配置缩写 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | --- | ------------ | --- |
| code |  `c`  | `string` | `""` | 代码是通过http状态码响应选择目标的规则。 | false |
| delay time |  `d`  | `string` | `""` | 延迟表示目标请求/响应的延迟。 | true |
| method |  `m`  | `string` | `""` | HTTP 方法 | false |
| path |  无  | `string` | `""` | Uri 的匹配路径基于通配符匹配。 | false |
| port |  无  | `int32` | `0` | 目标服务侦听的 TCP 端口。 | false |
| proxy-ports |  `p`  | `uintSlice` | `[]` | 由 HTTP 连接的端口之一组成，我们只会使用 proxy_ports 内的端口来攻击 HTTP 连接 | true |
| target |  `t`  | `string` | `""` | HTTP 目标: Request or Response | true |
#### 使用命令行模式模拟HTTP请求

通过运行HTTP请求命令，查看模拟HTTP请求场景支持的配置。

##### 模拟HTTP请求命令

具体命令如下所示：

```bash
chaosd attack http request help
```

输出结果如下所示：

```bash
request specific URL, only support GET 

 Usage:
  chaosd attack http request [flags]

Flags:
  -c, --count int          Number of requests to send (default 1)
  -p, --enable-conn-pool   Enable connection pool
      --url string         Request to send

Global Flags:
      --log-level string   the log level of chaosd. The value can be 'debug', 'info', 'warn' and 'error'
      --uid string         the experiment ID

```
##### 模拟HTTP请求相关配置说明

相关配置说明如下所示：

| 配置项 | 配置缩写 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | --- | ------------ | --- |
| count |  `c`  | `int` | `1` | 请求数量 | false |
| enable-conn-pool |  `p`  | `bool` | `false` | 连接池是否开启 | false |
| url |  无  | `string` | `""` | 发送目标 | true |
### 使用服务模式创建实验
要使用服务模式创建实验，请进行以下操作：

1. 以服务模式运行 chaosd。

   ```bash
   chaosd server --port 31767
   ```

2. 向 Chaosd 服务的路径 /api/attack/http 发送 `POST` HTTP 请求。

   ```bash
   curl -X POST 172.16.112.130:31767/api/attack/http -H "Content-Type:application/json" -d '{fault-configuration}'
   ```

在上述命令中，你需要按照故障类型在 `fault-configuration` 中进行配置。有关对应的配置参数，请参考下文中各个类型故障的相关参数说明和命令示例。

注意

在运行实验时，请注意保存实验的 UID 信息。当要结束 UID 对应的实验时，需要向 Chaosd 服务的路径 /api/attack/{uid} 发送 `DELETE` HTTP 请求。



#### 使用服务模式模拟HTTP中断

本节介绍如何使用服务模式模拟HTTP中断。

##### 模拟HTTP中断相关参数说明

相关配置说明如下所示：

| 配置项 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | ------------ | --- |
| action | string | 无 | 实验的行为,设置为 "abort" |
| code | `string` | `""` | 代码是通过http状态码响应选择目标的规则。 | false |
| method | `string` | `""` | HTTP 方法 | false |
| path | `string` | `""` | Uri 的匹配路径基于通配符匹配。 | false |
| port | `int32` | `0` | 目标服务侦听的 TCP 端口。 | false |
| proxy-ports | `uintSlice` | `[]` | 由 HTTP 连接的端口之一组成，我们只会使用 proxy_ports 内的端口来攻击 HTTP 连接 | true |
| target | `string` | `""` | HTTP 目标: Request or Response | true |
#### 使用服务模式模拟HTTP故障配置

本节介绍如何使用服务模式模拟HTTP故障配置。

##### 模拟HTTP故障配置相关参数说明

相关配置说明如下所示：

| 配置项 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | ------------ | --- |
| action | string | 无 | 实验的行为,设置为 "config" |
| file path | `string` | `""` | 配置文件路径. | true |
#### 使用服务模式模拟HTTP时延

本节介绍如何使用服务模式模拟HTTP时延。

##### 模拟HTTP时延相关参数说明

相关配置说明如下所示：

| 配置项 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | ------------ | --- |
| action | string | 无 | 实验的行为,设置为 "delay" |
| code | `string` | `""` | 代码是通过http状态码响应选择目标的规则。 | false |
| delay time | `string` | `""` | 延迟表示目标请求/响应的延迟。 | true |
| method | `string` | `""` | HTTP 方法 | false |
| path | `string` | `""` | Uri 的匹配路径基于通配符匹配。 | false |
| port | `int32` | `0` | 目标服务侦听的 TCP 端口。 | false |
| proxy-ports | `uintSlice` | `[]` | 由 HTTP 连接的端口之一组成，我们只会使用 proxy_ports 内的端口来攻击 HTTP 连接 | true |
| target | `string` | `""` | HTTP 目标: Request or Response | true |
#### 使用服务模式模拟HTTP请求

本节介绍如何使用服务模式模拟HTTP请求。

##### 模拟HTTP请求相关参数说明

相关配置说明如下所示：

| 配置项 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | ------------ | --- |
| action | string | 无 | 实验的行为,设置为 "request" |
| count | `int` | `1` | 请求数量 | false |
| enable-conn-pool | `bool` | `false` | 连接池是否开启 | false |
| url | `string` | `""` | 发送目标 | true |
