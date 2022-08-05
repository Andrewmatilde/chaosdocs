## 模拟自定义故障

本文主要介绍如何使用 `Chaosd` 模拟自定义故障.
该功能通过使用用户自定义脚本模拟自定义故障，支持通过命令行模式或服务模式创建实验。
### 使用命令行模式模拟自定义故障

本节介绍如何在命令行模式中模拟自定义故障。

在模拟自定义故障前，可运行以下命令行查看模拟自定义故障的相关配置项：
```bash
chaosd attack user-defined help
```
输出结果如下所示：
```bash
user defined attack related commands 

 Usage:
  chaosd attack user-defined <subcommand> [flags]

Flags:
  -a, --attack-cmd string    the command to be executed when attack
  -r, --recover-cmd string   the command to be executed when recover

Global Flags:
      --log-level string   the log level of chaosd. The value can be 'debug', 'info', 'warn' and 'error'
      --uid string         the experiment ID

```
##### 模拟自定义故障相关配置说明

相关配置说明如下所示：

| 配置项 | 配置缩写 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | --- | ------------ | --- |
| attack-cmd |  `a`  | `string` | `""` | 攻击命令 | true |
| recover-cmd |  `r`  | `string` | `""` | 恢复命令 | true |
### 使用服务模式创建实验
要使用服务模式创建实验，请进行以下操作：

1. 以服务模式运行 chaosd。

   ```bash
   chaosd server --port 31767
   ```

2. 向 Chaosd 服务的路径 /api/attack/user-defined 发送 `POST` HTTP 请求。

   ```bash
   curl -X POST 172.16.112.130:31767/api/attack/user-defined -H "Content-Type:application/json" -d '{fault-configuration}'
   ```

在上述命令中，你需要按照故障类型在 `fault-configuration` 中进行配置。有关对应的配置参数，请参考下文中各个类型故障的相关参数说明和命令示例。

注意

在运行实验时，请注意保存实验的 UID 信息。当要结束 UID 对应的实验时，需要向 Chaosd 服务的路径 /api/attack/{uid} 发送 `DELETE` HTTP 请求。


#### 模拟自定义故障通用参数说明

相关配置说明如下所示：

| 配置项 | 类型 | 默认值 | 说明 | 必须参数 |
| --- | --- | --- | ------------ | --- |
| attack-cmd | `string` | `""` | 攻击命令 | true |
| recover-cmd | `string` | `""` | 恢复命令 | true |

