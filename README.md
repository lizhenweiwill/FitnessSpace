# FitnessSpace

## 数据库设计

### 教练

> 系统用户，比如工作室的接待人员、教练员、系统维护人员，不包括会员。
> 现阶段只有教练员。Coach

| ID  | 姓名  | 昵称  | 密码     |
|-----|-----|-----|--------|
| 1   | 张发  | ZF  | 111111 |
| 2   | 王吉  | WJ  | 222222 |

### 课程

> 工作室的课程类型。Course

| ID  | 课程名称 |
|-----|------|
| 1   | 常规   |
| 2   | 拉伸   |
| 3   | 搏击   |
| 4   | 体态纠正 |

### 会员

> 健身工作室的签约会员。Member

| ID  | 名称  | 身份证号               | 手机号         | 微信    | 头像    |
|-----|-----|--------------------|-------------|-------|-------|
| 1   | 张三  | 111222200002200000 | 11100000000 | wx_xo | a.png |

### 会员-课程-购课详情表

> 【哪位会员】在【什么时间】购买【什么课程】【多少节】实际付款【多少元】

### 会员-课程-销课详情表

> 【哪位会员】在【什么时间】核销【什么课程】【多少节】实际【上课教练】

### 会员-课程-信息统计表

> 【哪位会员】的【什么课程】还剩【多少节】

- 购买课程后，剩余课程数量增加。
- 会员上课后，核销完毕数量减少。

## API

- 教练
    - 查询教练 GET /api/v1/coach/:id
    - 教练列表 GET /api/v1/coach/list
    - 增加教练 POST /api/v1/coach
    - 删除教练 DELETE /api/v1/coach
- 课程
    - 查询课程 GET /api/v1/course/:id
    - 课程列表 GET /api/v1/course/list
    - 增加课程 POST /api/v1/course
    - 删除课程 DELETE /api/v1/course
- 会员
    - 查询会员 GET /api/v1/member/:id
    - 会员列表 GET /api/v1/member/list
    - 新增会员 POST /api/v1/member
    - 修改信息 PUT /api/v1/member
- 记录
    - 记录统计 GET /api/v1/record/list
    - 购买记录 POST /api/v1/record/buy
    - 核销记录 POST /api/v1/record/end

## 环境

### Gin

### gorm

```go
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## SQL

### 教练

```sql

```
