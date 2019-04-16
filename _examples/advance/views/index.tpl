{{define "content"}}
    <h1 class="hello">This is content!!!!</h1>
    <p>123 + 333 = {{call $.add 123 333}}</p>
    <p>123 - 100= {{sub 123 100}}</p>
    <hr>
    <p><a href="/page">Page render</a></p>
{{end}}