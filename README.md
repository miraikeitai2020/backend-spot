# backend-spot

<p align="left">
  <a href="https://github.com/actions/setup-node/actions?query=workflow%3Abuild-test"><img alt="build-test status" src="https://github.com/miraikeitai2020/backend-spot/workflows/Go/badge.svg"></a>
</p>

<img src="./img/michishirube.png" width="400" alt="architecture" />

`backend-spot`APIはミライケータイプロジェクト2020 [mirai-cluster](https://github.com/miraikeitai2020/mirai-cluster)のサービスの一つです． 
APIの主な機能は  
- スポットの選出
- 寄り道スポットの選出
- 寄り道スポットの登録

## Description
### Endpoints
|Method|URL|Description|
|:-:|:-:|:-|
|POST|/query/spot|Get spot infomation|
|POST|/query/detour|Get detour infomation|
|POST|/mutation/add/spot|Add spot infomation|

## Other
- [Commit Rules](./docs/commit_rules.md)
- [Branch Rules](./docs/branch_rules.md)
- [Issue Rules](./docs/issue_rules.md)
