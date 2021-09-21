package send

import (
	"bytes"
	"log"
	"text/template"
)

func ReadTemplate(name string, data interface{}) string {
	t, err := template.ParseGlob("template/*.tpl")
	if err != nil {
		log.Fatalln(err)
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, name, &data); err != nil {
		log.Fatalln(err)
	}

	return tpl.String()
}
