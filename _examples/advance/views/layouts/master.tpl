<!-- /views/admin/master.html -->

<!doctype html>

<html>
    <head>
        <title>{{.Title}}</title>
        {{include "layouts/head"}}
    </head>
    <body>
        {{safeHTML "<!-- body is here!!! -->"}}
        {{template "content" .}}
        <hr>
        {{template "ad" .}}
        <hr>
        {{include "layouts/footer"}}
    </body>
</html>