package main

import "testing"
import "os"
import "github.com/stretchr/testify/assert"
// import "reflect"
// import "sort"
const expected_result_file_content = `{
    "a": 1
}
`

func TestRunForFilesWithoutParametrization(t *testing.T) {
    // given
    tmp_dir := t.TempDir()
    cue_file_content := "a: 1\n"

    cue_file_path := tmp_dir + "/some_file.json.cue"
    expected_result_file_path := tmp_dir + "/some_file.json"
    f, _ := os.Create(cue_file_path)
    defer f.Close()
    f.WriteString(cue_file_content)

    // when
    run_for_files([]string{cue_file_path}, map[string][]string{})

    // then
    _, err1 := os.Stat(expected_result_file_path)
    assert.Nil(t, err1)
    actual_json_file_content, _ := os.ReadFile(expected_result_file_path)
    assert.Equal(t, expected_result_file_content, string(actual_json_file_content))
}


const expected_result_file_content1 = `{
    "a": "dev"
}
`

const expected_result_file_content2 = `{
    "a": "prod"
}
`

func TestRunForFilesWithParametrization(t *testing.T) {
    //given
    tmp_dir := t.TempDir()
    cue_file_content := "a: _ @tag(env)\n"

    parametrizationMap := map[string][]string{
        "env": []string{"dev", "prod"},
    }

    cue_file_path := tmp_dir + "/some_file_{env}.json.cue"
    expected_result_file_path1 := tmp_dir + "/some_file_dev.json"
    expected_result_file_path2 := tmp_dir + "/some_file_prod.json"
    f, _ := os.Create(cue_file_path)
    defer f.Close()
    f.WriteString(cue_file_content)

    // when
    run_for_files([]string{cue_file_path}, parametrizationMap)

    // then
    _, err1 := os.Stat(expected_result_file_path1)
    assert.Nil(t, err1)
    actual_json_file_content1, _ := os.ReadFile(expected_result_file_path1)
    assert.Equal(t, expected_result_file_content1, string(actual_json_file_content1))

    _, err2 := os.Stat(expected_result_file_path2)
    assert.Nil(t, err2)
    actual_json_file_content2, _ := os.ReadFile(expected_result_file_path2)
    assert.Equal(t, expected_result_file_content2, string(actual_json_file_content2))
}


func TestParseFilename(t *testing.T) {
    // given
    filename := "/some_file_{env}.json.cue"

    // when
    result := parse_filename(filename)

    // then
    assert.Equal(t, "/some_file_{env}.json", result.no_cue_filename)
    assert.Equal(t, ".json", result.extension)
    assert.Equal(t, []string{"env"}, result.parametrizationVariables)
}

func TestCartesianProduct(t *testing.T) {
    // given
    input := map[string][]string{
        "z": []string{"a", "b"},
        "y": []string{"1", "2"},
    }

    // when
    result := cartesianProduct(input)

    // then
    expected_result := []map[string]string{
        {"z": "a", "y": "1"},
        {"z": "a", "y": "2"},
        {"z": "b", "y": "1"},
        {"z": "b", "y": "2"},
    }
    assert.ElementsMatch(t, result, expected_result)
}


