workflow "Build and Push" {
  on = "push"
  resolves = ["build image", "login to registry", "push to registry"]
}

action "build image" {
  uses = "actions/docker/cli@76ff57a"
  args = "build -t sjpuas/my-first-cli:0.0.1 ."
}

action "login to registry" {
  uses = "actions/docker/login@76ff57a"
  needs = ["build image"]
  secrets = ["DOCKER_USERNAME", "DOCKER_PASSWORD"]
}

action "push to registry" {
  uses = "actions/docker/cli@76ff57a"
  needs = ["login to registry"]
  args = "push sjpuas/my-first-cli:0.0.1"
}
