# Terraform

公式ドキュメント: [Documentation | Terraform | HashiCorp Developer](https://developer.hashicorp.com/terraform/docs)

`.tf` ファイルについてのドキュメントは [Configuration Language](https://developer.hashicorp.com/terraform/language) を参照する。

# ディレクトリ構成の参考資料

- [Terraform プロジェクトの効果的なディレクトリ構成パターン(ゼロから始める Terraform 講座～その2～) - ForgeVision Engineer Blog](https://techblog.forgevision.com/entry/Terraform/directory)
- [【初学者向け】Terraformの基本 #Terraform - Qiita](https://qiita.com/yutaroud/items/75915228d787f7b41fed)

# Google Cloud に対する認証

Terraformは以下の方法で認証することができる。

[アプリケーションのデフォルト認証情報の仕組み  |  Google Cloud](https://cloud.google.com/docs/authentication/application-default-credentials?hl=ja)

## [Application Default Credentials](https://cloud.google.com/sdk/gcloud/reference/auth/application-default)を使用するパターン

下記コマンドを実行するとブラウザが起動し、Googleのログイン画面が表示される。
認証するとデフォルトユーザのcredential情報が記載されたJSONが `~/.config/gcloud/application_default_credentials.json` に作成される。

```bash
gcloud auth application-default login
```

## `GOOGLE_APPLICATION_CREDENTIALS` 環境変数を使用するパターン

1. サービスアカウントのkeyを作成する
2. ローカルの任意のフォルダにkeyファイルを配置する
3. `GOOGLE_APPLICATION_CREDENTIALS` 環境変数に、keyファイルまでのpathを設定する

# バックエンドの設定

## Initialize

Backendの設定を変更した際は、必ず `terraform init` すること。

# 📝terraform 運用手順📝

terraformファイルが格納されたディレクトリへ移動する。

```bash
cd ./thanksfulness/terraform
```

## 変数を定義する

1. templateファイルをコピーする
    ```bash
    sh setup-tfvars.sh
    ```
    ※既にtfvarsファイルが存在する場合は何もしない。
2. 変数を定義する
    - オプショナル変数かどうかは、`variables.tf` を参照して判断すること

## リソースを反映する

1. 計画を確認する
    ```bash
    terraform plan -var-file="thanksfulness-prd.tfvars"
    ```
2. 適用する
    ```bash
    terraform apply -var-file="thanksfulness-prd.tfvars"
    ```

## 🚨リソースを削除する🚨

よく注意して実行すること。

```bash
terraform destroy
```

# 既存リソースをterraformに取り込む

[Import - Configuration Language | Terraform | HashiCorp Developer](https://developer.hashicorp.com/terraform/language/import)

# Provider を upgrade する

[Lock and upgrade provider versions | Terraform | HashiCorp Developer](https://developer.hashicorp.com/terraform/tutorials/configuration-language/provider-versioning#upgrade-the-aws-provider-version)

1. `terraform.tf` の `required_providers` ブロック内にある `version` を更新する
2. `terraform init -upgrade` コマンドを実行する
