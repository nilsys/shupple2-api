# コードスタイル及び開発フロー

## 新規API実装フロー

1. 新しくテーブルが必要な場合、/migrationsにテーブル作成のSQLを追加します。
2. データ参照、永続化(新規作成、更新)、削除に関してはrepositoryに実装します。
3. 依存関係が逆転しない様pkg/domain/repositoryにinterfaceを作成します。
4. 参照と登録系でrepositoryを分けます。
5. 命名規則は参照系の場合entity名_query.go登録系の場合entity名_command.goです。
6. 次はinterfaceの実装をpkg/adaptor/infrastructure/repositoryに作成します。命名規則は上記と同じです。
7. wireで依存(DI)を設定するので他の実装を参考にRepoSetを定義して下さい。
8. interfaceを継承して実装をします。
9. 次はアプリケーション層のserviceを実装します。
10. serviceに関してはinterfaceも用意しますが、依存関係に影響を与えない為、同ファイルに継承も実装します。
11. 7と同じ様にServiceSetを定義して下さい。
12. 実装、命名規則は5と同じです。また、参照と登録系はrepositoryと同じで分けて下さい。
13. Entityなどのdomainが新たに必要な場合は適宜domain/に追加します。
14. 次にcontrollerの実装です。
15. こちらも上記と同じ様に、pkg/adaptor/apiに参照と登録系を分けます。また、命名規則も合わせて下さい。
16. レスポンスモデルとEntityが異なる場合のconverterが必要な場合はpkg/adaptor/api/converterに実装
17. レスポンスモデルはpkg/adaptor/responseにモデルを実装
18. repositoryとservice、controllerの実装が終われば、cmd/wire.goに各Setを追記して下さい。

## 注意点

- 各レイヤーの責務と違う実装がされていないか
- レイヤーの依存関係が一方向であるか
- コメントは全てのコードに記述されているか
- ロジックをドメインモデルに閉じ込めているか
- ローカルでもgolang-ci-lintを実行しながら開発

## 命名規則（参考）
- 参照 -> FindBy, Show
- 登録 -> Store
- 更新 -> Store, Update
- 削除 -> Delete
