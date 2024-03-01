こちらは学習用に作成した掲示板アプリのバックエンドのリポジトリになります。フロントエンドのリポジトリは[こちら](https://github.com/progsystem926/bbs-nextjs-go-front)です。

# 概要

Next.js と Go 言語(echo)の学習のため、掲示板アプリを作成しました。
ログイン機能、ユーザ登録機能、投稿機能、投稿の削除機能などを実装しました。

## URL

https://bbs-nextjs-go.click

### 使い方

テストユーザとして、下記のアカウントでログインできます。  
Email: test1@example.com  
Password: password1

## 使用技術一覧

**バックエンド:** Go 1.21.7 / echo 4.11.4

- コード解析: golangci-lint
- 主要パッケージ: Goose / GORM / gqlgen

**フロントエンド:** TypeScript 5.3.3 / React 18.2.0 / Next.js 14.1.0

- コード解析: ESLint
- フォーマッター: Prettier
- テストフレームワーク: Jest
- CSS フレームワーク: Tailwind CSS / daisyUI
- 主要パッケージ: React Hook Form / Recoil / react-cookie / GraphQL Code Generator / graphql-request / Tanstack Query

**インフラ:** AWS(Route53 / Certificate Manager/ Systems Manager Parameter Store / ALB / VPC / ECR / ECS Fargate / RDS MySQL / S3) / Nginx

**CI / CD:** GitHub Actions / AWS(CodePipeline)

**環境構築:** Docker / Docker Compose

## 機能

- メールアドレスとパスワードを利用したユーザー登録 / ログイン・ログアウト機能
- 投稿の一覧表示機能
- 新規の投稿作成機能
- 投稿の削除機能
