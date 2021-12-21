## protocol_mail

> protocol_mail processor 根据raw_mail解析详细邮件内容，并提取字段放入 event 中。

### 如何使用?

将此 processor 添加到 filebeat 后，你可以在 filebeat processors 配置段中增加以下配置:

``` yaml
processors:
  - protocol_mail:
      # 源字段，protocol_mail processor 从此字段读取到一个文件名，然后按照分隔符提取前缀
      # 此配置默认值为 "row_mail"
      source_field: "row_mail"
      # 目标字段，protocol_mail processor 提取前缀成功后将其写入到目标字段中
      # 此配置默认值为 "mail"
      target_field: "mail"
      # processor 标记位，protocol_mail processor 处理成功后会将此字段设置为 true
      # 通常该字段用于标识作用，方便后面的 logstash 判断 event 是否被某个 processor 处理过
      # 此配置默认值为 "processors.protocol_mail"
      processors_field: "processors.protocol_mail"
      # 当无法找到 source_field 指定的字段时，如果该配置为 true，则忽略错误，继续处理 event
      # 此配置默认值为 false
      ignore_missing: true
      # 当出现一些错误时(例如上面的 source_field 找不到或者 source_field 不是个字符串等)忽略
      # 错误继续处理 event，可以将 ignore_failure 视为 ignore_missing 的更大范畴兼容
      # 此配置默认值为 true
      ignore_failure: true
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
