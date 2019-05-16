<!-- /views/page.html -->
<!doctype html>

<html>
    <head>
        <title>{{.Title}}</title>
        {{include "layouts/head"}}
    </head>

    <body>
        <a href="/"><- Back home!</a>
        {{template "ad" .}}
        <hr>
        {{include "layouts/footer"}}
    </body>
</html>