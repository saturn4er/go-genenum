{{- $validationPkg := import "github.com/go-ozzo/ozzo-validation/v4" "validation"}}
{{- $strconvPkg := import "strconv"}}
{{- $fmtPkg := import "fmt"}}
{{- $errorsPkg := import "errors"}}


{{- range $enum := $.enums }}
  type {{$enum.Name}} byte
  const (
  {{$enum.Name}}{{index $enum.Values 0}} {{$enum.Name}} = iota+1
  {{- range $value := (slice $enum.Values 1)}}
  {{$enum.Name}}{{$value}}
  {{- end}}
  )
  {{- if not .Helpers}}
  {{- continue }}
  {{- end }}
  {{- $receiverName := slice $enum.Name 0 1 | lCamelCase}}
  {{- if .Helpers.AllValues.VarName }}
    var {{.Helpers.AllValues.VarName}} = []{{$enum.Name}}{{"{"}}{{- range $value := $enum.Values -}}{{$enum.Name}}{{$value}}, {{- end -}}{{"}"}}
  {{- end }}
  {{- if .Helpers.AllValues.FuncName }}
    func {{.Helpers.AllValues.FuncName}}() []{{$enum.Name}} {
    return []{{$enum.Name}}{{"{"}}{{- range $value := $enum.Values -}}{{$enum.Name}}{{$value}}, {{- end -}}{{"}"}}
    }
  {{- end }}
  {{- if .Helpers.IsValid }}
      func ({{$receiverName}} {{$enum.Name}}) IsValid() bool {
        return {{$receiverName}} > 0 && {{$receiverName}} < {{addInts (len $enum.Values) 1}}
      }
  {{- end }}
  {{- range $category := .Helpers.Categories }}
    var {{$category.Name}}{{plural $enum.Name}} = []{{$enum.Name}}{
        {{- range $categoryValue := $category.Values }}
            {{$enum.Name}}{{$categoryValue}},
        {{- end }}
    }
    func ({{$receiverName}} {{$enum.Name}}) Is{{$category.Name}}() bool {
      if {{$receiverName}} < 1 || {{$receiverName}} > {{addInts (len $enum.Values) 1}} {
        return false
      }
      return []bool{false,
        {{- range $value := $enum.Values -}}
            {{- $isTrue := false -}}
            {{- range $isCategoryValue := $category.Values }}
                {{- if eq $value $isCategoryValue  }}
                  {{- $isTrue = true -}}
                {{- end }}
            {{- end -}}
            {{$isTrue}},
        {{- end -}}}[{{$receiverName}}]
    }
  {{- end }}
  {{- if .Helpers.Is }}
    {{- range $value := $enum.Values }}
      func ({{$receiverName}} {{$enum.Name}}) Is{{$value}}() bool {
        return {{$receiverName}} == {{$enum.Name}}{{$value}}
      }
    {{- end}}
  {{- end }}
  {{- if .Helpers.Validate }}
    type Invalid{{$enum.Name}}ValueError byte
    func (e Invalid{{$enum.Name}}ValueError) Error() string {
    return {{$fmtPkg.Ref "Sprintf"}}("invalid {{$enum.Name}}(%d)", e)
    }
    func ({{$receiverName}} {{$enum.Name}}) Validate() error {
      if {{$receiverName}} < 1 || {{$receiverName}} > {{addInts (len $enum.Values) 1}} {
        return Invalid{{$enum.Name}}ValueError({{$receiverName}})
      }
    return nil
    }
  {{- end }}
  {{- if .Helpers.String }}
    func ({{$receiverName}} {{$enum.Name}}) String() string {
    if {{$receiverName}} < 1 || {{$receiverName}} > {{len $enum.Values}} {
    return "{{$enum.Name}}("+{{$strconvPkg.Ref "FormatInt"}}(int64({{$receiverName}}), 10)+")"
    }
    const names = "{{range $i, $v := .Values}}{{if $i}}{{end}}{{$v}}{{end}}"
    {{$totalLength := 0}}
    var indexes = [...]int32{{"{0, "}} {{range $i, $v := .Values}}{{- $totalLength = addInts $totalLength (len $v) -}}{{$totalLength}}, {{- end -}}{{"}"}}

    return names[indexes[{{$receiverName}}-1]:indexes[{{$receiverName}}]]
    }
  {{- end}}
{{- end }}
