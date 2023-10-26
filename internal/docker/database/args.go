package database

func setInitializer() {
	buildDatabaseCMD = append(buildDatabaseCMD, "docker", "run", "-d")
}

func setDatabaseName(databaseName string) {
	buildDatabaseCMD = append(buildDatabaseCMD, "--name", databaseName)
}

func setNetworks(networks []string) {
	if len(networks) == 0 {
		return
	}

	for _, network := range networks {
		buildDatabaseCMD = append(buildDatabaseCMD, "--network", network)
	}
}

func setEnv(env string) {
	buildDatabaseCMD = append(buildDatabaseCMD, "-e", env)
}

func setExposedPorts(exposePorts []string) {
	if len(exposePorts) == 0 {
		return
	}

	for _, port := range exposePorts {
		buildDatabaseCMD = append(buildDatabaseCMD, "-p", port+":"+port)
	}
}

func setDatabaseType(databaseType string) {
	buildDatabaseCMD = append(buildDatabaseCMD, databaseType)
}
