{{define "content"}}
    <section>
        <h2>{{"{{"}}.RawContent{{"}}"}}</h2>
        <article>{{.RawContent}}</article>
    </section>
    <hr>
    <section>
        <h2>{{"{{"}}.HtmlContent{{"}}"}}</h2>
        <article>{{.HtmlContent}}</article>
    </section>
    <hr>
    <section>
        <h2>{{"{{"}}safeHTML .RawContent{{"}}"}}</h2>
        <article>{{safeHTML .RawContent}}</article>
    </section>
    <hr>
    <section>
        <h2>{{"{{"}}call $.tempConvertHTML .RawContent{{"}}"}}</h2>
        <article>{{call $.tempConvertHTML .RawContent}}</article>
    </section>
    <hr>
    <p><a href="/page">Page render</a></p>
{{end}}