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
        <p><img src="{{img}}"></p>

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