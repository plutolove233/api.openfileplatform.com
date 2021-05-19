# 系统错误码

正常 = "200"

账号或密码错误  = "400"



### 普通用户登录注册接口

 1. 登录接口

    | 接口详细 |                             |
    | -------- | --------------------------- |
    | 接口地址 | http://127.0.0.1:9090/login |
    | 请求方式 | POST                        |
    | 操作对象 | normal_users                |

    请求参数（form）

    | 参数         | 参数名称 | 必填 | 类型   | 示例   |
    | ------------ | -------- | ---- | ------ | ------ |
    | name         | 账号名称 | 是   | string | shyhao |
    | pwd          | 账号密码 | 是   | string | 123    |
    | verification | 验证码   | 是   | string | sD2q   |

    

 2. 注册接口