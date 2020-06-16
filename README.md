# SendGird

[Github(sendgrid-go)](https://github.com/sendgrid/sendgrid-go)

[日本語doc](https://sendgrid.kke.co.jp/docs/API_Reference/Web_API_v3/index.html)

### DynamicTemplate
```
  <html>
    <head>
      <title></title>
    </head>
    <body>
      <div>

        <p>
            <img src="画像のURL" height=300 weight=200>
        </p>

        ――――――――――<br>
        ■応募者情報<br>
        ・氏名：{{NAME_KANA}}<br>
        ・年齢：{{AGE}}<br>
        ・性別：{{SEX}}<br>
        ・電話番号：{{TEL1}}<br>
        ・電話番号2：{{TEL1}}<br>
        ・メールアドレス：{{EMAIL1}}<br>
        ・メールアドレス2：{{EMAIL2}}<br>
        {{#if ADDRESS}}
            ・住所：{{ADDRESS}}<br>
        {{else}}
            ・住所未入力<br>
        {{/if}}
        ――――――――――<br>

      </div>
    </body>
  </html>
```

## Usage
1. [こちら](https://sendgrid.kke.co.jp/blog/?p=11818)を参考にAPIキーを作成する。
2. `cat sendgrid.env.example | sed s/\<SENDGRID_API_KEY\>/<ここに作成したAPIキーを入力する>/g > sendgrid.env`を実施し、sendgrid.envを作成する。
3. sendgrid.envの`EMAIL=`の後に送信したいアドレスを入れる。  
※ DynamicTemplateを利用したい場合は別途テンプレートを登録する必要がある。
4. `make up`でコンテナを起動させる。

### メールの送信
- helperPost
```
$ curl http://0.0.0.0:8000/helper
```

- notHelperPost  
※ `"email": "test@example.com"`を変更する必要あり。。。
```
$ curl http://0.0.0.0:8000/notHelper
```

- sendDynamicTemplateEmail
```
$ curl http://0.0.0.0:8000/dynamicTemplate
```