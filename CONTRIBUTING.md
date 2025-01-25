# Contributing

We welcome any contributions and are willing to help you get that 100% completion, making our project and your gaming experience better.

## Contribution Guide

Follow the steps below to contribute to the main repository via pull request. You can learn about the details of pull requests [here](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests).

### 1. Fork the Official Repository

Firstly, you must visit the [perfect100 repository](https://github.com/lukojy3d/perfect100.git) and log into your account. The `fork` button is at the top right corner of the web page alongside buttons such as `watch` and `star`.

Now, you can clone your forked repository into your local environment.

```shell
git clone https://github.com/<YOUR-USERNAME>/perfect100.git
```

### 2. Create a New Branch

It would be best not to change your forked repository's main branch, as this might make upstream synchronization difficult. You can create a new branch with the appropriate name. General branch name format should start with `fix/` and `feat/`. `fix` is for minimal adjustments (similar to a bug fix), and `feat` is for adding a new feature, guide, game list, etc.

```shell
git checkout -b <NEW-BRANCH-NAME>
```

### 3. Making changes without local setup

Most parts of the process are automated by GitHub action to keep all information more structured.

Sometimes there is no need to run Golang on your local environment as you want to add a couple of game guides for existing games in which guide files are already generated. Adding a new game can be done with the following workflow:

1. Add game id and name to [games.yml](/games.yml). (You can find the game id on the Steam store page URL or steamdb.info).
2. Commit your changes and open a pull request.
3. Wait for GitHub actions to do the rest (it will generate an achievement list and stats of players)

### 4. Running guide generation (local development)

1. Set up Golang depending on your platform.
2. Add game title and id in [games.yml](/games.yml). Example

```yaml
description: List of added in progress or finnished game guides
please: Keep sorted by id!
games:
  - id: "620980" #Beat Saber
  - id: "881100" #Noita
  - id: "1091500" #Cyberpunk 2077
```

_You can find the Game ID on the Steam store page or steamdb.info_ 3. Do `go run main.go "<game name>"`. 4. Script will populate the achievement list and guide folder contents after a script successfully runs.

### 5. Code Commit to a fork

You can commit and push the changes to your local repository.

```shell
git add -A
git commit -m "<COMMIT-MESSAGE>"
git push
```

### 6. Open a Pull Request

You can now create a pull request on the GitHub webpage of your repository. As we like polished Steam profiles, we also like neat Pull Requests and Issues! So we have a few requirements:

1. Use proper commit names, based on [Conventional Commits specification](https://www.conventionalcommits.org):

`<type>[optional scope]: <description>`

Examples:

`feat(games): add 'Among Us'`

`feat(guides): create 'Cyberpunk 2077' 'Legends of the afterlife'`

`fix(Cyberpunk2077): typo 'Legends of the afterlife'`

1. Let's do quality over quantity! It's always better to take time and make clear, informative, and structured guides instead of confusing wallpaper-length paragraphs or simple lazy one-liners.

### 7. Create an issue

Not everyone is familiar with GitHub, and it might be difficult for first-time contributors to follow such a workflow. In such cases, write your content in an issue, and we will help you turn that issue into a pull request.
