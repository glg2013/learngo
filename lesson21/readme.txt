1. 基本介绍
1.0 项目介绍
Discover 是基于开发语言 PHP7.3，Laravel 框架开发。项目中采用的拓展如下：

dcat/laravel-admin
overtrue/laravel-pinyin
propaganistas/laravel-phone
spatie/laravel-enum
yxx/laravel-quick
zgldh/qiniu-laravel-storage
1.1 适用场景
生产加工羽绒，羽毛制品的厂家。

1.2 安装
Github 地址, gitee 地址
执行 composer install
将 .env.example 复制重命名为 .env, 并在 .env 设置数据库账号密码等信息。
执行 php artisan migrate 生成表结构。
执行 php artisan db:seed --class=InitSeeder 初始化数据库。