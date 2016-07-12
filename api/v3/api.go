package v3api

import "strings"

func ListApps(guids, names, orgGUIDS, spaceGUIDS []string) (string, error) {
	path := "/v3/apps?"
	if len(guids) > 0 {
		path = addParameter(path, "guids", guids)
	}
	if len(names) > 0 {
		path = addParameter(path, "names", names)
	}
	if len(spaceGUIDS) > 0 {
		path = addParameter(path, "space_guids", spaceGUIDS)
	}
	if len(orgGUIDS) > 0 {
		path = addParameter(path, "org_guids", orgGUIDS)
	}

	return path, nil
}

func addParameter(path string, parameter string, value []string) string {
	if strings.Contains(path, "=") {
		path += "&" + parameter + "="
	} else {
		path += "?" + parameter + "="
	}
	return path + strings.Join(value, ",")
}
