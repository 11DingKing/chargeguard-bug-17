# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

监管端给隐患检索设置了 200 毫秒超时，请求断开后数据库查询仍持续数秒。先不要改目标仓库代码，请沿 HTTP、service 到 repository 的调用过程定位 deadline 在哪里没有继续生效，并提供耗时或取消结果以及对应文件、符号作为证据。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-17
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-17.git
- parent SHA：a22b588573d4a53ca019b5e72591fb2e0adf5a0f

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-17.git bug-repro
cd bug-repro
git checkout --detach a22b588573d4a53ca019b5e72591fb2e0adf5a0f
go test ./internal/charging -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/charging -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:22: err=search hazards: query continued after deadline
FAIL
FAIL	chargeguard/internal/charging	0.037s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/charging -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:22: err=search hazards: query continued after deadline
FAIL
FAIL	chargeguard/internal/charging	0.002s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

根因结论必须准确写明出问题的 Go 文件、具体符号和完整失效机制，并由实际复现、源码调查和验证证据支撑；调查结束时目标仓库代码、测试和配置零改动。
