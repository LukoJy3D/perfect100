[![Go Report Card](https://goreportcard.com/badge/github.com/LukoJy3D/perfect100)](https://goreportcard.com/report/github.com/LukoJy3D/perfect100)
[![Star on GitHub](https://img.shields.io/github/stars/LukoJy3D/perfect100.svg?style=social)](https://github.com/LukoJy3D/perfect100/stargazers)

Repo is for those crazy achievement hunters who love going for that 100%. After that grinding completion, nothing better than sharing the knowledge with others!
Feel free to contribute, make feature/guide requests, and ask for help with hard achievements. This repo is and probably will always be a work-in-progress type of thing for me, so I hope we will find more minds alike :]

## Workflow

To keep all information more structured most parts of the process are automated. So to create a new guide:

1. Add game id and name to [games.json](/games.json). (game id can be found on the steam store page URL or steamdb.info).
2. Wait for GitHub actions to do the rest

## Requirements

As we like polished steam profiles, we also like neat Pull Requests and Issues!

1. Use proper commit names (according to [requirements](.github/workflows/commitlint.config.js)):

- type is one of `games`, `guides`, `multi`, `tools`
- the subject is always Sentence-cased
- Examples when editing games.json
  - games: Add "Among Us"
- Examples when adding guides
  - guides: Create 'Cyberpunk 2077' 'Legends of the afterlife'
  - guides: Create all 'MultiVersus' achievement guides
- Examples when changing multiple categories
  - multi: Add 'Lost Ark' and change readme
  - multi: Add 'Fall Guys' guides and update commitlint config
- Examples when changing scripts, readme, workflows, etc.
  - tools: Improve commitlint workflow
  - tools: Make build.go mo dynamic

2. Use labels when creating issues. Request it if you are missing any.

3. Let's do quality over quantity! It's always better to take time and make clear, informative, and structured guides instead of confusing wallpaper-length paragraphs or simple lazy one-liners.

4. Use one commit per PR. Some functionalities depend on it, so if needed - squash those commits or even try using the --amend flag when committing. It's a very handy thing to learn!

## To-do list

- [ ] Make a separation between human and AI-made guides
- [ ] Launch via GitHub pages
- [ ] More implementation with the discord community
- [ ] More contribution acknowledgment
- [ ] Markdown linters?
