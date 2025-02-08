package main

import "testing"
import "os"
import "github.com/stretchr/testify/assert"

const expected_result_file_content = `{
    "a": 1
}
`

func TestRunForFiles(t *testing.T) {
    //given
    tmp_dir := t.TempDir()
    cue_file_content := "a: 1\n"

    cue_file_path := tmp_dir + "/some_file.json.cue"
    expected_result_file_path := tmp_dir + "/some_file.json"
    f, _ := os.Create(cue_file_path)
    defer f.Close()
    f.WriteString(cue_file_content)

    // when
    run_for_files([]string{cue_file_path})

    // then
    _, err1 := os.Stat(expected_result_file_path)
    assert.Nil(t, err1)
    actual_json_file_content, _ := os.ReadFile(expected_result_file_path)
    assert.Equal(t, expected_result_file_content, string(actual_json_file_content))
}
