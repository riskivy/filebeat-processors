## protocol_mail

> protocol_mail processor 根据raw_mail解析详细邮件内容，并提取字段放入 event 中。

### 如何使用?

将此 processor 添加到 filebeat 后，你可以在 filebeat processors 配置段中增加以下配置:

``` yaml
processors:
  - protocol_mail:
      # 源字段，protocol_mail processor 从此字段读取到一个文件名，然后按照分隔符提取前缀
      # 此配置默认值为 "raw_mail"
      source_field: "raw_mail
      # processor 标记位，protocol_mail processor 处理成功后会将此字段设置为 true
      # 通常该字段用于标识作用，方便后面的 logstash 判断 event 是否被某个 processor 处理过
      # 此配置默认值为 "processors.protocol_mail"
      processors_field: "processors.protocol_mail"
```

### 如何调试?

你可以为 logstash 开启终端输出来实时观察日志处理情况:

``` sh
output {
  stdout {
    codec => rubydebug
  }
}
```
