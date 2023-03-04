[![Go Report Card](https://goreportcard.com/badge/github.com/LukoJy3D/perfect100)](https://goreportcard.com/report/github.com/LukoJy3D/perfect100)

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->

<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![Star on GitHub](https://img.shields.io/github/stars/LukoJy3D/perfect100.svg?style=social)](https://github.com/LukoJy3D/perfect100/stargazers)

Repo for those crazy achievement hunters who just love going for that 100%. After that grindy completion, nothing better as sharing the knowledge with others!
Feel free to contribute, make feature/guide requests and ask for help with hard achievements. This repo is and probably will always be a work in progress type of thing for me, so I hope we will find more minds alike :]

## Workflow

To keep all information more structured most parts of the process are automated. So to create new guide:

1. Add game id and name to [games.json](/games.json). (ID can be found in games store page url or steamdb.info).
2. Run `go run .\build.go`, this also will be triggered when opening a PR, so guide can be added over multiple PRs.
3. Dir/file structure will be generated under /guides. Every achievement will have dedicated markdown file, so feel free adding guides even one by one.

## Requirements

As we like polished steam profiles, we like neat Pull Requests and Issues just as well!

1. Use proper commit names (according to [requirements](.github/workflows/commitlint.config.js)):
- type is one of `games`, `guides`, `multi`, `tools`
- subject is always Sentense-cased
- Examples when editing games.json
  - games: Add "Among Us"
- Examples when adding guides
  - guides: Create 'Cyberpunk 2077' 'Legends of the afterlife'
  - guides: Create all 'MultiVersus' achievement guides
- Examples when changing multiple categories
  - multi: Add 'Lost Ark' and change readme
  - multi: Add 'Fall Guys' guides and update commitlint config
- Examples when changing scripts, readme, workflows and etc.
  - tools: Improve commitlint workflow
  - tools: Make build.go mo dynamic

2. Use labels when creating issues, request it if you are missing any.

3. Let's do quality over quantity! It's always better to take time and make clear, informative and structured guides instead of confusing walpaper lenght paragraphs or simple lazy oneliners.

4. Use one commit per PR. Some functionalities depend on it, so if needed - squash those commits or even try using --amend flag when commiting. It's very handy thing to learn!

## Contributors ✨

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->

<!-- prettier-ignore-start -->

<!-- markdownlint-disable -->

<!-- markdownlint-restore -->

<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->