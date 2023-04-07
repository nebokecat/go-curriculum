# BST - 共通プログラム
## これなに？
[テクノロジー戦略本部_共通_モノづくりプログラム](https://buysell-tech.atlassian.net/wiki/spaces/TECHNOLOGY/pages/2674262528)
<br />
共通プログラムをすすめるためのベースリポジトリ
<br />
基本的には、forkして利用してください
メンターになった人を[コードオーナー](https://docs.github.com/ja/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)に追加するようにしてください

## How
```bash
% git clone https://github.com/[YOUR_GITHUB_USERNAME]/tech-program.git

# server
# http://localhost:8080 にアクセスして graphiQL が表示されればOK！ 
% cd server 
% docker-compose up -d

# frontend
# http://localhost:3000 にアクセスして Hello Next.js の画面が表示されればOK！ 
% cd frontend 
% yarn install
% yarn dev
```
※ リソースDLに時間がかかるのですぐにアクセスしても表示できない場合があります

## なにが入ってる？

### server
- go:latest
- postgreSQL:latest
- gqlgen
- gorm
- gorm/driver/postgres

### frontend
- Next.js(React)
- TypeScript
- React Hook Form
- zod
- MUI
- Apollo Client
- GraphQL Code Generator
