package main
import "os"
import S "strings"
import "path"
import "cuelang.org/go/cmd/cue/cmd"
import "fmt"

var extension_out_map = map[string]string{
    ".json": "json",
    ".cue": "cue",
    ".yaml": "yaml",
    ".yml": "yaml",
    ".jsonl": "jsonl",
    ".ldjson": "jsonl",
    ".textproto": "textproto",
    ".proto": "proto",
    ".go": "go",
    ".txt": "text",
    "": "text",
}

func main() {
  files := os.Args[1:]
  fmt.Printf("l1%v\n", "debug")
  for _, file := range files {
    fmt.Printf("l2%v\n", file)
    no_cue_filename := S.TrimSuffix(file, ".cue")
    extension := path.Ext(no_cue_filename)
    out_format := extension_out_map[extension]
    command_string := "export"
    export_command, _ := cmd.New([]string{command_string, file, "--out", out_format, "--outfile", no_cue_filename, "--force"})
    export_command.Execute()
  }

}