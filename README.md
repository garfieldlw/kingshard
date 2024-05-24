#  kingshard

## Original README
[link](./README_Original.md)

## Overview

kingshard is a high-performance proxy for MySQL powered by Go. Just like other mysql proxies, you can use it to split the read/write sqls. Now it supports basic SQL statements (select, insert, update, replace, delete). The most important feature is the sharding function. Kingshard aims to simplify the sharding solution of MySQL. **The Performance of kingshard is about 80% compared to connecting to MySQL directly.**

## Feature

Based on the functions of the original project, some functions were added or modified:
* supports connect to MySQL 5.7, 8.0. The version of kingshard and the backend mysql must be consistent.
* Supports custom snowflake id algorithm. If necessary, you can modify the corresponding file.
* Use go.uber.org/zap logger to replace the original custom logger.


## License

kingshard is under the Apache 2.0 license. See the [LICENSE](./doc/License) directory for details.
