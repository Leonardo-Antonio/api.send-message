{{ define "message" }}

{{ template "header" }}

    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Assumenda delectus sapiente dicta? Ratione laborum atque, accusantium maiores animi sit at iure! Voluptate quae neque tempora eveniet qui adipisci, quia libero.</p>

    <strong>Name: </strong> <span>{{ .Name }}</span>
    <strong>Message: </strong> <span>{{ .Text }}</span>

{{ template "footer" }}

{{ end }}