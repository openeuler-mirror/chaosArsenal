[TOC]



## 一、chaosArsenal工具使用说明

### 1.1 `OpenAPI server`端

#### 1.1.1 启动`server`端服务命令

```bash
# 指定ip和端口。
[root@localhost chaos-arsenal]# arsenal server start --host 10.103.177.165 --port 9095

# 指定ip。
[root@localhost chaos-arsenal]# arsenal server start --host 10.103.177.165

# 不指定ip和端口，localhost和默认端口9095起http服务器。
[root@localhost chaos-arsenal]# arsenal server start
```

#### 1.1.2 关闭`server`端服务命令

```bash
[root@localhost chaos-arsenal]# arsenal server stop --signal stop
```



### 1.2 故障查询命令

#### 1.2.1 `cli`

- 单一参数信息查询

  ```bash
  # Query all removed faults information
  [root@localhost chaos-arsenal]# arsenal query --status injected
  [
          {
                  "UUID": "3e0fea306512eadd",
                  "InteractiveMode": "http",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "lost",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "injected",
                  "InjectTime": "2023-12-11T14:28:12.538+08:00",
                  "UpdateTime": "2023-12-11T14:28:12.538+08:00"
          }
  ]
  
  # Query all faults information via Domain name
  [root@localhost chaos-arsenal]# arsenal query --domain file
  [
          {
                  "UUID": "6b94ca860d73b41c",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "removed",
                  "InjectTime": "2023-11-04T02:27:19.139+08:00",
                  "UpdateTime": "2023-11-04T02:30:22.283+08:00"
          },
          {
                  "UUID": "088f902f32885a03",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "Injected",
                  "InjectTime": "2023-11-04T02:30:29.290+08:00",
                  "UpdateTime": "2023-11-04T02:30:29.290+08:00"
          }
  ]
  
  # Query all faults information via fault type
  [root@localhost chaos-arsenal]# arsenal query --domain file --fault-type corruption
  [
          {
                  "UUID": "6b94ca860d73b41c",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "removed",
                  "InjectTime": "2023-11-04T02:27:19.139+08:00",
                  "UpdateTime": "2023-11-04T02:30:22.283+08:00"
          },
          {
                  "UUID": "088f902f32885a03",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "Injected",
                  "InjectTime": "2023-11-04T02:30:29.290+08:00",
                  "UpdateTime": "2023-11-04T02:30:29.290+08:00"
          }
  ]
  
  # Query all faults information via object
  [root@localhost chaos-arsenal]# arsenal query --object '/mnt/chaos-arsenal/kk.sh'
  [
          {
                  "UUID": "6b94ca860d73b41c",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "removed",
                  "InjectTime": "2023-11-04T02:27:19.139+08:00",
                  "UpdateTime": "2023-11-04T02:30:22.283+08:00"
          },
          {
                  "UUID": "088f902f32885a03",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "Injected",
                  "InjectTime": "2023-11-04T02:30:29.290+08:00",
                  "UpdateTime": "2023-11-04T02:30:29.290+08:00"
          }
  ]
  
  # Query fault information via fault UUID
  [root@localhost chaos-arsenal]# arsenal query --uuid 088f902f32885a03
  [
          {
                  "UUID": "088f902f32885a03",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "Injected",
                  "InjectTime": "2023-11-04T02:30:29.290+08:00",
                  "UpdateTime": "2023-11-04T02:30:29.290+08:00"
          }
  ]
  
  # Query fault information via injected time
  # 根据注入时间查找相关表项
  [root@localhost chaos-arsenal]# arsenal query --inject-time '2023-11-04T02:27'
  [
          {
                  "UUID": "6b94ca860d73b41c",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "removed",
                  "InjectTime": "2023-11-04T02:27:19.139+08:00",
                  "UpdateTime": "2023-11-04T02:30:22.283+08:00"
          }
  ]
  
  # Query fault information via databse update time
  # 根据表项更新时间查找相关表项，模糊匹配时间。
  [root@localhost chaos-arsenal]# arsenal query --update-time '2023-11-04T02:30:29.290'
  [
          {
                  "UUID": "088f902f32885a03",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "Injected",
                  "InjectTime": "2023-11-04T02:30:29.290+08:00",
                  "UpdateTime": "2023-11-04T02:30:29.290+08:00"
          }
  ]
  ```

- 多参数查询

  ```bash
  [root@localhost chaos-arsenal]# arsenal query --status removed --domain network
  [
          {
                  "UUID": "50965e6790749df7",
                  "InteractiveMode": "cli",
                  "Env": "hardware",
                  "Domain": "network",
                  "FaultType": "unavailable",
                  "Object": "ens19",
                  "Flags": "--interface ens19",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "removed",
                  "InjectTime": "2023-12-12T16:38:33.692+08:00",
                  "UpdateTime": "2023-12-12T16:39:57.074+08:00"
          },
  ]
  ```

  

#### 1.2.2 `http`

- 单一参数查询

  ```bash
  [root@localhost chaos-arsenal]# curl -X 'GET' 'http://10.103.177.165:9095/arsenal/v1/faults?status=injected'
  {"code":200,"infos":[{"domain":"file","FaultType":"lost","flags":"--path /mnt/chaos-arsenal/kk.sh","injectTime":"2023-12-11T14:28:12.538+08:00","proactiveCleanup":true,"status":"injected","updateTime":"2023-12-11T14:28:12.538+08:00","uuid":"3e0fea306512eadd"}]}
  ```

- 多参数查询

  ```bash
  [root@localhost chaos-arsenal]# curl -X 'GET' 'http://10.103.177.165:9095/arsenal/v1/faults?inject-time=2023-12-11T14:28&domain=file'
  {"code":200,"infos":[{"domain":"file","FaultType":"lost","flags":"--path /mnt/chaos-arsenal/kk.sh","injectTime":"2023-12-11T14:28:12.538+08:00","proactiveCleanup":true,"status":"injected","updateTime":"2023-12-11T14:28:12.538+08:00","uuid":"3e0fea306512eadd"}]}
  ```



### 1.3  重复注入、清理检查

- 重复注入

  ```bash
  # 已经对文件/mnt/chaos-arsenal/kk.sh做故障注入。
        {
                  "UUID": "088f902f32885a03",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "Injected",
                  "InjectTime": "2023-11-04T02:30:29.290+08:00",
                  "UpdateTime": "2023-11-04T02:30:29.290+08:00"
          }
  
  # 再次对/mnt/chaos-arsenal/kk.sh注入其他故障将会报错。
  [root@localhost chaos-arsenal]# arsenal inject os file lost --path /mnt/chaos-arsenal/kk.sh
  Error: handle ops failed([088f902f32885a03 file-lost --length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh] has been injected)
  ```

- 重复清理

  ```bash
  [root@localhost chaos-arsenal]# arsenal query all
  [
          {
                  "UUID": "6b94ca860d73b41c",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "removed",
                  "InjectTime": "2023-11-04T02:27:19.139+08:00",
                  "UpdateTime": "2023-11-04T02:30:22.283+08:00"
          },
          {
                  "UUID": "088f902f32885a03",
                  "InteractiveMode": "cli",
                  "Env": "os",
                  "Domain": "file",
                  "FaultType": "corruption",
                  "Object": "/mnt/chaos-arsenal/kk.sh",
                  "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
                  "Private": "",
                  "ProactiveCleanup": true,
                  "Status": "Injected",
                  "InjectTime": "2023-11-04T02:30:29.290+08:00",
                  "UpdateTime": "2023-11-04T02:30:29.290+08:00"
          }
  ]
  [root@localhost chaos-arsenal]# arsenal inject os file corruption --length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh
  Error: handle ops failed([088f902f32885a03 file-corruption --length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh] has been injected)
  ```

  

### 1.4 timeout延迟定时清理

`timeout`的参数格式为`xd:xh:xm:xs`的任意组合形式。

- cli

  ```bash
  # timeout延迟执行
  [root@localhost chaos-arsenal]# arsenal inject os file corruption --length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh --timeout 1m:10s
  
  # 数据库记录信息
  {
      "UUID": "a67dc3c95dbd6668",
      "InteractiveMode": "cli",
      "Env": "os",
      "Domain": "file",
      "FaultType": "corruption",
      "Object": "/mnt/chaos-arsenal/kk.sh",
      "Flags": "--length 2 --offset 1 --path /mnt/chaos-arsenal/kk.sh",
      "Private": "1007203,1m:10s",
      "ProactiveCleanup": true,
      "Status": "Injected",
      "InjectTime": "2023-11-06T01:50:35.279+08:00",
      "UpdateTime": "2023-11-06T01:50:35.279+08:00"
  }
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X POST -H "Content-Type: application/json" -d '{"path":"/mnt/chaos-arsenal/kk.sh"}' 'http://10.103.177.165:9095/arsenal/v1/faults/os/file-lost?timeout=1m:10s'
  {"code":200,"id":"f9c45ef7a1cc9264","message":"success"}
  
  # 数据库记录信息
  {
      "UUID": "f9c45ef7a1cc9264",
      "InteractiveMode": "http",
      "Env": "os",
      "Domain": "file",
      "FaultType": "lost",
      "Object": "/mnt/chaos-arsenal/kk.sh",
      "Flags": "--path /mnt/chaos-arsenal/kk.sh",
      "Private": "1m:10s",
      "ProactiveCleanup": true,
      "Status": "Injected",
      "InjectTime": "2023-11-06T01:56:37.253+08:00",
      "UpdateTime": "2023-11-06T01:56:37.253+08:00"
  }
  ```

  

### 1.5 cli命令自动补全

`arsenal completion bash > /etc/bash_completion.d/arsenal`需要重新启动终端才生效。

```bash
[root@localhost chaos-arsenal]# arsenal completion bash -h
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

        source <(arsenal completion bash)

To load completions for every new session, execute once:

#### Linux:

        arsenal completion bash > /etc/bash_completion.d/arsenal

#### macOS:

        arsenal completion bash > $(brew --prefix)/etc/bash_completion.d/arsenal

You will need to start a new shell for this setup to take effect.

Usage:
  arsenal completion bash

Flags:
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```



## 二、OS类故障注入能力

### 2.1 file-readonly（文件只读）

#### 原理

- 故障注入
  通过置文件`Immutable`属性实现文件只读，文件不可写，不可删除，类似于`shell`命令`chattr +i`。

- 故障清理

  清理文件文件`Immutable`属性。

注入是判断是不是已经被置相关属性。

<font color=red>注：文件正在写、正在使用等场景执行故障注入，可能会导致失败。</font>

#### 对象查找

```bash
[root@localhost chaos-arsenal]# realpath kk.sh
/mnt/chaos-arsenal/kk.sh
```

#### 功能测试

- cli

  ```bash
  arsenal inject os file readonly --path /mnt/chaos-arsenal/kk.sh
  arsenal remove os file readonly --path /mnt/chaos-arsenal/kk.sh
  ```
  
- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "file", "fault-type": "readonly", "params": {"path":"/mnt/chaos-arsenal/kk.sh"}}'
  {"code":200,"id":"9eae9d6d34fd25f5","message":"success"}
  
  # 检查文件属性
  [root@localhost chaos-arsenal]# lsattr /mnt/chaos-arsenal/kk.sh
  ----i--------------- /root/chaos-arsenal/kk.sh
  # 写文件，删除文件失败
  [root@localhost chaos-arsenal]# rm -rf kk.sh
  rm: cannot remove 'kk.sh': Operation not permitted
  [root@localhost chaos-arsenal]# echo "test" > kk.sh
  -bash: kk.sh: Operation not permitted
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/9eae9d6d34fd25f5
  {"code":200,"id":"","message":"success"}
  
  # 检查文件属性
  [root@localhost chaos-arsenal]# lsattr /mnt/chaos-arsenal/kk.sh
  --------------e----- /mnt/chaos-arsenal/kk.sh
  
  # 测试延迟执行
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "file", "fault-type": "readonly", "timeout": "1m:30s", "params": {"path":"/mnt/chaos-arsenal/kk.sh"}}'
  ```
  
  

### 2.2 file-unexecuted（文件不可执行）

#### 原理

- 故障注入
  通过修改文件`UGO`属性实现文件不可执行（类似于`chomd`），在修改文件属性之前先在对象文件所在目录备份文件原有属性，文件名为`%s_arsenal_backup_attr`，备份的内容示例：`755`。 -- 路径固化

- 故障清理

  读取文件权限备份文件`%s_arsenal_backup_attr`内容，并修改为文件原有权限。
  
  <font color=red>注：文件正在写、正在使用等场景执行故障注入，可能会导致失败。</font>

#### 对象查找

```bash
[root@localhost chaos-arsenal]# realpath ./kk.sh
/mnt/chaos-arsenal/kk.sh
```

#### 功能测试

- cli

  ```bash
  arsenal inject os file unexecuted --path /mnt/chaos-arsenal/kk.sh
  arsenal remove os file unexecuted --path /mnt/chaos-arsenal/kk.sh
  ```
  
- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "file", "fault-type": "unexecuted", "params": {"path":"/mnt/chaos-arsenal/kk.sh"}}'
  {"code":200,"id":"777cf7857986694a","message":"success"}
  
  # 检查文件属性，执行权限被移除。
  [root@localhost chaos-arsenal]# ll kk.sh
  -rw-r--r-- 1 root root 16 Dec  8 11:21 kk.sh
  # 文件路径存在属性备份文件。
  [root@localhost chaos-arsenal]# ll kk.sh-file-unexecuted-backup-attr
  -rw-r--r-- 1 root root 3 Dec 11 10:54 kk.sh-file-unexecuted-backup-attr
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/777cf7857986694a
  {"code":200,"id":"","message":"success"}
  
  # 检查文件属性，可执行权限被加回
  [root@localhost chaos-arsenal]# ll ./kk.sh
  -rwxr-xr-x 1 root root 16 Dec  8 11:21 ./kk.sh
  ```
  
  

### 2.3 file-lost（文件丢失）

#### 原理

- 故障注入
  通过文件重命名的方式模拟文件丢失，重命名文件名为`%s_arsenal_backup", filePath`。

- 故障清理

  将文件`%s%s", filePath, "_arsenal_backup`重命名为原文件名。
  
  <font color=red>注：文件正在写、正在使用等场景执行故障注入，可能会导致失败。</font>

#### 对象查找

```bash
[root@localhost chaos-arsenal]# realpath ./kk.sh
/mnt/chaos-arsenal/kk.sh
```

#### 功能测试

- cli

  ```bash
  arsenal inject os file lost --path /mnt/chaos-arsenal/kk.sh
  arsenal remove os file lost --path /mnt/chaos-arsenal/kk.sh
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "file", "fault-type": "lost", "params": {"path":"/mnt/chaos-arsenal/kk.sh"}}'
  {"code":200,"id":"1c825d3ea10023e1","message":"success"}
  
  # 原文件kk.sh被重命名为kk.sh_arsenal_backup
  [root@localhost chaos-arsenal]# ll |grep kk.sh
  -rwxr-xr-x 1 root root   16 Dec  8 11:21 kk.sh_arsenal_backup
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/1c825d3ea10023e1
  {"code":200,"id":"","message":"success"}
  
  # 文件被重命名为相应文件
  [root@localhost chaos-arsenal]# ll | grep kk.sh
  -rwxr-xr-x 1 root root        5 10月 25 17:35 kk.sh
  ```

  

### 2.4 file-corruption（文件损坏） 

建议使用在小文件场景下，大文件场景可能会报错，在`offset`较大场景也会导致注入时间过长。

#### 原理

- 故障注入
  向文件给定位置`offset`处开始随机写入`length`长度的数据，构造数据损坏，可以根据实际使用场景，自己指定文件的备份路径，备份文件名为`%s-file-corruption-backup`。

- 故障清理

  清理已经损坏的文件，从备份文件恢复。
  
  <font color=red>注：文件正在写、正在使用等场景执行故障注入，可能会导致失败。</font>

#### 对象查找

```bash
[root@localhost chaos-arsenal]# realpath kk.sh
/mnt/chaos-arsenal/kk.sh
```

#### 功能测试

- cli

  ```bash
  # offset为1，那么就是从文件的第二个字节开始，损坏长度为3个字节
  arsenal inject os file corruption --length 3 --offset 1 --path /mnt/chaos-arsenal/kk.sh
  arsenal remove os file corruption --length 3 --offset 1 --path /mnt/chaos-arsenal/kk.sh
  
  # offset为1，那么就是从文件的第二个字节开始，损坏长度为3个字节，备份路径为/home目录
  arsenal inject os file corruption --length 3 --offset 1 --path /mnt/chaos-arsenal/kk.sh --backup-path /home
  arsenal remove os file corruption --length 3 --offset 1 --path /mnt/chaos-arsenal/kk.sh
  ```

- http

  场景一：不带备份参数
  
  ```bash
  # 构造目标文件
  [root@localhost chaos-arsenal]# echo "arsenal-os test" > kk.sh
  [root@localhost chaos-arsenal]# cat kk.sh
  arsenal-os test
  
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "file", "fault-type": "corruption", "params": {"path":"/mnt/chaos-arsenal/kk.sh", "offset": "1", "length": "2"}}'
  {"code":200,"id":"d6de4e1541782675","message":"success"}
  
  # 原文件数据发生损坏
  [root@localhost chaos-arsenal]# cat kk.sh
  aAxYnal-os test
  # 原文件kk.sh被备份到/mnt/chaos-arsenal/文件夹下，并被命名为kk.sh-arsenal-os-file-corruption-backup
  [root@localhost chaos-arsenal]# ll kk.sh-file-corruption-backup
  -rw-r--r-- 1 root root 16 Dec  8 10:55 kk.sh-file-corruption-backup
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/d6de4e1541782675
  {"code":200,"id":"","message":"success"}
  
  # 文件内容恢复
  [root@localhost chaos-arsenal]# cat kk.sh
  arsenal-os test
  # 备份文件被删除
  [root@localhost chaos-arsenal]# ll kk.sh-file-corruption-backup
  ls: cannot access 'kk.sh-file-corruption-backup': No such file or directory
  ```
  
  
  
  场景二：带备份参数
  
  ```bash
  # 带备份路径的故障注入
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "file", "fault-type": "corruption", "params": {"path":"/mnt/chaos-arsenal/kk.sh", "offset": "1", "length": "2", "backup-path": "/tmp"}}'
  {"code":200,"id":"62fd3aa0a67b7415","message":"success"}
  
  # 原文件数据发生损坏
  [root@localhost chaos-arsenal]# cat kk.sh
  aP4enal-os test
  # 原文件kk.sh被备份到/tmp/文件夹下，并被命名为kk.sh-file-corruption-backup
  [root@localhost chaos-arsenal]# ll /tmp/kk.sh-file-corruption-backup
  -rw-r--r-- 1 root root 16 Dec  8 11:00 /tmp/kk.sh-file-corruption-backup
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/62fd3aa0a67b7415
  {"code":200,"id":"","message":"success"}
  
  # 文件内容恢复
  [root@localhost chaos-arsenal]# cat kk.sh
  arsenal-os test
  # 备份文件被删除
  [root@localhost chaos-arsenal]# ll /tmp/kk.sh-file-corruption-backup
  ls: cannot access 'kk.sh-file-corruption-backup': No such file or directory
  ```
  
  



### 2.5 filesystem-mountpoint-inode-exhaustion（挂载点inode耗尽）

#### 原理

- 故障注入
  多线程向指定挂载点批量创建文件消耗`inode`。

- 故障清理

  多线程删除指定挂载点下批量创建的文件。

- 特别说明
  该能力故障注入和清理在挂载点`inode`数量比较多时会耗费一定时间。

#### 对象查找

```bash
[root@localhost chaos-arsenal]# df -i
Filesystem                  Inodes IUsed   IFree IUse% Mounted on
devtmpfs                   1982863   463 1982400    1% /dev
tmpfs                      1987054     1 1987053    1% /dev/shm
tmpfs                      1987054   616 1986438    1% /run
tmpfs                      1987054    18 1987036    1% /sys/fs/cgroup
/dev/mapper/openeuler-root 4587520 79876 4507644    2% /
tmpfs                      1987054     5 1987049    1% /tmp
/dev/vda1                    65536   347   65189    1% /boot
/dev/mapper/openeuler-home 2695168    58 2695110    1% /home
tmpfs                      1987054     5 1987049    1% /run/user/0
```

#### 功能测试

- cli

  ```bash
  arsenal inject os filesystem mountpoint-inode-exhaustion --path /dev/shm
  arsenal remove os filesystem mountpoint-inode-exhaustion --path /dev/shm
  ```
  
- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "filesystem", "fault-type": "mountpoint-inode-exhaustion", "params": {"path":"/dev/shm"}}'
  {"code":200,"id":"17e02e8ce923d958","message":"success"}
  
  # 目标挂载点inode被耗尽
  [root@localhost chaos-arsenal]# df -i
  Filesystem                  Inodes   IUsed   IFree IUse% Mounted on
  devtmpfs                   1982863     463 1982400    1% /dev
  tmpfs                      1987054 1987054       0  100% /dev/shm
  tmpfs                      1987054     616 1986438    1% /run
  tmpfs                      1987054      18 1987036    1% /sys/fs/cgroup
  /dev/mapper/openeuler-root 4587520   79876 4507644    2% /
  tmpfs                      1987054       5 1987049    1% /tmp
  /dev/vda1                    65536     347   65189    1% /boot
  /dev/mapper/openeuler-home 2695168      58 2695110    1% /home
  tmpfs                      1987054       5 1987049    1% /run/user/0
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/17e02e8ce923d958
  {"code":200,"id":"","message":"success"}
  
  # 再次查看挂载点inode相关信息，inode可用量恢复正常，且挂载点下没有arsenal相关文件残留。
  [root@localhost chaos-arsenal]# df -i
  Filesystem                  Inodes IUsed   IFree IUse% Mounted on
  devtmpfs                   1982863   463 1982400    1% /dev
  tmpfs                      1987054     1 1987053    1% /dev/shm
  tmpfs                      1987054   616 1986438    1% /run
  tmpfs                      1987054    18 1987036    1% /sys/fs/cgroup
  /dev/mapper/openeuler-root 4587520 79876 4507644    2% /
  tmpfs                      1987054     5 1987049    1% /tmp
  /dev/vda1                    65536   347   65189    1% /boot
  /dev/mapper/openeuler-home 2695168    58 2695110    1% /home
  tmpfs                      1987054     5 1987049    1% /run/user/0
  [root@localhost chaos-arsenal]# ll /dev/shm/
  total 0
  ```



### 2.6 filesystem-mountpoint-space-full（挂载点磁盘空间满）

#### 原理

- 故障注入
  `dd`命令往挂载点下写一个大文件，文件名为`$mountpoint/filesystem-mountpoint-space-full-image`

- 故障清理

  将挂载点下大文件删除。

- 特别说明

  该能力故障注入和故障清理在挂载点空间比较大时会耗费一定时间。

#### 对象查找

```bash
[root@localhost chaos-arsenal]# df -i
Filesystem                  Inodes IUsed   IFree IUse% Mounted on
devtmpfs                   1982863   463 1982400    1% /dev
tmpfs                      1987054     1 1987053    1% /dev/shm
tmpfs                      1987054   616 1986438    1% /run
tmpfs                      1987054    18 1987036    1% /sys/fs/cgroup
/dev/mapper/openeuler-root 4587520 79876 4507644    2% /
tmpfs                      1987054     5 1987049    1% /tmp
/dev/vda1                    65536   347   65189    1% /boot
/dev/mapper/openeuler-home 2695168    58 2695110    1% /home
tmpfs                      1987054     5 1987049    1% /run/user/0
```

#### 功能测试

- cli

  ```bash
  arsenal inject os filesystem mountpoint-space-full --path /dev/shm
  arsenal remove os filesystem mountpoint-space-full --path /dev/shm
  ```
  
- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "filesystem", "fault-type": "mountpoint-space-full", "params": {"path":"/dev/shm"}}'
  {"code":200,"id":"42d3a16ba979c127","message":"success"}
  
  # 目标挂载点磁盘空间被耗尽
  [root@localhost chaos-arsenal]# df -h
  Filesystem                  Size  Used Avail Use% Mounted on
  devtmpfs                    7.6G     0  7.6G   0% /dev
  tmpfs                       7.6G  7.6G     0 100% /dev/shm
  tmpfs                       7.6G   25M  7.6G   1% /run
  tmpfs                       7.6G     0  7.6G   0% /sys/fs/cgroup
  /dev/mapper/openeuler-root   69G  3.4G   62G   6% /
  tmpfs                       7.6G     0  7.6G   0% /tmp
  /dev/vda1                   976M  126M  783M  14% /boot
  /dev/mapper/openeuler-home   41G  188M   38G   1% /home
  tmpfs                       1.6G     0  1.6G   0% /run/user/0
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/42d3a16ba979c127
  {"code":200,"id":"","message":"success"}
  
  # 再次查看挂载点inode相关信息，可用磁盘空间恢复正常，挂载点路径没有`arsenal`数据残留。
  [root@localhost chaos-arsenal]# df -h
  Filesystem                  Size  Used Avail Use% Mounted on
  devtmpfs                    7.6G     0  7.6G   0% /dev
  tmpfs                       7.6G     0  7.6G   0% /dev/shm
  tmpfs                       7.6G   25M  7.6G   1% /run
  tmpfs                       7.6G     0  7.6G   0% /sys/fs/cgroup
  /dev/mapper/openeuler-root   69G  3.4G   62G   6% /
  tmpfs                       7.6G     0  7.6G   0% /tmp
  /dev/vda1                   976M  126M  783M  14% /boot
  /dev/mapper/openeuler-home   41G  188M   38G   1% /home
  tmpfs                       1.6G     0  1.6G   0% /run/user/0
  [root@localhost chaos-arsenal]# ll /dev/shm/
  total 0
  ```



### 2.7 filesystem-io-overload（IO过载）

#### 原理

- 故障注入
  底层调用`stress-ng`跑`io`压力测试，可以根据需求，设定进程`nice`值。

- 故障清理

  将后台运行的`stress-ng`相关进程`kill`掉。

  

#### 对象查找

#### 功能测试

- cli

  ```bash
  arsenal inject os filesystem io-overload --iomix 4 --iomix-bytes 1G
  arsenal remove os filesystem io-overload --iomix 4 --iomix-bytes 1G
  ```

- http

  ```bash
  # 故障注入前bi/bo数值
  [root@localhost chaos-arsenal]# vmstat 1
  procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
   r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
   1  0      0 15629928  13588 102016    0    0    79   395  403  566  1  1 93  6  0
   0  0      0 15629928  13588 102016    0    0     0     0  495  295  0  0 100  0  0
   0  0      0 15629960  13588 102016    0    0     0     0  602  353  0  0 100  0  0
   0  0      0 15629960  13588 102016    0    0     0     0  517  297  0  0 100  0  0
   0  0      0 15629896  13588 102016    0    0     0     0  603  327  0  0 100  0  0
   0  0      0 15629864  13588 102016    0    0     0     0  418  256  0  0 100  0  0
  
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "filesystem", "fault-type": "io-overload", "params": {"iomix":"4", "iomix-bytes": "1G"}}'
  {"code":200,"id":"6e536259b5b56953","message":"success"}
  
  # 注入后的bi/bo数值
  [root@localhost chaos-arsenal]# vmstat 1
  procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
   r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
   1  4      0 15503628   8928 112800    0    0    75   379  389  544  1  1 93  5  0
   0  1      0 15497052  10504 117844    0    0   284  7584 6343 9539  1  2 85 12  0
   0  1      0 15489476  12156 123176    0    0   212  7988 6733 10168  1  2 85 12  0
   0  1      0 15482732  13732 128492    0    0   188  7408 6453 9975  1  2 85 12  0
   1  0      0 15476116  15344 133200    0    0   164  7756 6666 10279  1  2 85 12  0
   0  1      0 15468476  16848 139232    0    0   152  7576 6396 9831  1  2 86 12  0
   
   # 清理故障
  [root@localhost chaos-arsenal]#  curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/6e536259b5b56953
  {"code":200,"id":"","message":"success"}
  
  # 清理后注入后的bi/bo数值
  [root@localhost chaos-arsenal]# vmstat 1
  procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
   r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
   1  0      0 15609604   7808 104144    0    0    75   398  404  569  1  1 93  6  0
   0  0      0 15609604   7820 104132    0    0     0    88  582  334  0  0 100  0  0
   0  0      0 15609856   7820 104184    0    0     0     0  521  323  0  0 100  0  0
   0  0      0 15609856   7820 104184    0    0     0     0  484  295  0  0 100  0  0
   0  0      0 15609856   7820 104184    0    0     0     0  495  323  0  0 100  0  0
   0  0      0 15609888   7820 104184    0    0     0     0  604  348  0  0 100  0  0
  ```

### 2.8 cpu-overload（cpu过载）

#### 原理

`--cpu-load`为可选参数，如果不输入，默认为`100`。

`--nice`的取值范围为`-20~19`，如果设定值小于`-20`则值为`-20`，如果设定值大于`19`，则值为`19`。

- 故障注入
  底层调用`stress-ng`跑`cpu`压力测试，可以根据需求，设定进程`nice`值。

- 故障清理

  将后台运行的`stress-ng`相关进程`kill`掉。

  

#### 对象查找

```bash
[root@localhost chaos-arsenal]# cat /proc/cpuinfo | grep -w "processor"
processor       : 0
processor       : 1
processor       : 2
processor       : 3
processor       : 4
processor       : 5
processor       : 6
processor       : 7
```

#### 功能测试

- cli

  ```bash
  arsenal inject os cpu overload --cpu 1 --cpu-load 60 --taskset 1
  arsenal remove os cpu overload --cpu 1 --cpu-load 60 --taskset 1
  
  arsenal inject os cpu overload --cpu 2 --cpu-load 60 --taskset 1-2
  arsenal remove os cpu overload --cpu 2 --cpu-load 60 --taskset 1-2
  
  arsenal inject os cpu overload --cpu 1 --cpu-load 60 --taskset 1 --nice -20
  arsenal remove os cpu overload --cpu 1 --cpu-load 60 --taskset 1 --nice -20
  
  # 在timeout未到时清理对应故障
  arsenal inject os cpu overload --cpu 1 --cpu-load 60 --taskset 1 --timeout 10s
  arsenal remove os cpu overload --cpu 1 --cpu-load 60 --taskset 1 
  
  arsenal inject os cpu overload --cpu 3 --cpu-load 60 --taskset 1-2,7 -add
  arsenal remove os cpu overload --cpu 3 --cpu-load 60 --taskset 1-2,7
  ```
  
- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "cpu", "fault-type": "overload", "params": {"cpu":"2", "cpu-load":"60", "taskset":"2-3"}}'
  {"code":200,"id":"466b04b95b77a432","message":"success"}
  
  # top命令查看对应cpu使用率为60%
  %Cpu2  : 59.3 us,  0.9 sy,  0.0 ni, 38.9 id,  0.0 wa,  0.9 hi,  0.0 si,  0.0 st
  %Cpu3  : 60.7 us,  0.0 sy,  0.0 ni, 39.3 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
  
  # 清理故障，后台无stress-ng cpu相关进程。
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/466b04b95b77a432
  {"code":200,"id":"","message":"success"}
  [root@localhost chaos-arsenal]# ps aux | grep stress-ng | grep -v grep
  
  # 注入带nice值的过载故障。
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "cpu", "fault-type": "overload", "params": {"cpu":"2", "cpu-load":"60", "taskset":"2-3", "nice":"-20"}}'
  {"code":200,"id":"849ba6d5017c39c5","message":"success"}
  
  # top命令查看对应cpu使用率为60%,nice值为-20.
  %Cpu2  : 59.7 us,  0.0 sy,  0.0 ni, 38.9 id,  0.0 wa,  0.7 hi,  0.7 si,  0.0 st
  %Cpu3  : 59.3 us,  0.0 sy,  0.0 ni, 40.7 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
   32840 root       0 -20   20460   5104   3412 R  60.4   0.0   1:04.77 stress-ng-cpu
   32839 root       0 -20   20460   5104   3412 R  59.7   0.0   1:04.76 stress-ng-cpu
  ```



### 2.9 cpu-offline（cpu离线）

#### 原理

- 注入
  `echo 1 > /sys/devices/system/cpu/cpu%d/online`

- 清理

  `echo 0 > /sys/devices/system/cpu/cpu%d/online`

  

#### 对象查找

```bash
[root@localhost chaos-arsenal]# cat /proc/cpuinfo | grep -w "processor"
processor       : 0
processor       : 1
processor       : 2
processor       : 3
processor       : 4
processor       : 5
processor       : 6
processor       : 7
```

#### 功能测试

- cli

  ```bash
  # 注入1核下线
  arsenal inject os cpu offline --cpuid 1
  arsenal remove os cpu offline --cpuid 1
  
  # 注入1和4核下线
  arsenal inject os cpu offline --cpuid 1,4
  arsenal remove os cpu offline --cpuid 1,4
  
  # 注入1到4核下线
  arsenal inject os cpu offline --cpuid 1-4
  arsenal remove os cpu offline --cpuid 1-4
  
  # 注入1,2,4核下线
  arsenal inject os cpu offline --cpuid 1-2,4
  arsenal remove os cpu offline --cpuid 1-2,4
  
  # 注入cpu核1下线，设定timeout时间为1m:10s，在1m:10s未到时提前清理故障
  arsenal inject os cpu offline --cpuid 1 --timeout 1m:10s
  arsenal remove os cpu offline --cpuid 1()
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "cpu", "fault-type": "offline", "params": {"cpuid":"2"}}'
  {"code":200,"id":"4f88130f7adf8893","message":"success"}
  
  # top命令看到对应cpu核心被offline
  %Cpu0  :  0.0 us,  0.0 sy,  0.0 ni,100.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
  %Cpu1  :  0.0 us,  0.0 sy,  0.0 ni, 99.7 id,  0.0 wa,  0.0 hi,  0.0 si,  0.3 st
  %Cpu3  :  0.0 us,  0.0 sy,  0.0 ni,100.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/os/cpu-overload/0138262c68acd769
  {"code":200,"id":"","message":"success"}
  ```



### 2.10 memory-overload（内存过载）

#### 原理

- 故障注入
  底层调用`stress-ng`跑`cpu`压力测试，可以根据需求，设定进程`nice`值。

- 故障清理

  将后台运行的`stress-ng`相关进程`kill`掉。

  

#### 对象查找

```bash
[root@localhost chaos-arsenal]# free -h
              total        used        free      shared  buff/cache   available
Mem:           15Gi       179Mi        14Gi        24Mi       806Mi        14Gi
Swap:         7.9Gi          0B       7.9Gi
```

#### 功能测试

- cli

  ```bash
  # 根据百分比跑内存压力
  arsenal inject os memory overload --vm 10 --vm-bytes 50%
  arsenal remove os memory overload --vm 10 --vm-bytes 50%
  
  # 根据内存使用量跑内存压力
  arsenal inject os mem memload --vm 2 --vm-bytes 1G
  arsenal remove os mem memload --vm 2 --vm-bytes 1G
  ```
  
- http

  ```bash
  # 注入故障 -- 指定百分比
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "memory", "fault-type": "overload", "params": {"vm":"10", "vm-bytes":"50%"}}'
  {"code":200,"id":"799e54c8f640453b","message":"success"}
  
  # 查看内存使用量为系统可以量的50%
  [root@localhost chaos-arsenal]# free -h
                total        used        free      shared  buff/cache   available
  Mem:           15Gi       7.3Gi       7.1Gi        24Mi       806Mi       7.5Gi
  Swap:         7.9Gi          0B       7.9Gi
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/799e54c8f640453b
  {"code":200,"id":"","message":"success"}
  
  # 内存使用量恢复正常
  [root@localhost chaos-arsenal]# free -h
                total        used        free      shared  buff/cache   available
  Mem:           15Gi       181Mi        14Gi        24Mi       806Mi        14Gi
  Swap:         7.9Gi          0B       7.9Gi
  
  # 注入故障 -- 指定内存使用量
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "memory", "fault-type": "overload", "params": {"vm":"10", "vm-bytes":"10G"}}'
  {"code":200,"id":"90a8c3caafc42197","message":"success"}
  [root@localhost chaos-arsenal]# free -h
                total        used        free      shared  buff/cache   available
  Mem:           15Gi        10Gi       4.2Gi        24Mi       806Mi       4.5Gi
  Swap:         7.9Gi          0B       7.9Gi
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/90a8c3caafc42197
  {"code":200,"id":"","message":"success"}
  ```



### 2.11 process-choking（进程卡顿）

#### 原理

- 故障注入
  间隔`interval`交替向目标进程发送`SIGSTOP`和`SIGCOUNT`信号实现进程卡顿的效果。

- 故障清理

  将后台运行的相关进程`kill`掉。

  

#### 对象查找

```bash
ps aux
```

#### 功能测试

- cli

  ```bash
  arsenal inject os process choking --pid 33827 --interval 2
  arsenal remove os process choking --pid 33827 --interval 2
  ```

- http

  ```bash
  # 查找对象
  [root@localhost chaos-arsenal]# ps aux | grep while | grep -v grep
  root       33827  0.0  0.0 213900  3500 pts/2    S    19:10   0:00 /bin/bash ./while.sh
  
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "process", "fault-type": "choking", "params": {"pid":"33827", "interval":"2"}}'
  {"code":200,"id":"7bc2b3cfcf571f48","message":"success"}
  
  # 目标进程会切到T状态，停留的时间为`interval`
  [root@localhost chaos-arsenal]# ps aux | grep while | grep -v grep
  root       33827  0.0  0.0 213900  3500 pts/2    T    19:10   0:00 /bin/bash ./while.sh
  [root@localhost chaos-arsenal]# ps aux | grep while | grep -v grep
  root       33827  0.0  0.0 213900  3500 pts/2    S    19:10   0:00 /bin/bash ./while.sh
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/7bc2b3cfcf571f48
  {"code":200,"id":"","message":"success"}
  
  # 目标进程正常运行，不处于T状态
  [root@localhost chaos-arsenal]# ps aux | grep while | grep -v grep
  root       33827  0.0  0.0 213900  3500 pts/2    S    19:10   0:00 /bin/bash ./while.sh
  ```



### 2.12 process-exit-abnormally（进程异常退出）

#### 原理

- 故障注入
  通过向给定进程发送`SIGKILL`信号让进程异常退出。

- 故障清理

  NA

  

#### 对象查找

```bash
ps aux
```

#### 功能测试

- cli

  ```bash
  arsenal inject os process exit-abnormal --pid 33827
  ```

- http

  ```bash
  # 查找对象
  [root@localhost chaos-arsenal]# ps aux | grep while | grep -v grep
  root       33827  0.0  0.0 213900  3500 pts/2    S    19:10   0:00 /bin/bash ./while.sh
  
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "process", "fault-type": "exit-abnormal", "params": {"pid":"33827"}}'
  {"code":200,"id":"000dbf74b635d8b1","message":"success"}
  
  # 找不到目标进程
  [root@localhost chaos-arsenal]# ps aux | grep while | grep -v grep
  ```



### 2.13 process-hang（进程hang住）

#### 原理

- 故障注入
  向给定进程发送`SIGSTOP`信号进入`T`状态。

- 故障清理

  向给定进程发送`SIGCOUNT`信号让进程恢复正常运行。

  

#### 对象查找

```bash
ps aux
```

#### 功能测试

- cli

  ```bash
  arsenal inject os process hang --pid 35449
  arsenal remove os process hang --pid 35449
  ```

- http

  ```bash
  # 查找对象
  [root@localhost chaos-arsenal]# ps aux | grep while | grep -v grep
  root       35449  0.0  0.0 213900  3436 pts/2    S+   19:36   0:00 /bin/bash ./while.sh
  
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "process", "fault-type": "hang", "params": {"pid":"35449"}}'
  {"code":200,"id":"90a347aca246f456","message":"success"}
  
  # 进程一直处于T状态
  [root@localhost chaos-arsenal]# ps aux | grep while | grep -v grep
  root       35449  0.0  0.0 213900  3440 pts/2    T    19:36   0:00 /bin/bash ./while.sh
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/90a347aca246f456
  {"code":200,"id":"","message":"success"}
  
  # 进程处于非T状态
  [root@localhost chaos-arsenal]# ps aux | grep while | grep -v grep
  root       35449  0.0  0.0 213900  3440 pts/2    S    19:36   0:00 /bin/bash ./while.sh
  ```



### 2.14 system-panic（系统panic）

#### 原理

- 故障注入
  `echo c > /proc/sysrq-trigger`模拟系统异常宕机。

- 故障清理

  NA

  

#### 对象查找

#### 功能测试

- cli

  ```bash
  arsenal inject os system panic
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "system", "fault-type": "panic"}'
  
  # 系统触发宕机
  ```



### 2.15 system-reboot-abnormal（系统异常重启）

#### 原理

- 故障注入
  `echo b > /proc/sysrq-trigger`模拟系统异常重启。

- 故障清理

  NA

  

#### 对象查找

#### 功能测试

- cli

  ```bash
  arsenal inject os system reboot-abnormal
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "system", "fault-type": "reboot-abnormal"}'
  
  # 系统触发重启
  ```



### 2.16 system-oom（系统oom killer）

#### 原理

- 故障注入
  `echo f > /proc/sysrq-trigger`触发内核oom killer。

- 故障清理

  NA

  <font color=red>注：`oom`可能会将`arsenal`服务端`kill`掉。</font>

#### 对象查找

#### 功能测试

- cli

  ```bash
  arsenal inject os system oom
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "system", "fault-type": "oom"}'
  {"code":200,"id":"e4656fb2d4f8de55","message":"success"}
  
  # dmesg查看内核日志
  [  323.403616] Out of memory: Kill process 1439 (arsenal) score 1 or sacrifice child
  [  323.405043] Killed process 1439 (arsenal) total-vm:1602844kB, anon-rss:18104kB, file-rss:11384kB, shmem-rss:0kB
  [  323.407260] oom_reaper: reaped process 1439 (arsenal), now anon-rss:0kB, file-rss:0kB, shmem-rss:0kB
  ```





### 2.17 system-file-systems-readonly（系统文件系统只读）

#### 原理

- 故障注入
  `echo u > /proc/sysrq-trigger`模拟文件系统只读。

- 故障清理

  NA

  

#### 对象查找

#### 功能测试

- cli

  ```bash
  arsenal inject os system file-systems-readonly
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "system", "fault-type": "file-systems-readonly"}'
  {"code":500,"message":"unable to update database after retrying 3 times, error(open data base failed: cleanup database faults table failed: database connection failed: attempt to write a readonly database)"}
  
  # 系统/挂载点只读
  [root@localhost /]# echo "arsenal test" > /test.txt
  -bash: /test.txt: Read-only file system
  ```



### 2.18 system-service-stop（系统服务异常停止）

#### 原理

- 故障注入
  通过`service`或者`systemctl`将特定服务暂停。

- 故障清理

  NA

  

#### 对象查找



#### 功能测试

- cli

  ```bash
  arsenal inject os system service-stop --name sshd
  arsenal remove os system service-stop --name sshd
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "system", "fault-type": "service-stop", "params": {"name": "sshd"}}'
  {"code":200,"id":"f9f41f03a5d757fb","message":"success"}
  
  # sshd服务被禁用
  [root@localhost chaos-arsenal]# service sshd status                                                                                                                       Redirecting to /bin/systemctl status sshd.service
  ● sshd.service - OpenSSH server daemon
     Loaded: loaded (/usr/lib/systemd/system/sshd.service; enabled; vendor preset: enabled)
     Active: inactive (dead) since Mon 2023-12-11 20:31:20 CST; 4s ago
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/f9f41f03a5d757fb
  {"code":200,"id":"","message":"success"}
  
  # sshd服务恢复正常
  [root@localhost chaos-arsenal]# service sshd status
  Redirecting to /bin/systemctl status sshd.service
  ● sshd.service - OpenSSH server daemon
     Loaded: loaded (/usr/lib/systemd/system/sshd.service; enabled; vendor preset: enabled)
     Active: active (running) since Mon 2023-12-11 20:33:04 CST; 39s ago
  ```



### 2.19 system-service-restart（系统服务异常重启）

#### 原理

- 故障注入
  通过`service`或者`systemctl`将特定服务重启。

- 故障清理

  NA

  

#### 对象查找

#### 功能测试

- cli

  ```bash
  arsenal inject os system service-restart --name sshd
  ```
  
- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "system", "fault-type": "service-restart", "params": {"name": "sshd"}}'
  {"code":200,"id":"f9f41f03a5d757fb","message":"success"}
  ```



### 2.20 system-time-jump（系统时间跳变）

#### 原理

- 故障注入
  通过`date`命令设定指定的系统时间。

- 故障清理

  通过`hwclock`同步系统时间。

  

#### 对象查找

#### 功能测试

- cli

  ```bash
  # 向后跳变
  arsenal inject os system time-jump --direction backwards --interval 1h
  arsenal remove os system time-jump --direction backwards --interval 1h
  
  # 向前跳变
  arsenal inject os system time-jump --direction forwards --interval 1h
  arsenal remove os system time-jump --direction forwards --interval 1h
  
  # 设定复杂的跳变时间
  arsenal inject os system time-jump --direction forwards --interval 1h:1m:1s
  arsenal remove os system time-jump --direction forwards --interval 1h:1m:1s
  ```

- http

  ```bash
  # 故障注入前时间
  [root@localhost chaos-arsenal]# date
  Mon Dec 11 20:48:30 CST 2023
  
  # 故障注入
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "os", "domain": "system", "fault-type": "time-jump", "params": {"direction": "backwards", "interval": "1h:5m"}}'
  {"code":200,"id":"d612f88ef1798205","message":"success"}
  
  # 故障注入后时间
  [root@localhost chaos-arsenal]# date
  Mon Dec 11 19:43:36 CST 2023
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/d612f88ef1798205
  d{"code":200,"id":"","message":"success"}
  
  # 系统时间恢复正常
  [root@localhost chaos-arsenal]# date
  Mon Dec 11 20:49:45 CST 2023
  ```



## 三、硬件类故障注入能力

### 3.1 disk-offline（磁盘异常离线）

#### 原理

通过磁盘`sysfs`下状态控制节点来设定磁盘的状态。

- 故障注入
  `echo offline > /sys/block/$dev/device/state`

- 故障清理

  `echo running > /sys/block/$dev/device/state`

  

#### 对象查找

```bash
[root@localhost ~]# lsblk
NAME               MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda                  8:0    0  120G  0 disk
├─sda1               8:1    0    1G  0 part /boot
└─sda2               8:2    0  119G  0 part
  ├─openeuler-root 253:0    0   70G  0 lvm  /
  ├─openeuler-swap 253:1    0  7.9G  0 lvm  [SWAP]
  └─openeuler-home 253:2    0 41.1G  0 lvm  /home
sdb                  8:16   0    5G  0 disk
├─sdb1               8:17   0    1G  0 part
└─sdb2               8:18   0    1G  0 part
```



#### 功能测试

- cli

  ```bash
  arsenal inject hardware disk offline --device sdb
  arsenal remove hardware disk offline --device sdb
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "disk", "fault-type": "offline", "params": {"device": "sdb"}}'
  {"code":200,"id":"ec0b26b38682b9de","message":"success"}
  
  # 从硬盘中读取数据失败
  [root@localhost chaos-arsenal]# dd if=/dev/sdb of=./text.img bs=1M count=10 iflag=direct
  dd: failed to open '/dev/sdb': No such device or address
  
  # 恢复故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/ec0b26b38682b9de
  {"code":200,"id":"","message":"success"}
  
  # 磁盘状态恢复支撑，可以正常从磁盘中读取数据
  [root@localhost chaos-arsenal]# dd if=/dev/sdb of=./text.img bs=1M count=10 iflag=direct
  10+0 records in
  10+0 records out
  10485760 bytes (10 MB, 10 MiB) copied, 0.105663 s, 99.2 MB/s
  ```



### 3.2 disk-blocked（磁盘卡死）

#### 原理

通过磁盘`sysfs`下状态控制节点来设定磁盘的状态。

- 故障注入
  `echo blocked > /sys/block/$dev/device/state`

- 故障清理

  `echo running > /sys/block/$dev/device/state`

  

#### 对象查找

#### 功能测试

- cli

  ```bash
  arsenal inject hardware disk blocked --device sdb
  arsenal remove hardware disk blocked --device sdb
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "disk", "fault-type": "blocked", "params": {"device": "sdb"}}'
  {"code":200,"id":"687673553f08a80b","message":"success"}
  
  # dd命令从磁盘中读取数据
  [root@localhost chaos-arsenal]# dd if=/dev/sdb of=./text.img bs=1M count=10 iflag=direct
  
  # dd进程等待IO进入D状态
  [root@localhost ~]# ps aux | grep "dd if" | grep -v grep
  root        1703  0.0  0.0 213448  2764 pts/0    D+   14:54   0:00 dd if=/dev/sdb of=./text.img bs=1M count=10 iflag=direct
  
  # 恢复故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/687673553f08a80b
  {"code":200,"id":"","message":"success"}
  
  # dd进程成功从磁盘中读取到数据
  [root@localhost chaos-arsenal]# dd if=/dev/sdb of=./text.img bs=1M count=10 iflag=direct
  10+0 records in
  10+0 records out
  10485760 bytes (10 MB, 10 MiB) copied, 388.393 s, 27.0 kB/s
  ```



### 3.3 pcie-offline（pcie设备异常离线）

#### 原理

- 故障注入
  通过`pcie sysfs(echo 1 > /sys/bus/pci/devices/$bdf/remove)`节点移除对应的`pcie`设备，因为要通过记录`pcie`设备的`root bus`，来判断注入是不是合法的，一个`root bus`只能存在一个`pcie`设备`offline`故障，`root bus`备份文件路径为`../logs/pcie-$root_bus-$bdf`。

- 故障清理

  通过`pcie sysfs (/sys/devices/pci$root_bus/pci_bus/$parent_bus/rescan)`下`parent bus rescan`节点，重新扫描的方式将被移除的设备加回。

#### 对象查找

```bash
[root@localhost chaos-arsenal]# lspci -D
0000:00:00.0 Host bridge: Intel Corporation 440FX - 82441FX PMC [Natoma] (rev 02)
0000:00:01.0 ISA bridge: Intel Corporation 82371SB PIIX3 ISA [Natoma/Triton II]
0000:00:01.1 IDE interface: Intel Corporation 82371SB PIIX3 IDE [Natoma/Triton II]
0000:00:01.2 USB controller: Intel Corporation 82371SB PIIX3 USB [Natoma/Triton II] (rev 01)
0000:00:01.3 Bridge: Intel Corporation 82371AB/EB/MB PIIX4 ACPI (rev 03)
0000:00:02.0 VGA compatible controller: Cirrus Logic GD 5446
0000:00:03.0 Communication controller: Virtio: Virtio console
0000:00:0a.0 SCSI storage controller: Virtio: Virtio block device
0000:00:0b.0 SCSI storage controller: Virtio: Virtio block device
0000:00:12.0 Ethernet controller: Intel Corporation 82540EM Gigabit Ethernet Controller (rev 03)
0000:00:18.0 USB controller: Intel Corporation 82801I (ICH9 Family) USB UHCI Controller #1 (rev 03)
0000:00:18.1 USB controller: Intel Corporation 82801I (ICH9 Family) USB UHCI Controller #2 (rev 03)
0000:00:18.2 USB controller: Intel Corporation 82801I (ICH9 Family) USB UHCI Controller #3 (rev 03)
0000:00:18.7 USB controller: Intel Corporation 82801I (ICH9 Family) USB2 EHCI Controller #1 (rev 03)
0000:00:19.0 USB controller: NEC Corporation uPD720200 USB 3.0 Host Controller (rev 03)
```

#### 功能测试

- cli

  ```bash
  arsenal inject hardware pcie offline --bdf 0000:00:02.0
  arsenal remove hardware pcie offline --bdf 0000:00:02.0
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "pcie", "fault-type": "offline", "params": {"bdf": "0000:00:02.0"}}'
  {"code":200,"id":"4c484dd697d54153","message":"success"}
  
  # 找不到对应的pcie设备
  [root@localhost chaos-arsenal]# lspci -D | grep "0000:00:02.0"
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/4c484dd697d54153
  {"code":200,"id":"","message":"success"}
  
  # 对应的pcie设备被加回
  [root@localhost chaos-arsenal]# lspci -D | grep "0000:00:02.0"
  0000:00:02.0 VGA compatible controller: Cirrus Logic GD 5446
  ```

  

### 3.4 pcie-reset-abnormal（pcie设备异常重置） - - 缺少查验方法

#### 原理

- 故障注入
  通过`pcie sysfs (eccho 1 > /sys/bus/pci/devices/$bdf/reset)`节点对`pcie`设备做`reset`操作。

- 故障清理

  自动恢复，不需要手动清理。

#### 对象查找

```bash
[root@localhost chaos-arsenal]# find /sys/devices/ -name "reset" | grep pci | grep -v grep
/sys/devices/pci0000:17/0000:17:00.0/0000:18:00.0/0000:19:03.0/0000:1a:00.1/reset
/sys/devices/pci0000:17/0000:17:00.0/0000:18:00.0/0000:19:03.0/reset
```

#### 功能测试

- cli

  ```bash
  # 注入故障
  arsenal inject hardware pcie reset-abnormal --bdf 0000:1a:00.1
  ```
  
- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "pcie", "fault-type": "reset-abnormal", "params": {"bdf": "0000:00:02.0"}}'
  {"code":200,"id":"50965e6790749df7","message":"success"}
  ```
  
  

### 3.5 network-corrupt（网络错包）

#### 原理

- 故障注入

  ```bash
  tc qdisc add dev $interface root netem corrupt $percent
  ```

- 故障清理

  ```bash
  tc qdisc del dev $interface root netem corrupt $percent
  ```


#### 对象查找

```bash
ifconfig
ip addr show
```

#### 功能测试

- cli

  ```bash
  arsenal inject hardware network corrupt --interface ens19 --percent 50%
  arsenal remove hardware network corrupt --interface ens19 --percent 50%
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "corrupt", "params": {"interface":"ens19", "percent": "50%"}}'
  {"code":200,"id":"aa3e396349d7e0d6","message":"success"}
  
  # 网络错包导致packet loss
  root@:chaos-arsenal# ping 10.103.177.165
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=0.290 ms
  ......
  64 bytes from 10.103.177.165: icmp_seq=25 ttl=64 time=0.221 ms
  ^C
  --- 10.103.177.165 ping statistics ---
  25 packets transmitted, 12 received, 52% packet loss, time 24537ms
  rtt min/avg/max/mdev = 0.221/0.399/1.995/0.481 ms
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/aa3e396349d7e0d6
  {"code":200,"id":"","message":"success"}
  
  # ping目标注入不再有packet loss
  root@:chaos-arsenal# ping 10.103.177.165
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=0.347 ms
  ......
  64 bytes from 10.103.177.165: icmp_seq=9 ttl=64 time=0.252 ms
  ^C
  --- 10.103.177.165 ping statistics ---
  9 packets transmitted, 9 received, 0% packet loss, time 8195ms
  rtt min/avg/max/mdev = 0.208/0.245/0.347/0.039 ms
  ```

  

### 3.6 network-loss（网络丢包）

#### 原理

- 故障注入

  ```bash
  tc qdisc add dev $interface root netem loss $percent
  ```

- 故障清理

  ```bash
  tc qdisc del dev $interface root netem loss $percent
  ```


#### 对象查找

```bash
ifconfig
ip addr show
```

#### 功能测试

- cli

  ```bash
  arsenal inject hardware network loss --interface ens19 --percent 50%
  arsenal remove hardware network loss --interface ens19 --percent 50%
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "loss", "params": {"interface":"ens19", "percent": "50%"}}'
  {"code":200,"id":"9010a6ea19e94dc5","message":"success"}
  
  # 网络错包导致丢包
  root@:chaos-arsenal# ping 10.103.177.165
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=0.290 ms
  ......
  64 bytes from 10.103.177.165: icmp_seq=25 ttl=64 time=0.221 ms
  ^C
  --- 10.103.177.165 ping statistics ---
  25 packets transmitted, 12 received, 52% packet loss, time 24537ms
  rtt min/avg/max/mdev = 0.221/0.399/1.995/0.481 ms
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/9010a6ea19e94dc5
  {"code":200,"id":"","message":"success"}
  
  # ping目标注入不再有packet loss
  root@:chaos-arsenal# ping 10.103.177.165
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=0.347 ms
  ......
  64 bytes from 10.103.177.165: icmp_seq=9 ttl=64 time=0.252 ms
  ^C
  --- 10.103.177.165 ping statistics ---
  9 packets transmitted, 9 received, 0% packet loss, time 8195ms
  rtt min/avg/max/mdev = 0.208/0.245/0.347/0.039 ms
  ```



### 3.7 network-duplicate（网络重复包）

#### 原理

- 故障注入

  ```bash
  tc qdisc add dev $interface root netem duplicate $percent
  ```

- 故障清理

  ```bash
  tc qdisc del dev $interface root netem duplicate $percent
  ```

  

#### 对象查找

```bash
ifconfig
ip addr show
```

#### 功能测试

- cli

  ```bash
  arsenal inject hardware network duplicate --interface ens19 --percent 50%
  arsenal remove hardware network duplicate --interface ens19 --percent 50%
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "duplicate", "params": {"interface":"ens19", "percent": "50%"}}'
  {"code":200,"id":"bc68c3aeb8bbac32","message":"success"}
  
  # ping命令可以看到重复包
  root@:chaos-arsenal# ping 10.103.177.165
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=0.296 ms
  64 bytes from 10.103.177.165: icmp_seq=2 ttl=64 time=0.300 ms
  64 bytes from 10.103.177.165: icmp_seq=2 ttl=64 time=0.300 ms (DUP!)
  64 bytes from 10.103.177.165: icmp_seq=3 ttl=64 time=0.344 ms
  64 bytes from 10.103.177.165: icmp_seq=3 ttl=64 time=0.344 ms (DUP!)
  64 bytes from 10.103.177.165: icmp_seq=4 ttl=64 time=0.261 ms
  64 bytes from 10.103.177.165: icmp_seq=4 ttl=64 time=0.261 ms (DUP!)
  64 bytes from 10.103.177.165: icmp_seq=5 ttl=64 time=0.280 ms
  64 bytes from 10.103.177.165: icmp_seq=6 ttl=64 time=0.260 ms
  64 bytes from 10.103.177.165: icmp_seq=6 ttl=64 time=0.260 ms (DUP!)
  64 bytes from 10.103.177.165: icmp_seq=7 ttl=64 time=0.241 ms
  ^C
  --- 10.103.177.165 ping statistics ---
  7 packets transmitted, 7 received, +4 duplicates, 0% packet loss, time 6134ms
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/bc68c3aeb8bbac32
  {"code":200,"id":"","message":"success"}
  
  # ping目标注入不再有packet loss
  root@:chaos-arsenal# ping 10.103.177.165
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=0.347 ms
  ......
  64 bytes from 10.103.177.165: icmp_seq=9 ttl=64 time=0.252 ms
  ^C
  --- 10.103.177.165 ping statistics ---
  9 packets transmitted, 9 received, 0% packet loss, time 8195ms
  rtt min/avg/max/mdev = 0.208/0.245/0.347/0.039 ms
  ```

  

### 3.8 network-delay（网络延迟）

#### 原理

通过`tc`命令匹配`ip + port`的包做延迟操作。

- 故障注入

  在对应`chain`规则列表末尾添加相应规则。

  ```bash
  # 指定网络接口
  tc qdisc add dev $interface root netem delay $delay_time
  
  # 指定目标ip+端口
  tc qdisc add dev $interface root handle 1: prio bands 4
  tc qdisc add dev $interface parent 1:4 handle 40: netem delay $delay_time
  tc filter add dev $interface protocol ip parent 1:0 prio 4 u32 match ip dst $dst_ip match ip dport $dport 0xffff flowid 1:4
  ```

- 故障清理

  ```bash
  # 指定网络接口
  tc qdisc del dev $interface root netem delay $delay_time
  
  # 指定目标ip+端口
  tc qdisc del dev $interface root handle 1: prio bands 4
  ```
  

#### 对象查找

```bash
ifconfig
ip addr show
```

#### 功能测试

- cli

  ```bash
  # 指定目标网络接口
  arsenal inject hardware network delay --interface ens19 --delay 100ms
  arsenal remove hardware network delay --interface ens19 --delay 100ms
  
  # 指定目标IP延时
  arsenal inject hardware network delay --interface ens19 --destination 10.103.176.177 --delay 100ms
  arsenal remove hardware network delay --interface ens19 --destination 10.103.176.177 --delay 100ms
  
  # 指定目标IP端口号延时	
  arsenal inject hardware network delay --interface ens19 --destination 10.103.176.177 --destination-port 22 --delay 100ms
  arsenal remove hardware network delay --interface ens19 --destination 10.103.176.177 --destination-port 22 --delay 100ms
  ```

- http

  ```bash
  # 指定目标网络接口
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "delay", "params": {"interface":"ens19", "delay": "100ms"}}'
  {"code":200,"id":"051b5833630bf937","message":"success"}
  
  # ping命令可以观察到延时
  root@:chaos-arsenal# ping 10.103.177.165
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=100 ms
  64 bytes from 10.103.177.165: icmp_seq=2 ttl=64 time=100 ms
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/051b5833630bf937
  {"code":200,"id":"","message":"success"}
  
  # ping目标注入不再有延时
  root@:chaos-arsenal# ping 10.103.177.165
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=0.348 ms
  64 bytes from 10.103.177.165: icmp_seq=2 ttl=64 time=0.238 ms
  
  # 指定目标IP延时
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "delay", "params": {"interface":"ens19", "delay": "100ms", "destination": "10.103.176.177"}}'
  {"code":200,"id":"435b5f4be102c487","message":"success"}
  
  # ping 目标ip延迟，非目标ip不延迟。
  [root@localhost chaos-arsenal]# ping 10.103.176.177
  PING 10.103.176.177 (10.103.176.177) 56(84) bytes of data.
  64 bytes from 10.103.176.177: icmp_seq=1 ttl=64 time=100 ms
  64 bytes from 10.103.176.177: icmp_seq=2 ttl=64 time=100 ms
  64 bytes from 10.103.176.177: icmp_seq=3 ttl=64 time=100 ms
  64 bytes from 10.103.176.177: icmp_seq=4 ttl=64 time=100 ms
  
  --- 10.103.176.177 ping statistics ---
  4 packets transmitted, 4 received, 0% packet loss, time 3003ms
  rtt min/avg/max/mdev = 100.243/100.284/100.316/0.028 ms
  [root@localhost chaos-arsenal]# ping 10.103.177.165
  PING 10.103.176.172 (10.103.176.172) 56(84) bytes of data.
  64 bytes from 10.103.176.172: icmp_seq=1 ttl=64 time=0.363 ms
  64 bytes from 10.103.176.172: icmp_seq=2 ttl=64 time=0.243 ms
  64 bytes from 10.103.176.172: icmp_seq=3 ttl=64 time=0.242 ms
  
  --- 10.103.176.172 ping statistics ---
  3 packets transmitted, 3 received, 0% packet loss, time 2055ms
  rtt min/avg/max/mdev = 0.242/0.282/0.363/0.056 ms
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/435b5f4be102c487
  {"code":200,"id":"","message":"success"}
  
  # ping目标注入不再有延迟。
  [root@localhost chaos-arsenal]# ping 10.103.176.177
  PING 10.103.176.177 (10.103.176.177) 56(84) bytes of data.
  64 bytes from 10.103.176.177: icmp_seq=1 ttl=64 time=1.97 ms
  64 bytes from 10.103.176.177: icmp_seq=2 ttl=64 time=0.241 ms
  64 bytes from 10.103.176.177: icmp_seq=3 ttl=64 time=0.263 ms
  
  --- 10.103.176.177 ping statistics ---
  3 packets transmitted, 3 received, 0% packet loss, time 2058ms
  rtt min/avg/max/mdev = 0.241/0.826/1.974/0.811 ms
  [root@localhost chaos-arsenal]# ping 10.103.176.172
  PING 10.103.176.172 (10.103.176.172) 56(84) bytes of data.
  64 bytes from 10.103.176.172: icmp_seq=1 ttl=64 time=0.402 ms
  64 bytes from 10.103.176.172: icmp_seq=2 ttl=64 time=0.266 ms
  64 bytes from 10.103.176.172: icmp_seq=3 ttl=64 time=0.279 ms
  
  --- 10.103.176.172 ping statistics ---
  3 packets transmitted, 3 received, 0% packet loss, time 2085ms
  rtt min/avg/max/mdev = 0.266/0.315/0.402/0.061 ms
  
  # 指定目标ip+port
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "delay", "params": {"interface":"ens19", "delay": "1s", "destination": "10.103.176.177", "destination-port": "22"}}'
  {"code":200,"id":"051b5833630bf937","message":"success"}
  
  # ssh到目标主机时间明显变大。
  [root@localhost chaos-arsenal]# time ssh root@10.103.176.177
  root@10.103.176.177's password:
  
  real    0m8.171s
  user    0m0.008s
  sys     0m0.003s
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/051b5833630bf937
  {"code":200,"id":"","message":"success"}
  
  # ssh到目标主机无明显延时。
  [root@localhost chaos-arsenal]# time ssh root@10.103.176.177
  root@10.103.176.177's password:
  
  real    0m0.732s
  user    0m0.009s
  sys     0m0.004s
  ```



### 3.9 network-reorder（网络乱序）

#### 原理

- 故障注入

  ```bash
  tc qdisc add dev $interface root netem delay $delay reorder $percent $relatper
  ```

- 故障清理

  ```bash
  tc qdisc del dev $interface root netem delay $delay reorder $percent $relatper
  ```


#### 对象查找

```bash
ifconfig
ip addr show
```

#### 功能测试

- cli

  ```bash
  arsenal inject hardware network reorder --interface ens19 --delay 100ms --percent 25% --relatper 50%
  arsenal remove hardware network reorder --interface ens19 --delay 100ms --percent 25% --relatper 50%
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "reorder", "params": {"interface":"ens19", "delay": "100ms", "percent": "25%", "relatper": "50%"}}'
  {"code":200,"id":"051b5833630bf937","message":"success"}
  
  # 在其他主机ping目标接口对应的ip，发包的时间间隔为0.01s，可以看到明显乱序。
  root@:chaos-arsenal# ping 10.103.177.165 -i 0.01
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=2 ttl=64 time=0.257 ms
  64 bytes from 10.103.177.165: icmp_seq=7 ttl=64 time=0.212 ms
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=100 ms
  64 bytes from 10.103.177.165: icmp_seq=3 ttl=64 time=100 ms
  64 bytes from 10.103.177.165: icmp_seq=4 ttl=64 time=100 ms
  64 bytes from 10.103.177.165: icmp_seq=5 ttl=64 time=100 ms
  64 bytes from 10.103.177.165: icmp_seq=6 ttl=64 time=100 ms
  64 bytes from 10.103.177.165: icmp_seq=8 ttl=64 time=100 ms
  64 bytes from 10.103.177.165: icmp_seq=9 ttl=64 time=100 ms
  64 bytes from 10.103.177.165: icmp_seq=10 ttl=64 time=100 ms
  64 bytes from 10.103.177.165: icmp_seq=17 ttl=64 time=0.147 ms
  64 bytes from 10.103.177.165: icmp_seq=11 ttl=64 time=100 ms
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/051b5833630bf937
  {"code":200,"id":"","message":"success"}
  
  # 在其他主机ping目标接口对应的ip，无乱序
  root@:chaos-arsenal# ping 10.103.177.165 -i 0.01
  PING 10.103.177.165 (10.103.177.165) 56(84) bytes of data.
  64 bytes from 10.103.177.165: icmp_seq=1 ttl=64 time=0.356 ms
  64 bytes from 10.103.177.165: icmp_seq=2 ttl=64 time=0.247 ms
  64 bytes from 10.103.177.165: icmp_seq=3 ttl=64 time=0.256 ms
  64 bytes from 10.103.177.165: icmp_seq=4 ttl=64 time=0.252 ms
  64 bytes from 10.103.177.165: icmp_seq=5 ttl=64 time=0.280 ms
  64 bytes from 10.103.177.165: icmp_seq=6 ttl=64 time=0.341 ms
  64 bytes from 10.103.177.165: icmp_seq=7 ttl=64 time=0.220 ms
  64 bytes from 10.103.177.165: icmp_seq=8 ttl=64 time=0.180 ms
  64 bytes from 10.103.177.165: icmp_seq=9 ttl=64 time=0.205 ms
  64 bytes from 10.103.177.165: icmp_seq=10 ttl=64 time=1.46 ms
  64 bytes from 10.103.177.165: icmp_seq=11 ttl=64 time=0.212 ms
  64 bytes from 10.103.177.165: icmp_seq=12 ttl=64 time=0.228 ms
  64 bytes from 10.103.177.165: icmp_seq=13 ttl=64 time=0.225 ms
  ```

  

### 3.10 network-down（网卡down）

#### 原理

- 故障注入

  ```bash
  # 方法1
  nmcli connection down $interface
  
  # 方法2
  ifconfig $interface down
  
  # 方法3
  ifdown $interface
  ```

- 故障清理

  ```bash
  # 方法1
  nmcli connection up $interface
  
  # 方法1
  ifconfig $interface up
  
  # 方法2
  ifup $interface
  ```


#### 对象查找

```bash
ifconfig
ip addr show
```

#### 功能测试

- cli

  ```bash
  arsenal inject hardware network down --interface ens19
  arsenal remove hardware network down --interface ens19
  ```

- http

  ```bash
  [root@localhost chaos-arsenal]# ip addr show ens19
  3: ens19: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
      link/ether fe:fc:fe:79:2f:90 brd ff:ff:ff:ff:ff:ff
      inet 10.103.177.168/20 brd 10.103.191.255 scope global noprefixroute ens19
         valid_lft forever preferred_lft forever
      inet6 fe80::14ca:d86:eb16:1fbf/64 scope link noprefixroute
         valid_lft forever preferred_lft forever
  
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "down", "params": {"interface":"ens19"}}'
  {"code":200,"id":"83736d6b3c44520f","message":"success"}
  
  # 无法ping通网卡ens19设定的ip
  root@:chaos-arsenal# ping 10.103.177.168
  PING 10.103.177.168 (10.103.177.168) 56(84) bytes of data.
  64 bytes from 10.103.177.168: icmp_seq=10 ttl=64 time=0.353 ms
  64 bytes from 10.103.177.168: icmp_seq=11 ttl=64 time=0.322 ms
  64 bytes from 10.103.177.168: icmp_seq=12 ttl=64 time=0.221 ms
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/71fc0d9c1954e369
  {"code":200,"id":"","message":"success"}
  
  # 可以ping通网卡ens19设定的ip
  root@:chaos-arsenal# ping 10.103.177.168
  PING 10.103.177.168 (10.103.177.168) 56(84) bytes of data.
  From 10.103.176.172 icmp_seq=24 Destination Host Unreachable
  From 10.103.176.172 icmp_seq=25 Destination Host Unreachable
  From 10.103.176.172 icmp_seq=26 Destination Host Unreachable
  ```
  
  

### 3.11 network-unavailable（网络不可用）

#### 原理

- 故障注入
  通过`iptables`将出入指定网卡的所有类型报文`drop`掉。

  ```bash
   iptables -I INPUT -i $interface -j DROP
   iptables -I OUTPUT -o $interface -j DROP
  ```

- 故障清理

  ```bash
  iptables -D INPUT -i $interface -j DROP
  iptables -D OUTPUT -o $interface -j DROP
  ```


#### 对象查找

```bash
ifconfig
ip addr show
```

#### 功能测试

- cli

  ```bash
  arsenal inject hardware network unavailable --interface ens19
  arsenal remove hardware network unavailable --interface ens19
  ```

- http

  ```bash
  3: ens19: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
      link/ether fe:fc:fe:78:46:18 brd ff:ff:ff:ff:ff:ff
      altname enp0s19
      inet 10.103.176.174/20 brd 10.103.191.255 scope global noprefixroute ens19
  
  # 故障注入
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "unavailable", "params": {"interface":"ens19"}}'
  {"code":200,"id":"197c5068d891a750","message":"success"}
  
  # 其他节点ping不同目的接口对应的ip，100% packet loss。
  root@:chaos-arsenal# ping 10.103.176.174
  PING 10.103.176.174 (10.103.176.174) 56(84) bytes of data.
  
  --- 10.103.176.174 ping statistics ---
  48 packets transmitted, 0 received, 100% packet loss, time 48168ms
  
  # 清理故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/197c5068d891a750
  {"code":200,"id":"","message":"success"}
  
  # 其他节点可正常ping不同目的接口对应的ip。
  root@:chaos-arsenal# ping 10.103.176.174
  PING 10.103.176.174 (10.103.176.174) 56(84) bytes of data.
  64 bytes from 10.103.176.174: icmp_seq=1 ttl=64 time=0.048 ms
  64 bytes from 10.103.176.174: icmp_seq=2 ttl=64 time=0.039 ms
  64 bytes from 10.103.176.174: icmp_seq=3 ttl=64 time=0.045 ms
  64 bytes from 10.103.176.174: icmp_seq=4 ttl=64 time=0.060 ms
  ^C
  --- 10.103.176.174 ping statistics ---
  4 packets transmitted, 4 received, 0% packet loss, time 3067ms
  rtt min/avg/max/mdev = 0.039/0.048/0.060/0.007 ms
  ```



### 3.12 network-package-drop（网络协议报文丢失）

#### 原理

其中`interface`参数需要根据输入的`chain`类型来指定。

| chain类型                       | interface       |
| ------------------------------- | --------------- |
| INPUT, FORWARD and PREROUTING   | --in-interface  |
| FORWARD, OUTPUT and POSTROUTING | --out-interface |

- 故障注入

  用`iptables`在对应`chain`规则列表末尾添加相应规则。

  ```bash
  iptables -A INPUT --protocol icmp -j DROP $interface eth0 --source $sip --source-port $sport --destination $dip --destination-port $dport
  ```

- 故障清理

  ```bash
  iptables -D INPUT --protocol icmp -j DROP $interface eth0 --source $sip --source-port $sport --destination $dip --destination-port $dport
  ```


#### 对象查找

```bash
ifconfig
ip addr show
```

#### 功能测试

- cli

  ```bash
  # 指定目的端IP的icmp出报文丢失
  arsenal inject hardware network package-drop --chain OUTPUT --interface ens19 --protocol icmp --destination 10.103.176.177
  arsenal remove hardware network package-drop --chain OUTPUT --interface ens19 --protocol icmp --destination 10.103.176.177 
  
  # 指定目的端端口号的icmp出报文丢失
  arsenal inject hardware network package-drop --chain OUTPUT --interface ens19 --protocol icmp --destination-port 22
  arsenal remove hardware network package-drop --chain OUTPUT --interface ens19 --protocol icmp --destination-port 22 
  
  # 指定目的端IP和端口号的icmp出报文丢失
  arsenal inject hardware network package-drop --chain OUTPUT --interface ens19 --protocol icmp --destination 10.103.176.177 --destination-port 22
  arsenal remove hardware network package-drop --chain OUTPUT --interface ens19 --protocol icmp --destination 10.103.176.177 --destination-port 22
  
  # 指定目的端IP的icmp入报文丢失
  arsenal inject hardware network package-drop --chain INPUT --interface ens19 --protocol icmp --source 10.103.176.177
  arsenal remove hardware network package-drop --chain INPUT --interface ens19 --protocol icmp --source 10.103.176.177
  
  # 指定源端IP和端口号的icmp入报文丢失
  arsenal inject hardware network package-drop --chain INPUT --interface ens19 --protocol icmp --source 10.103.176.177 --source-port 22 
  arsenal remove hardware network package-drop --chain INPUT --interface ens19 --protocol icmp --source 10.103.176.177 --source-port 22
  
  # 指定目的端端口号的icmp入报文丢失
  arsenal inject hardware network package-drop --chain INPUT --interface ens19 --protocol icmp --destination-port 22
  arsenal remove hardware network package-drop --chain INPUT --interface ens19 --protocol icmp --destination-port 22
  ```

- http

  ```bash
  # 注入故障
  [root@localhost chaos-arsenal]# curl -X 'POST' 'http://10.103.177.165:9095/arsenal/v1/faults' -H "Content-Type: application/json" -d '{"env": "hardware", "domain": "network", "fault-type": "package-drop", "params": {"interface":"ens19", "chain": "OUTPUT", "protocol": "icmp", "destination": "10.103.176.177"}}'
  {"code":200,"id":"f2f98166bad3a998","message":"success"}
  
  # 无法ping通目标地址
  [root@localhost chaos-arsenal]# ping -I ens19 10.103.176.177
  PING 10.103.176.177 (10.103.176.177) from 10.103.176.174 ens19: 56(84) bytes of data.
  ping: sendmsg: Operation not permitted
  ping: sendmsg: Operation not permitted
  ping: sendmsg: Operation not permitted
  ping: sendmsg: Operation not permitted
  ping: sendmsg: Operation not permitted
  ^C
  --- 10.103.176.177 ping statistics ---
  5 packets transmitted, 0 received, 100% packet loss, time 4081ms
  
  # 恢复故障
  [root@localhost chaos-arsenal]# curl -X DELETE http://10.103.177.165:9095/arsenal/v1/faults/f2f98166bad3a998
  {"code":200,"id":"","message":"success"}
  
  # 可以ping通目标地址
  [root@localhost chaos-arsenal]# ping -I ens19 10.103.176.177
  PING 10.103.176.177 (10.103.176.177) from 10.103.176.174 ens19: 56(84) bytes of data.
  64 bytes from 10.103.176.177: icmp_seq=1 ttl=64 time=0.208 ms
  64 bytes from 10.103.176.177: icmp_seq=2 ttl=64 time=0.147 ms
  64 bytes from 10.103.176.177: icmp_seq=3 ttl=64 time=0.271 ms
  ^C
  --- 10.103.176.177 ping statistics ---
  3 packets transmitted, 3 received, 0% packet loss, time 2025ms
  rtt min/avg/max/mdev = 0.147/0.208/0.271/0.050 ms
  ```



