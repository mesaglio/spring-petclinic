package main

import (
	"context"

	"dagger/my-module/internal/dagger"
)

type MyModule struct{}

func (m *MyModule) BuildMaven(ctx context.Context, source *dagger.Directory) *dagger.File {
	return dag.Java().
		WithJdk("17").
		WithMaven("3.9.5").
		WithProject(source.WithoutDirectory("dagger")).
		Maven([]string{"package"}).
		File("target/spring-petclinic-3.3.0-SNAPSHOT.jar")
}

func (m *MyModule) BuildGradle(ctx context.Context, source *dagger.Directory) *dagger.File {
	return dag.Container().
		From("openjdk:17").
		WithExec([]string{"microdnf", "install", "git", "findutils"}).
		// WithMountedCache("/project/.gradle", dag.CacheVolume("gradle-cache")).
		WithWorkdir("/project").
		WithMountedDirectory("/project",
			source.WithoutDirectory("dagger")).
		WithExec([]string{"./gradlew", "build"}).
		File("build/libs/spring-petclinic-3.3.0.jar")
}
