package wbear

import "html/template"


var htmlText *template.Template = new(template.Template)

var htmlTextDefault string = `<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />

   

    <!--[if lt IE 9]>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.js"></script>
    <![endif]-->
  </head>

  <body>


    <header>
      <h1>HTTP Method is not supported</h1>
    </header>

  </body>
</html>
`


func HTTPMethodFailHTML(html string) error {
  var err error 
  htmlText, err = template.ParseFiles(html)
  if err != nil {
    return err
  } 
  return nil
}
