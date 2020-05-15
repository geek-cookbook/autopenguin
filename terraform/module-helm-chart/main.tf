resource "github_repository" "helm-harbor" {
  name        = "helm-harbor"
  description = "A Helm chart for Harbor (supports Istio mTLS)"
  homepage_url = "https://goharbor.io"
  has_issues = true
  private = false

  template {
    owner = "geek-cookbook"
    repository = "template-helm-chart"
  }
}

resource "github_branch" "gh-pages" {
  repository = "helm-harbor"
  branch     = "gh-pages"
}

resource "github_branch_protection" "protect-master" {
  repository     = "helm-harbor"
  branch         = "master"
  enforce_admins = true
  require_signed_commits = true

  # required_status_checks {
  #   strict   = false
  #   contexts = ["ci/travis"]
  # }

  # required_pull_request_reviews {
  #   dismiss_stale_reviews = true
  #   dismissal_users       = ["foo-user"]
  #   dismissal_teams       = ["${github_team.example.slug}", "${github_team.second.slug}"]
  # }

  # restrictions {
  #   users = ["foo-user"]
  #   teams = ["${github_team.example.slug}"]
  #   apps  = ["foo-app"]
  # }
}


# resource "github_repository_collaborator" "repo_collaborator_autopenguins" {
#   repository = "helm-harbor"
#   username   = "orgs/geek-cookbook/teams/autopenguins"
#   permission = "push"
# }

# We need this secret for chart-releaser action
resource "github_actions_secret" "cr_token" {
  repository       = "helm-harbor"
  secret_name      = "CR_TOKEN"
  plaintext_value  = "${var.github_token}"
}

# # We need this to trigger a webhook when an action fails
# resource "github_actions_secret" "slack_webhook" {
#   repository       = "helm-harbor"
#   secret_name      = "SLACK_WEBHOOK"
#   plaintext_value  = "${var.slack_webhook}"
# }