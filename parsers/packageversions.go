package parsers

func ParsePackageVersions(packageJson PackageJSONData, filters map[string]bool, exclude bool) [][2]string {
	deps := packageJson.Dependencies
	var packageVersions [][2]string
	for dep, version := range deps {
		if len(filters) == 0 {
			packageVersions = append(packageVersions, [2]string{dep, version})
		} else {
			_, inFilter := filters[dep]
			if exclude && !inFilter {
				packageVersions = append(packageVersions, [2]string{dep, version})
			}
			if !exclude && inFilter {
				packageVersions = append(packageVersions, [2]string{dep, version})
			}
		}
	}
	return packageVersions
}
