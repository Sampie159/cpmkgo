package cpmkgo

import (
	"fmt"
	"os"
)

func SetupProject(language, projectName string) error {
	if !isValid(language) {
		return fmt.Errorf("Invalid language: %s", language)
	}

	projectPath, err := os.Getwd()
	if err != nil {
		return err
	}

	projectPath += "/" + projectName
	err = os.MkdirAll(projectPath+"/src", 0755)
	if err != nil {
		return err
	}

	err = createFiles(language, projectName, projectPath)
	fmt.Println("Project", projectName, "created successfully!")

	return err
}

func isValid(language string) bool {
	return language == "c" || language == "cpp"
}

func createFiles(language, projectName, projectPath string) error {
	mainPath := projectPath + "/src/main." + language
	cmakePath := projectPath + "/CMakeLists.txt"
	cmakeSrcPath := projectPath + "/src/CMakeLists.txt"
	mainContent := ""
	cmakeContent := ""
	cmakeSrcContent := ""

	if language == "c" {
		mainContent = `#include <stdio.h>

		int main(void) {
			printf("Hello, World!\n");

			return 0;
		}`

		cmakeContent = `cmake_minimum_required(VERSION 3.10)

		project(` + projectName + `)

		set(CMAKE_C_STANDARD 17)
		set(CMAKE_C_STANDARD_REQUIRED True)
		set(CMAKE_C_FLAGS "-Wall -Wextra -Wpedantic")

		set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR})

		add_subdirectory(src)`

		cmakeSrcContent = `add_executable(
			` + projectName + `
			main.c
		)`
	} else {
		mainContent = `#include <iostream>

		int main() {
			std::cout << "Hello World!\n";

			return 0;
		}`

		cmakeContent = `cmake_minimum_required(VERSION 3.10)

		project(` + projectName + `)

		set(CMAKE_CXX_STANDARD 20)
		set(CMAKE_CXX_STANDARD_REQUIRED True)
		set(CMAKE_CXX_FLAGS "-Wall -Wextra -Wpedantic")

		set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR})

		add_subdirectory(src)`

		cmakeSrcContent = `add_executable(
		` + projectName + `
		main.cpp
		)`
	}

	err := os.WriteFile(mainPath, []byte(mainContent), 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(cmakePath, []byte(cmakeContent), 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(cmakeSrcPath, []byte(cmakeSrcContent), 0644)

	return nil
}
