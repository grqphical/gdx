package analysis_test

import (
	"gdx/analysis"
	"reflect"
	"testing"
)

func TestGodotProjectConfigParser(t *testing.T) {
	example := `; Engine configuration file.
; It's best edited using the editor UI and not directly,
; since the parameters that go here are not all obvious.
;
; Format:
;   [section] ; section goes between []
;   param=value ; assign values to parameters

config_version=5

[application]

config/name="Strategy Game"
run/main_scene="uid://rrg65t6adsor"
config/features=PackedStringArray("4.4", "GL Compatibility")
config/icon="res://icon.svg"

[input]

forward={
"deadzone": 0.2,
"events": [Object(InputEventKey,"resource_local_to_scene":false,"resource_name":"","device":-1,"window_id":0,"alt_pressed":false,"shift_pressed":false,"ctrl_pressed":false,"meta_pressed":false,"pressed":false,"keycode":0,"physical_keycode":87,"key_label":0,"unicode":119,"location":0,"echo":false,"script":null)
]
}
back={
"deadzone": 0.2,
"events": [Object(InputEventKey,"resource_local_to_scene":false,"resource_name":"","device":-1,"window_id":0,"alt_pressed":false,"shift_pressed":false,"ctrl_pressed":false,"meta_pressed":false,"pressed":false,"keycode":0,"physical_keycode":83,"key_label":0,"unicode":115,"location":0,"echo":false,"script":null)
]
}`

	expectedInputs := []analysis.InputConfig{
		{
			Name: "forward",
		},
		{
			Name: "back",
		},
	}

	projectConfig, err := analysis.ParseGodotProjectFile([]byte(example))
	if err != nil {
		t.Errorf("error while parsing: %s\n", err)
	}

	if projectConfig.ApplicationName != `"Strategy Game"` {
		t.Errorf("expected '%s', got '%s'\n", `"Strategy Game"`, projectConfig.ApplicationName)
	}

	if !reflect.DeepEqual(projectConfig.InputConfigs, expectedInputs) {
		t.Errorf("expected '%+v', got '%+v'\n", expectedInputs, projectConfig.InputConfigs)
	}

}
